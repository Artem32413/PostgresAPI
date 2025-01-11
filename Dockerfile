# FROM golang:latest AS builder  
# WORKDIR /postgresApi  
# COPY go.mod go.sum ./  
# RUN go mod download  
# COPY ./ ./  
# RUN go build -o main .  

# EXPOSE 8081  
# CMD [ "./main" ]  

# FROM golang:latest AS builder
# WORKDIR /postgresApi 
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o main

# FROM alpine:3.13
# RUN apk update && apk --no-cache add bash
# COPY --from=builder /postgresApi/main /bin/  
# COPY --from=builder /postgresApi/wait-for-it.sh /bin/wait-for-it.sh 
# EXPOSE 8081
# RUN chmod +x /bin/wait-for-it.sh /bin/main 
# Этап сборки
FROM golang:latest AS builder  
WORKDIR /postgresApi  

# Копируем файлы модуля и загружаем зависимости
COPY go.mod go.sum ./  
RUN go mod download  

# Копируем остальные файлы
COPY . .  

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Этап выполнения
FROM alpine:3.13
RUN apk update && apk --no-cache add bash

# Копируем скомпилированное приложение и скрипт
COPY --from=builder /postgresApi/main /bin/  
COPY --from=builder /postgresApi/wait-for-it.sh /bin/wait-for-it.sh 

# Открываем порт и устанавливаем права
EXPOSE 8081
RUN chmod +x /bin/wait-for-it.sh /bin/main

# Команда для запуска приложения
CMD [ "/bin/wait-for-it.sh", "localhost:5432", "--timeout=30", "--", "/bin/main" ]


FROM golang:latest AS builder  
WORKDIR /newappapilatest  
COPY go.mod go.sum ./  
RUN go mod download  
COPY ./ ./  
RUN go build -o main .  

EXPOSE 8080  
CMD [ "./main" ]  
  
# # Используем официальный образ Golang для сборки  
# FROM golang:latest AS builder  

# # Устанавливаем рабочую директорию  
# WORKDIR /app  

# # Копируем файлы зависимостей и загружаем зависимости  
# COPY go.mod go.sum ./  
# RUN go mod download  

# # Копируем остальную часть приложения  
# COPY . .  

# # Собираем бинарный файл  
# RUN go build -o main .  

# # Используем легкий образ для выполнения приложения  
# FROM alpine:latest  

# # Устанавливаем рабочую директорию  
# WORKDIR /root/  

# # Копируем скомпилированный бинарный файл из предыдущего этапа  
# COPY --from=builder /app/main .  

# # Открываем порт для приложения  
# EXPOSE 8080  

# # Команда запуска приложения  
# CMD ["./main"]
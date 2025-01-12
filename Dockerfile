FROM golang:latest AS builder  
WORKDIR /postgresApi  
COPY go.mod go.sum ./  
RUN go mod download  
COPY ./ ./  
RUN go build -o main .  

EXPOSE 8080  
CMD [ "./main" ]  

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
# FROM golang:latest AS builder  
# WORKDIR /postgresApi  
# COPY go.mod go.sum ./  
# RUN go mod download  
# COPY . .  
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .
# FROM alpine:3.13
# RUN apk update && apk --no-cache add bash
# COPY --from=builder /postgresApi/main /bin/  
# # COPY --from=builder /postgresApi/wait-for-it.sh /bin/wait-for-it.sh 
# EXPOSE 8080
# # RUN chmod +x /bin/wait-for-it.sh /bin/main
# # CMD [ "/bin/wait-for-it.sh", "localhost:5432", "--timeout=30", "--", "/bin/main" ]
# CMD [ "./main" ]

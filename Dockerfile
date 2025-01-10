# FROM golang:latest AS builder  
# WORKDIR /postgresApi  
# COPY go.mod go.sum ./  
# RUN go mod download  
# COPY ./ ./  
# RUN go build -o main .  

# EXPOSE 8080  
# CMD [ "./main" ]  
FROM golang:latest AS builder
WORKDIR /postgresApi 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:3.13
RUN apk update && apk --no-cache add bash
COPY --from=builder /postgresApi /bin/.
RUN ["chmod", "+x", "/bin/wait-for-it.sh"]

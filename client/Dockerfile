FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o client_main client_main.go

EXPOSE 8081
CMD ["./client_main"]

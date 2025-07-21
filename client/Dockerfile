# Dockerfile.client
FROM golang:latest

WORKDIR /app
COPY . .
RUN go build -o client client_main.go

CMD ["./client"]

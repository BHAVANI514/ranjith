# client/Dockerfile.client
FROM golang:1.22-alpine
WORKDIR /app
COPY client_main.go .
COPY ../go.mod .
RUN go build -o client_main client_main.go
CMD ["./client_main"]

# Dockerfile.server
FROM golang:1.22-alpine
WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go build -o main main.go
EXPOSE 8081
CMD ["./main"]

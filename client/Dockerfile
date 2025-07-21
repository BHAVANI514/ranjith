FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY client/client_main.go ./client_main.go
# Remove any problematic COPY for `definition` unless needed

RUN go mod download
RUN go build -o main main.go
RUN go build -o client_main client_main.go

EXPOSE 8081
CMD ["./main"]

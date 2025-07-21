FROM golang:1.22-alpine

WORKDIR /app

COPY . .         # module definition 

RUN go build -o main main.go

EXPOSE 8081
CMD ["./main"]

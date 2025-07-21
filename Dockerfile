FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .         # module definition 
COPY go.sum .         # ✅ required for dependency integrity
COPY main.go .        # your main file

RUN go mod tidy       # optional, ensures deps are clean
RUN go build -o main main.go

EXPOSE 8081
CMD ["./main"]

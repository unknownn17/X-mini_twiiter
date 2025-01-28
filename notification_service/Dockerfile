FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY . .
RUN go build -o main cmd/main.go

CMD ["./main"]
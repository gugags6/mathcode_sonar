FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN go mod init teste

COPY . .

RUN go build -o math

CMD ["./math"]
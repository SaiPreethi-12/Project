# Build Stage
FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

# Run Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
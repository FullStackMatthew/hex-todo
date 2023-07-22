# build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

RUN apk add --update git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/hex-todo

# final stage
FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/hex-todo .

# Remove the line for copying migrations, as we don't need it in the main.go approach.

CMD ["./hex-todo"]

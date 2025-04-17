FROM golang:1.24-alpine AS builder
LABEL authors="paullima"

WORKDIR /app

COPY . .

RUN go build -o loadtester .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/loadtester .

ENTRYPOINT ["/app/loadtester"]

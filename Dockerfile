FROM golang:1.20 AS builder

WORKDIR /src
COPY . .
RUN go build -o=census3 -ldflags="-s -w" ./cmd/census3

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/census3 /app/census3
ENTRYPOINT ["/app/census3"]


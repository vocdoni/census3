FROM golang:1.18 AS builder

WORKDIR /src
COPY . .
RUN go build -o=tokenscan -ldflags="-s -w" ./cmd/tokenscan

FROM debian:11.6-slim

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/tokenscan /app/tokenscan
ENTRYPOINT ["/app/tokenscan"]


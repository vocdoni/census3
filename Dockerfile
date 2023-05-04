FROM golang:1.20 AS builder

ARG SQLC_MIGRATIONS
ENV SQLC_MIGRATIONS=${SQLC_MIGRATIONS}

WORKDIR /src
COPY . .
RUN go mod tidy && go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN sqlc generate -f ${SQLC_MIGRATIONS}
RUN go build -o=census3 -ldflags="-s -w" ./cmd/census3

FROM debian:11.6-slim

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/census3 /app/census3
ENTRYPOINT ["/app/census3"]


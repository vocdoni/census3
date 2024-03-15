FROM golang:1.22.1 AS builder

WORKDIR /src
COPY . .

# Purge Go cache to ensure fresh dependency resolution
RUN go clean -modcache
RUN go clean -cache

RUN --mount=type=cache,sharing=locked,id=gomod,target=/go/pkg/mod/cache \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod/cache go build -o=census3 -ldflags="-s -w  -X=github.com/vocdoni/census3/internal.Version=$(git describe --always --tags --dirty --match='v[0-9]*')" ./cmd/census3

FROM debian:bookworm-slim AS base

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/census3 /app/census3
COPY --from=builder /src/db/initial_data/tokens.json /app/tokens.json

ENTRYPOINT ["/app/census3"]


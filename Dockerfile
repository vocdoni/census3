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
RUN ls /go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/

FROM debian:bookworm-slim AS base

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/census3 /app/census3
COPY --from=builder /src/db/initial_data/tokens.json /app/tokens.json

# Support for go-rapidsnark witness calculator (https://github.com/iden3/go-rapidsnark/tree/main/witness)
COPY --from=builder /go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/libwasmer.so \
                    /go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/libwasmer.so
# Support for go-rapidsnark prover (https://github.com/iden3/go-rapidsnark/tree/main/prover)
RUN apt-get update && \
	apt-get install --no-install-recommends -y libc6-dev libomp-dev openmpi-common libgomp1 curl && \
	apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/app/census3"]


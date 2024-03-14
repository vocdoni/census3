FROM golang:1.22.1 AS builder

WORKDIR /src
ENV CGO_ENABLED=1
RUN go env -w GOCACHE=/go-cache
COPY . .
RUN --mount=type=cache,target=/go-cache go mod download
RUN --mount=type=cache,target=/go-cache go build -o=census3 -ldflags="-s -w  -X=github.com/vocdoni/census3/internal.Version=$(git describe --always --tags --dirty --match='v[0-9]*')" ./cmd/census3

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


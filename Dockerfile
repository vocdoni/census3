FROM golang:1.21.1 AS builder

WORKDIR /src
COPY . .
RUN go build -o=census3 -ldflags="-s -w" ./cmd/census3

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/census3 /app/census3

# Support for go-rapidsnark witness calculator (https://github.com/iden3/go-rapidsnark/tree/main/witness)
COPY --from=builder /go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/libwasmer.so \
                    /go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/libwasmer.so
# Support for go-rapidsnark prover (https://github.com/iden3/go-rapidsnark/tree/main/prover)
RUN apt-get update && \
	apt-get install -y libc6-dev libomp-dev openmpi-common libgomp1 curl && \
	apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/app/census3"]


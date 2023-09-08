
# Vocdoni Census3

![census3-header](https://i.postimg.cc/HkgKdRYB/census3-header.png)

[![GoDoc](https://godoc.org/github.com/vocdoni/census3?status.svg)](https://godoc.org/github.com/vocdoni/census3)
[![Go Report Card](https://goreportcard.com/badge/github.com/vocdoni/census3)](https://goreportcard.com/report/github.com/vocdoni/census3)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

[![Join Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.gg/xFTh8Np2ga)
[![Twitter Follow](https://img.shields.io/twitter/follow/vocdoni.svg?style=social&label=Follow)](https://twitter.com/vocdoni)

<center>
<b>⚠ This project is currently a MVP and will be change ⚠</b>
</center>

---

## Description

Census3 is an API service to create censuses for elections with holders of a single token or a combination of a some of them. The service creates a list of holder addresses and balances and keeps it updated in realtime, for every registered token. Then, allows to create a merkle tree census (compatible with [Vocdoni](https://vocdoni.io/)) with these holders, using its balances as vote weight. 

#### Suported contract types
The service suports the following list of token types:
* ERC20
* ERC721
* ERC1155 (*coming soon*)
* ERC777
* Nation3 (veNation)
* wANT


#### About censuses
 - Census3 uses [go.vocdoni.io/dvote/tree/arbo](go.vocdoni.io/dvote/tree/arbo) to build the censuses merkle trees.
 - The censuses are published on [IPFS](https://ipfs.tech/) after its creation. 
 - The censuses can be created with the holders of just one token or a combination of some of them (*coming soon*).
 - The censuses are *zk-friendly* and can be used also for anonymous voting.


#### API Defintion
Check out the API endpoints definitions in the [`./api` folder](./api).

---

## Documentation

### How to run the Census3 API service

#### Using the CLI
1. Build the project:

```sh
go build -o census3 ./cmd/census3
```

2. Run the CLI:

```sh
./census3 --help
Usage of ./census3:
      --connectKey string      connect group key for IPFS connect
      --dataDir string         data directory for persistent storage (default "<$HOME>/.census3")
      --logLevel string        log level (debug, info, warn, error) (default "info")
      --port int               HTTP port for the API (default 7788)
      --web3Providers string   the list of URL's of available web3 providers (separated with commas)
```

Example:

```sh
./census3 --web3Providers https://goerli.rpcendpoint.io/,https://mainnet.rpcendpoint.io/ --logLevel debug

# or just:
# go run ./cmd/census3 --web3Providers https://goerli.rpcendpoint.io/,https://mainnet.rpcendpoint.io/ --logLevel debug
```

#### Using Docker

1. Create your config file using the [`.env` file](.env) as a template and save it the root.
```sh
# A web3 endpoint provider
WEB3_PROVIDERS=https://rpc-endoint.example1.com,https://rpc-endoint.example2.com

# Internal port for the service (80 and 443 are used by traefik)
PORT=7788

# Domain name for TLS
# DOMAIN=your.own.domain.xyz
DOMAIN=localhost

# Log level (info, debug, warning, error)
LOGLEVEL=debug

# IPFS connect key for discovering nodes
CONNECT_KEY=yourIPFSConnectKey
```

2. Run the services with `docker compose`:
```sh
docker compose up -d
```

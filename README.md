
# Vocdoni Census3

![census3-header](https://i.postimg.cc/HkgKdRYB/census3-header.png)

[![GoDoc](https://godoc.org/github.com/vocdoni/census3?status.svg)](https://godoc.org/github.com/vocdoni/census3)
[![Go Report Card](https://goreportcard.com/badge/github.com/vocdoni/census3)](https://goreportcard.com/report/github.com/vocdoni/census3)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

[![Join Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.gg/xFTh8Np2ga)
[![Twitter Follow](https://img.shields.io/twitter/follow/vocdoni.svg?style=social&label=Follow)](https://twitter.com/vocdoni)

<center>
<b>⚠ This project is currently a MVP and will be changed ⚠</b>
</center>

---

## Description

Census3 is an API service to create censuses for elections with holders of a single token or a combination of some of them. The service creates a list of holder addresses and balances and keeps it updated in realtime, for every registered token. Then, allows to create a merkle tree census (compatible with [Vocdoni](https://vocdoni.io/)) with these holders, using its balances as vote weight. 

#### Supported contract types
The service supports the following list of token types:
* ERC20
* ERC721
* ERC1155 (*coming soon*)
* ERC777
* Nation3 (veNation)
* wANT


#### About censuses
 * The censuses are published on [IPFS](https://ipfs.tech/) after its creation. 
 * Census3 uses [go.vocdoni.io/dvote/tree/arbo](go.vocdoni.io/dvote/tree/arbo) to build the censuses merkle trees.
 * The censuses can be created with the holders of just one token or a combination of tokens, using **complex strategies**  (*coming soon*).
 * The censuses are *zk-friendly* and can be used also for anonymous voting.

#### About complex strategies
A strategy is a definition of a group of previously created tokens and how their scanned holders must be combined to create a census.
* Must support combinations of tokens which contains:
  * An operator, which is a function associated with a tag (e.g. `AND`) that are used to combine token holders and define how to combine them.
  * Two token symbols (e.g. `BTC`), that identifies the token holders to combine.
  * Must have the following format: `<token_symbol> <operator> <token_symbol>`, e.g. `BTC OR ETH`.
* Must support groups of combinations, e.g. `USDC AND (ETH OR (BTC AND DAI))`

---

## Documentation

1. [How to run the Census3 service](#how-to-run-the-census3-api-service)
2. [Basic example](#basic-example)
3. [API definition](#api-defintion)


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

1. Create your config file using the [`.env` file](example.env) as a template and save it the root.
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

### Basic example

0. Starts the API service on `localhost:7788` with a web3 provider for `mainnet`
1. Register a new `erc20` token from `mainnet (chainId: 1)` by its contract address:

```sh
curl -X POST \
      --json '{"ID": "0xFE67A4450907459c3e1FFf623aA927dD4e28c67a", "type": "erc20", "chainID": 1}' \
      http://localhost:7788/api/tokens
```

2. Wait to that the API service completes the token sync. It could take up to 10-20 minutes, even more, based on the number of holders and transactions. You can check the token sync status getting the token info:
```sh
curl -X GET \
      http://localhost:7788/api/tokens/0xFE67A4450907459c3e1FFf623aA927dD4e28c67a
```

3. When the API ends, and the token reaches `synced` status (`token.status.synced = true`), it's time to create a new census based on the default token strategy. This strategy is created during the token registration and just contains the holders of this token. To create the census with token holders, you need to know the `token.defaultStrategy` (from token info endpoint):
```sh
curl -X POST \
        --json '{"strategyID": <strategyId>, "anonymous": true}" \
        http://localhost:7788/api/censuses
```
4. The last request will return a `queueId` which identifies the census creation and publication processes on the API queue. It will be completed in background. We can check if the task is done, it raised an error or was successfully completed:
```sh
curl -X GET \
        http://localhost:7788/censuses/queue/<queueId>
```

You can check and run the example using the [`example.sh` file](./example.sh):
```sh
sh ./example.sh

-> creating token...
Ok
-> created, waiting 4m to token scan
-> getting token info...
{"id":"0xFE67A4450907459c3e1FFf623aA927dD4e28c67a","type":"erc20","decimals":18,"startBlock":16976695,"symbol":"NEXT","totalSupply":"1000000000000000000000000000","name":"Connext","status":{"atBlock":18092468,"synced":true,"progress":100},"size":644,"defaultStrategy":1,"chainID":1}
-> enter the strategyId:
1
-> creating census...
{"queueId":"cd234ba75988e04e1e7a3234e48ff4033633142f"}
-> waiting 1m to census publication
-> enter the enqueue census:
cd234ba75988e04e1e7a3234e48ff4033633142f
{"done":true,"error":null,"census":{"censusId":1,"strategyId":1,"merkleRoot":"73368af290f4d0dfcb25b12060184bb3e5ad4147c5e5949de6729800c3629509","uri":"ipfs://bafybeiehspu3xrpshzjcvexl52u756cwfjobcwjz7ol4as44zfpvnlchsu","size":644,"weight":"5180125781955736442164650279357953853238828163172892166520872906800","anonymous":true}}
```

### API Definition
Check out the API endpoints definitions, accepted requests and expected responses in the [`./api` folder](./api).

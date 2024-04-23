<p align="center" width="100%">
    <img src="https://developer.vocdoni.io/img/vocdoni_logotype_full_white.svg" />
</p>

<p align="center" width="100%">
    <a href="https://github.com/vocdoni/census3/commits/main/"><img src="https://img.shields.io/github/commit-activity/m/vocdoni/census3" /></a>
    <a href="https://github.com/vocdoni/census3/issues"><img src="https://img.shields.io/github/issues/vocdoni/census3" /></a>
    <a href="https://github.com/vocdoni/census3/actions/workflows/main.yml/"><img src="https://github.com/vocdoni/census3/actions/workflows/main.yml/badge.svg" /></a>
    <a href="https://pkg.go.dev/github.com/vocdoni/census3"><img src="https://godoc.org/go.vocdoni.io/census3?status.svg"></a>
    <a href="https://discord.gg/xFTh8Np2ga"><img src="https://img.shields.io/badge/discord-join%20chat-blue.svg" /></a>
    <a href="https://twitter.com/vocdoni"><img src="https://img.shields.io/twitter/follow/vocdoni.svg?style=social&label=Follow" /></a>
</p>


  <div align="center">
    Vocdoni is the first universally verifiable, censorship-resistant, anonymous, and self-sovereign governance protocol. <br />
    Our main aim is a trustless voting system where anyone can speak their voice and where everything is auditable. <br />
    We are engineering building blocks for a permissionless, private and censorship resistant democracy.
    <br />
    <a href="https://developer.vocdoni.io/"><strong>Explore the developer portal »</strong></a>
    <br />
    <h3>More About Us</h3>
    <a href="https://vocdoni.io">Vocdoni Website</a>
    |
    <a href="https://vocdoni.app">Web Application</a>
    |
    <a href="https://explorer.vote/">Blockchain Explorer</a>
    |
    <a href="https://law.mit.edu/pub/remotevotingintheageofcryptography/release/1">MIT Law Publication</a>
    |
    <a href="https://chat.vocdoni.io">Contact Us</a>
    <br />
    <h3>Key Repositories</h3>
    <a href="https://github.com/vocdoni/vocdoni-node">Vocdoni Node</a>
    |
    <a href="https://github.com/vocdoni/vocdoni-sdk/">Vocdoni SDK</a>
    |
    <a href="https://github.com/vocdoni/ui-components">UI Components</a>
    |
    <a href="https://github.com/vocdoni/ui-scaffold">Application UI</a>
    |
    <a href="https://github.com/vocdoni/census3">Census3</a>
  </div>

# census3

Census3 is an API service which facilitates the creation of Vocdoni censuses whose eligible voters are defined by token-holders of some cryptocurrency token(s). The service has a list of registered tokens whose holder addresses and balances it keeps updated in real time. 

The Census3 service then uses this token-holder information to create a merkle tree census (compatible with [Vocdoni](https://vocdoni.io/)) according to some given eligibility criteria. 

This code is written in golang and is meant to be used in conjunction with other Vocdoni tools, such as the [API](https://developer.vocdoni.io/vocdoni-api/vocdoni-api). 

The best place to learn about interacting with the Census3 Service is the [developer portal](https://developer.vocdoni.io/).

### Table of Contents
- [Getting Started](#getting-started)
- [Reference](#reference)
- [Examples](#examples)
- [Preview](#preview)
- [Disclaimer](#disclaimer)
- [Contributing](#contributing)
- [License](#license)


## Getting Started

You can run the Census3 service locally for testing or deploy it yourself.

#### Using the CLI

You can locally host your own Census3 service by first building this code:
```sh
go build -o census3 ./cmd/census3
```

Then you can run the binary with the following possible options:
```sh
./census3 --help
Usage of ./census3:
      --adminToken string          the admin UUID token for the API
      --connectKey string          connect group key for IPFS connect
      --dataDir string             data directory for persistent storage (default "~/.census3")
      --logLevel string            log level (debug, info, warn, error) (default "info")
      --port int                   HTTP port for the API (default 7788)
      --scannerCoolDown duration   the time to wait before next scanner iteration (default 2m0s)
      --initialTokens string       path of the initial tokens json file
      --web3Providers string       the list of URLs of available web3 providers
      --gitcoinEndpoint string     Gitcoin Passport API access token
      --gitcoinCooldown duration   Gitcoin Passport API cooldown (default 6h0m0s)
      --poapAPIEndpoint string     POAP API endpoint
      --poapAuthToken string       POAP API access token
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
# Domain name for TLS
DOMAIN=your.own.domain.xyz
# A web3 endpoint provider
CENSUS3_WEB3PROVIDERS="https://rpc-endoint.example1.com,https://rpc-endoint.example2.com"
# Internal port for the service (80 and 443 are used by traefik)
CENSUS3_PORT=7788
# Log level (info, debug, warning, error)
CENSUS3_LOGLEVEL="debug"
# IPFS connect key for discovering nodes
# CONNECT_KEY=yourIPFSConnectKey
CENSUS3_CONNECTKEY="census3key"
# Internal data folder
CENSUS3_DATADIR="./.census3"
# Scanner cooldown duration
CENSUS3_SCANNERCOOLDOWN="20s"
# UUID admin token for protected endpoints
CENSUS3_ADMINTOKENS="UUID"
# POAP API configuration
CENSUS3_POAPAPIENDPOINT="poapAPIEndpoint"
CENSUS3_POAPAUTHTOKEN="yourPOAPAuthToken"
# Gitcoin API configuration
CENSUS3_GITCOINAPIENDPOINT="gitcoinAPIEndpoint"
CENSUS3_INITIALTOKENS="/app/initial_tokens.json"
```

2. Run the services with `docker compose`:
```sh
docker compose up -d
```

## Reference

The Census3 usage and API is documented at the [developer portal](https://developer.vocdoni.io/protocol/census/on-chain/census3#api-defintion). We recommend reading this documentation before trying to run your own examples.

## Examples

The following shows a basic example for creating a Census3 census using a locally-running instance:

0. Start the API service on `localhost:7788` with a web3 provider for `mainnet`
1. Register a new `erc20` token from `mainnet (chainId: 1)` by its contract address:

```sh
curl -X POST \
      --json '{"ID": "0xFE67A4450907459c3e1FFf623aA927dD4e28c67a", "type": "erc20", "chainID": 1}' \
      http://localhost:7788/api/tokens
```

1. Wait for the API service to complete the token synchronization. It could take up to 10-20 minutes, even more, based on the number of holders and transactions (for this reason, high-traffic tokens like ETH are infeasible). You can check the token sync status by fetching the token info:
```sh
curl -X GET \
      http://localhost:7788/api/tokens/0xFE67A4450907459c3e1FFf623aA927dD4e28c67a
```

1. When the token reaches `synced` status (`token.status.synced = true`), you can create a new census based on the default token strategy. This strategy is created during the token registration and just contains the holders of the token. To create the census with token holders, you need to know the `token.defaultStrategy` (from the token info endpoint):
```sh
curl -X POST \
        --json '{"strategyID": <strategyId>, "anonymous": true}" \
        http://localhost:7788/api/censuses
```
1. The previous request will return a `queueId` which identifies the census creation and publication processes on the API queue. Census publication will be completed in the background. We can check if the task is done and whether it raised an error or was successfully completed:
```sh
curl -X GET \
        http://localhost:7788/censuses/queue/<queueId>
```

You can also automatically run this example using the [`example.sh` file](./example.sh):
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

## Preview

You can see the Census3 API in use with the Vocdoni SDK at the SDK's [census3 example](https://github.com/vocdoni/vocdoni-sdk/tree/examples/token-based/examples/token-based)

## Disclaimer

The Census3 service is WIP. Please beware that it can be broken at any time if the release is `alpha` or `beta`. We encourage you to review this repository and the developer portal for any changes.

## Contributing 

While we welcome contributions from the community, we do not track all of our issues on Github and we may not have the resources to onboard developers and review complex pull requests. That being said, there are multiple ways you can get involved with the project. 

Please review our [development guidelines](https://developer.vocdoni.io/development-guidelines).

## License

This repository is licensed under the [GNU Affero General Public License v3.0.](./LICENSE)


    Vocdoni Census3 Service
    Copyright (C) 2024 Vocdoni Association

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

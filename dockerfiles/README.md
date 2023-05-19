# Dockerfiles

## Deploy `census3` on production

It includes Traefik proxy to redirect the HTTP calls and Watchtower to keep the container image updated. Steps:

1. Go to `./census3`
2. Rename `.env.example` to `.env` and modify it with your configuration.
3. Run `docker-compose up`

## Run `census3` for testing

It includes Traefik proxy to redirect the HTTP calls and Foundry Anvil to start a local ethereum testnet based on a hardfork
of the Web3 endpoint defined.

1. Go to `./testsuite`
2. Rename `.env.example` to `.env` and modify it with your configuration.
3. Run `docker-compose up`
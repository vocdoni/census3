# Endpoints

## GET /api

### Supported contracts

- Request:
    `GET /supportedContracts`
- Response:
    `{"ok": true, "contracts": ["a", "b", ...]}`

### Add contract

- Request:
    `GET /addContract/<contractAddress>/<contractType>/<startBlock>`
- Response:
    `{"ok": true}`

### List contracts

- Request:
    `GET /listContracts`
- Response:
    `{"ok": true, "contracts": ["a", "b", ...]}`

### Get contract

- Request:
    `GET /getContract/<contractAddress>`
- Response:
    `{"ok": true, "token": {"name": "tokenName","address": "0x123","type": "erc20","symbol": "tn","decimals": 18,"totalSupply": "10000000","startBlock": 1000,"lastBlock": 9000,"lastRoot": "0x123","lastSnapshot": 5}}`

### Get balances

- Request:
    `GET /balances/<contractAddress>`
- Response:
    `{"ok": true, "balances": {"0x123": 100000, "0x456": 200000, ...}}`

### Rescan contract

- Request:
    `GET /rescan/<contractAddress>`
- Response:
    `{"ok": true}`

### Get latest root

- Request:
    `GET /root/<contractAddress>`
- Response:
    `{"ok": true, "root": "0x123"}`

### Get root at block

- Request:
    `GET /root/<contractAddress>/<blockNumber>`
- Response:
    `{"ok": true, "root": "0x123"}`

### Queue export

- Request:
    `GET /queryExport/<contractAddress>`
- Response:
    `{"ok": true, "block": "123"}`

### Fetch export

- Request:
    `GET /fetchExport/<contractAddress>/<blockNumber>`
- Response:
    `{"ok": true, "data": [1,2,3,4,a,b,c,d,...]}`

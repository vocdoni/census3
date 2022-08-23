# Vocdoni Token State

**This is WIP**

The aim for this tool is to provide a GoLang library and a HTTP/API service that can be used to fetch an updated list of token holders for a Ethereum token (just ERC20 at the beggining). It must allow to get the holders list, amounts on a specific block.

The token holders list and the amounts are stored into a zkSnarks friendly Merkle Tree (currently Arbo). 

---

### Examples

Start the token scan service, the API will listen at port 7788 HTTP. 
```bash
go run ./cmd/tokenscan --url=wss://mainnet.infura.io/ws/v3/INFURA_TOKEN --port=7788
```


Add a new contract (Aragon Network Token) and an initial block number to start the scan.
```bash
curl 127.0.0.1:7788/api/addContract/0xa117000000f279D81A1D3cc75430fAA017FA5A2e/11000000
```

The service will start scaning the token transfers from the initial block provided.

The Balances can be fetched by calling `balances` method.
```bash
curl 127.0.0.1:7788/api/balances/0xa117000000f279D81A1D3cc75430fAA017FA5A2e
```


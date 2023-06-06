
# Vocdoni Census3

![census3-header](https://i.postimg.cc/HkgKdRYB/census3-header.png)

**⚠ This project is currently a MVP and will be change ⚠**

The aim for this tool is to provide a Golang library and a HTTP/API service that can be used to fetch an updated list of token holders for a Ethereum token.
The list of supported tokens is:

- ERC20
- ERC721
- ERC1155
- ERC777
- Nation3 (veNation)
- wANT

It makes it possible to obtain an updated list of holders and their balances. In the future, it will be possible to obtain this list for a specific block.

The list of token holders and their balances is stored in a zkSnarks friendly Merkle Tree (currently Arbo) and is shared using IPFS.

---

### API Defintion
Check out the API endpoints definitions in the [`./api` folder](./api).

### Run the API service

Start the token scan service, the API will listen at port 7788 HTTP: 

**Using Go**
```bash
go run ./cmd/tokenscan --url=wss://mainnet.infura.io/ws/v3/INFURA_TOKEN --port=7788
```

**Using Docker**

Edit [`.env` file](.env) with your own information.
```bash
docker-compose up -d
```

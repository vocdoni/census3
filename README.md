# Token State

**This is WIP**

The aim for this tool is to provide a GoLang library and a HTTP/API service that can be used to fetch an updated list of token holders for a Ethereum token (just ERC20 at the beggining). It must allow to get the holders list, amounts on a specific block.

The token holders list and the amounts are stored into a MerkleTree data storage (currently Graviton, but any other might be used). To this end, the tool will provide a MerkleRoot for a Token+Block and any MerkleProof of a token holder. The Proof can be used anywhere offchain to demonstrate the tokens holding and the Root can be added to any blockchain in order to validate it.

### Examples

Scan for the Aragon contract

```
go run ./cmd/tokenscan -url=wss://mainnet.infura.io/ws/v3/INFURA_TOKEN -contract=0x960b236A07cf122663c4303350609A66A7B288C0 -from=3000000 
```


# Entity Bridge

This pkg scans an ERC20 token on any EVM compatible chain and creates a Vocdoni entity on another EVM compatible chain.
The stored entity atributes are based on the data scrapped from the ERC20 token.

### Example

Create TMOON entity on xDAI:

```go run entitybridge.go --tokenContract 0x106c8eBaD6D9A71c962Da4088721221de9BD4fB7 --registryContract 0x00cEBf9E1E81D3CC17fbA0a49306EBA77a8F26cD --resolverContract 0x80629aF85C5623fDFDD3744c3192824be72B06F6 --web3Home <xdai endpoint> --web3Foreign <xdai endpoint> --gatewayURL <gateway url> --ethSigner <privkey> --logLevel debug```

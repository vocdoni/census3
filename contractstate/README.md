# Contract State

This directory contains all the logic for fetching token information and token holders from an Ethereum contract.

- `const.go`: General constants used in the package. Also contains the list of supported tokens. If you want to support a new token, add it here.
- `state.go`: Contains the `State` struct, which is the main struct for storing token information and token holders.
- `web3.go`: Contains the `Web3` struct, which is the main struct for interacting with the Ethereum blockchain via the defined methods. If you want to add a new token that requires a different way of fetching the token holders, add the custom logic here.

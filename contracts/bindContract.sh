#!/usr/bin/env bash

set -euo pipefail

abi() {
  abigen --abi="$1" --pkg="$2" --out="$3"
}

abi './erc/erc20/ERC20.abi' ERC20Contract './erc/erc20/erc20.go'
abi './erc/erc721/ERC721.abi' ERC721Contract './erc/erc721/erc721.go'
# abi './erc/erc1155/ERC1155.abi' ERC1155Contract './erc/erc1155/erc1155.go'
abi './erc/erc777/ERC777.abi' ERC777Contract './erc/erc777/erc777.go'

# abi './nation3/vestedToken/veNation.abi' Nation3VestedTokenContract './nation3/vestedToken/veNation.go'
# abi './aragon/want/want.abi' AragonWrappedANTTokenContract './aragon/want/want.go'
# abi './poap/poap.abi' POAPContract './poap/poap.go'
abi './farcaster/keyRegistry/KeyRegistry.abi' FarcasterKeyRegistry './farcaster/keyRegistry/keyregistry.go'
abi './farcaster/idRegistry/IDRegistry.abi' FarcasterIDRegistry './farcaster/idRegistry/idregistry.go'
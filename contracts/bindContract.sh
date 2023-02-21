#!/usr/bin/env sh

set -euo pipefail

abi() {
  abigen --abi="$1" --pkg="$2" --out="$3"
}

abi './openzeppelin/erc20/ERC20.abi' ERC20Contract './openzeppelin/erc20/erc20.go'
abi './openzeppelin/erc721/ERC721.abi' ERC721Contract './openzeppelin/erc721/erc721.go'
abi './openzeppelin/erc1155/ERC1155.abi' ERC1155Contract './openzeppelin/erc1155/erc1155.go'
abi './openzeppelin/erc777/ERC777.abi' ERC777Contract './openzeppelin/erc777/erc777.go'

abi './nation3/vestedToken/veNation.abi' Nation3VestedTokenContract './nation3/vestedToken/veNation.go'
abi './aragon/want/want.abi' AragonWrappedANTTokenContract './aragon/want/want.go'

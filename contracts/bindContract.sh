#!/bin/sh

abigen --abi=./openzeppelin/erc20/ERC20.abi --pkg=ERC20Contract --out=./openzeppelin/erc20/erc20.go
abigen --abi=./openzeppelin/erc721/ERC721.abi --pkg=ERC721Contract --out=./openzeppelin/erc721/erc721.go
abigen --abi=./openzeppelin/erc1155/ERC1155.abi --pkg=ERC1155Contract --out=./openzeppelin/erc1155/erc1155.go
abigen --abi=./openzeppelin/erc777/ERC777.abi --pkg=ERC777Contract --out=./openzeppelin/erc777/erc777.go

abigen --abi=./nation3/vestedToken/veNation.abi --pkg=Nation3VestedTokenContract --out=./nation3/vestedToken/veNation.go

abigen --abi=./aragon/want/want.abi --pkg=AragonWrappedANTTokenContract --out=./aragon/want/want.go
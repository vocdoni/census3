package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	rpcURL          string
	contractAddress string
	blockNumber     int64
)

func init() {
	flag.StringVar(&rpcURL, "rpc", "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID", "Ethereum RPC URL")
	flag.StringVar(&contractAddress, "contract", "", "ERC20 contract address")
	flag.Int64Var(&blockNumber, "block", 0, "Block number")
}

var transferEventSignature = []byte("Transfer(address,address,uint256)")
var transferEventSigHash = common.BytesToHash(crypto.Keccak256(transferEventSignature))

func main() {
	flag.Parse()

	if rpcURL == "" || contractAddress == "" || blockNumber == 0 {
		log.Fatalf("All flags (rpc, contract, block) are required")
	}
	fmt.Println("event signature hash:", transferEventSigHash.Hex())

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	blockNum := big.NewInt(blockNumber)
	contractAddr := common.HexToAddress(contractAddress)

	query := ethereum.FilterQuery{
		FromBlock: blockNum,
		ToBlock:   blockNum,
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{{transferEventSigHash}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to retrieve logs: %v", err)
	}

	transferEventABI := `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

	contractABI, err := abi.JSON(strings.NewReader(transferEventABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	for _, vLog := range logs {
		event := struct {
			From  common.Address
			To    common.Address
			Value *big.Int
		}{}

		err := contractABI.UnpackIntoInterface(&event, "Transfer", vLog.Data)
		if err != nil {
			log.Fatalf("Failed to unpack log data: %v", err)
		}

		event.From = common.HexToAddress(vLog.Topics[1].Hex())
		event.To = common.HexToAddress(vLog.Topics[2].Hex())

		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Printf("From: %s\n", event.From.Hex())
		fmt.Printf("To: %s\n", event.To.Hex())
		fmt.Printf("Value: %s\n", event.Value.String())
	}
}

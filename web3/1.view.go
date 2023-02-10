package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	Number string
}

func main() {
	client, err := rpc.Dial("https://mainnet.infura.io/v3/f76d1d451ebb4b1b8c4eb48377a7454a")
	if err != nil {
		log.Fatalf("Could not connect to Infura: %v", err)
	}

	var lastBlock Block
	err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("Cannot get the latest block:", err)
		return
	}

	fmt.Printf("Latest block: %v\n", lastBlock.Number)
}

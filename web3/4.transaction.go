package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func openTx1(tx *types.Transaction) {

	from := ""
	if msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), big.NewInt(1)); err == nil {
		//fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		from = msg.From().Hex()

	} else {
		fmt.Println(err.Error())
		fmt.Println("fuck")
	}

	fmt.Println(fmt.Sprintf("from %s to %s value %s ", from, tx.To(), tx.Value().String()))
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/f76d1d451ebb4b1b8c4eb48377a7454a")

	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(16591801)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		openTx1(tx)
	}

	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)       // false
}

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func openTx(tx *types.Transaction) {

	from := ""
	if msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), big.NewInt(1)); err == nil {
		fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
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

	txHash := common.HexToHash("0x17a5deb15b958684836fd80268ae005c8457ffe135a4b1201a633bd554ff3e5e")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	openTx(tx)

	//if msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), big.NewInt(1)); err == nil {
	//	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
	//
	//} else {
	//	fmt.Println(err.Error())
	//	fmt.Println("fuck")
	//}
	//
	//fmt.Println(tx.To())
	//fmt.Println(tx.Data())
	//
	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)       // false
}

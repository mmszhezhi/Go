package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/f76d1d451ebb4b1b8c4eb48377a7454a")
	if err != nil {
		log.Fatal(err)
	}

	//account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	account := common.HexToAddress("0xc12D36790612C8e093fCD20a3f0454F9637c0d29")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	//blockNumber := big.NewInt(5532993)
	//balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}

package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/uaGUMcMAZhXm16k8nb8b")
	if err != nil {
		log.Fatal("Whoops something went wrong!", err)
	}

	ctx := context.Background()
	tx, err := conn.TransactionReceipt(ctx, common.HexToHash("0xd336078cca5c3235775e95161524c631623be4446c3b4e9d8d744268bfdeece4"))
	if (err == nil) {
		fmt.Println(tx)
	}

	bal, err := conn.BalanceAt(ctx, common.HexToAddress("0xa1201A92B56CB45860808c1F2fDe9f482cCD2790"), nil)
	if (err == nil) {
		fmt.Println(bal , "wei")
	}
}
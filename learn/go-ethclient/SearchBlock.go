package go_ethclient

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// 2.查询区块信息
func SearchBlock(client *ethclient.Client) {
	number := big.NewInt(5671744)

	header1, err1 := client.HeaderByNumber(context.Background(), number)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("5671744区块高度", header1.Number.Uint64())
	fmt.Println("5671744区块高度", header1.Time)
	fmt.Println("5671744区块高度", header1.Difficulty.Uint64())
	fmt.Println("5671744区块高度", header1.Hash().Hex())

	//最新区块头信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	fmt.Println("最新区块高度", header.Number.Uint64())     // 9564109
	fmt.Println("最新区块高度", header.Time)                // 1712798400
	fmt.Println("最新区块高度", header.Difficulty.Uint64()) // 0
	fmt.Println("最新区块高度", header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}
	//获取完整区块数据
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(len(block.Transactions()))   // 70
	//交易数量 指定区块的交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("交易数量", count) // 70

}

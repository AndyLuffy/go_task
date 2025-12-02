package go_ethclient

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeBlock(client *ethclient.Client) {

	headers := make(chan *types.Header)
	client.SubscribeNewHead(context.Background(), headers)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {

		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("区块哈希", header.Hash().Hex())

			fmt.Println("区块头", header.Hash().Hex())
			fmt.Println("区块高度", header.Number.Uint64())
			fmt.Println("时间", header.Time)
			fmt.Println("nonce", header.Nonce.Uint64())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("区块哈希", block.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())    // 3477413
			fmt.Println(block.Time())               // 1529525947
			fmt.Println(block.Nonce())              // 130524141876765836
			fmt.Println(len(block.Transactions()))  // 7
		}
	}
}

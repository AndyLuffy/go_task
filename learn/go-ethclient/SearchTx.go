package go_ethclient

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 查询交易详细信息
func SearchTx(client *ethclient.Client) {

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	fmt.Println("block hash", block.Hash().Hex())
	//0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("交易哈希", tx.Hash().Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println("交易转账value", tx.Value().String()) // 100000000000000000
		fmt.Println(tx.Gas())                         // 21000
		fmt.Println(tx.GasPrice().Uint64())           // 100000000000
		fmt.Println("nonce", tx.Nonce())              // 245132
		fmt.Println("input", tx.Data())               // []
		fmt.Println("交易接收方", tx.To().Hex())           // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		//通过交易获取还原交易sender
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}
		//通过交易查询交易的收据数据来判断交易是否成功
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		fmt.Println("isPending", isPending)

		fmt.Println(receipt.Status) // 1
		//事件
		fmt.Println(receipt.Logs) // []
		break
	}

	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	//获取交易的完整数据
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false

}

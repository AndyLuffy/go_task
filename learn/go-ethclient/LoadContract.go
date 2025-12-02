package go_ethclient

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/learn/go_task/learn/go-ethclient/store"
)

const (
	contractAddr = "0xFbf9C9Daf1ec5a783aB70965054f4bD80922F0F5"
)

func LoadContract(client *ethclient.Client) {

	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract
}

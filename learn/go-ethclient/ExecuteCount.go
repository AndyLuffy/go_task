package go_ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/learn/go_task/learn/go-ethclient/count"
)

const countAddress = "0x2e89eDd129D4D37237726737E944a779A11246e7"

func ExexuteCount(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {

	countContract, err := count.NewCount(common.HexToAddress(countAddress), client)

	if err != nil {
		log.Fatal("Failed to instantiate contract:", err)
	}

	/*opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := countContract.Add(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())*/

	value, err := countContract.Count(&bind.CallOpts{Context: context.Background()})
	fmt.Println("count:", value)

}

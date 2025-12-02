package go_ethclient

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/learn/go_task/learn/go-ethclient/store"
)

const (
	storeContractAddr = "0xFbf9C9Daf1ec5a783aB70965054f4bD80922F0F5"
)

func ExecuteContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {

	storeContract, err := store.NewStore(common.HexToAddress(storeContractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value1111122"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	//调用合约
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	//查询合约
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value in contract:", valueInContract)
	fmt.Println("value in contract:", hex.EncodeToString(valueInContract[:]))
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}

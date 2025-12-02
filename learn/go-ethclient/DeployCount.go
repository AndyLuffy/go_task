package go_ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/learn/go_task/learn/go-ethclient/count"
)

func DeployCount(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//交易nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		log.Fatal("cannot get nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("cannot get gasPrice")
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("cannot get chainId")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal("cannot get auth")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	address, tx, _, err := count.DeployCount(auth, client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

}

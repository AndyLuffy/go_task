package go_ethclient

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func CreateWallet() {
	//成随机私钥
	//privateKey, err := crypto.GenerateKey()

	//

	privateKey, err := crypto.HexToECDSA("***********")

	if err != nil {
		log.Fatal(err)
	}
	//将其转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥", hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)           //转byte
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("公钥address is:", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

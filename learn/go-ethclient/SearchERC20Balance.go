package go_ethclient

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/learn/go_task/learn/go-ethclient/erc20"
)

// 查询ERC20代币
func SearchErc20Balance(client *ethclient.Client) {

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")
	instance, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}

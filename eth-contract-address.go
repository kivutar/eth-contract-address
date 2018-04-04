package main

import (
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	address := flag.String("address", "", "Ethereum address of the account")
	nonce := flag.Uint64("nonce", 0, "Nonce of the account")
	flag.Parse()

	result := crypto.CreateAddress(common.HexToAddress(*address), *nonce)

	fmt.Println(result.String())
}

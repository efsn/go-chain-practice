package main

import (
	"fmt"
	"go-chain-practice/pkg/blockchain"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockChain()

	bc.AddBlock("send 1 btc to oho")
	bc.AddBlock("send 2 btc to oho")

	for _, block := range bc.Blocks {
		fmt.Printf("pre hash: %x\n", block.PreBlockHash)
		fmt.Printf("data: %x\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Valid()))
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"go-chain-practice/src/blockchain"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

func main() {
	bc := blockchain.NewBlockChain()

	defer func(DB *bolt.DB) {
		if err := DB.Close(); err != nil {
			log.Panic(err)
		}
	}(bc.DB)

	bc.AddBlock("send 1 btc to oho")
	bc.AddBlock("send 2 btc to oho")

	err := bc.ForEach(func(k []byte, v []byte) error {
		block := blockchain.Deserialize(v)
		fmt.Printf("pre hash: %x\n", block.PreBlockHash)
		fmt.Printf("data: %x\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Valid()))
		fmt.Println()

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

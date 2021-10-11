package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
	Nonce        int
}

func (r *Block) Serialize() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(r)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

func Deserialize(blockBytes []byte) *Block {
	block := &Block{}
	decoder := gob.NewDecoder(bytes.NewBuffer(blockBytes))
	if err := decoder.Decode(block); err != nil {
		log.Panic(err)
	}
	return block
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         []byte(data),
		PreBlockHash: preBlockHash,
		Hash:         []byte{},
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("genesis block", []byte{})
}

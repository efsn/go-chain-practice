package blockchain

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const (
	dbFile       = "blockchain.db"
	blocksBucket = "blocks"
)

type BlockChain struct {
	tip []byte
	DB  *bolt.DB
}

type ChainIterator struct {
	currentHash []byte
	DB          *bolt.DB
}

func (r *BlockChain) AddBlock(data string) {
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			log.Panic("no exist blockchain found")
		}

		// get last chain block
		lh := b.Get(b.Get([]byte("l")))
		lb := Deserialize(lh)
		next := NewBlock(data, lb.Hash)

		err := b.Put(next.Hash, next.Serialize())
		if err != nil {
			log.Panic(err)
		}

		if err := b.Put([]byte("l"), next.Hash); err != nil {
			log.Panic(err)
		}

		r.tip = next.Hash
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

func (r *BlockChain) lastBlock(preHash []byte) error {
	return r.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			log.Panic("no exist blockchain found")
		}

		return nil
	})
}

func (r *BlockChain) GetBlock(hash []byte) *Block {
	var block *Block
	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			log.Panicf("%s not exist", blocksBucket)
		}
		block = Deserialize(b.Get(hash))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}

func (r *BlockChain) ForEach(fn func(k []byte, v []byte) error) error {
	return r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			log.Panicf("%s not exist", blocksBucket)
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil && string(k) != "l"; k, v = c.Next() {
			if err := fn(k, v); err != nil {
				return err
			}
		}
		return nil
	})
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			fmt.Println("no exist blockchain found, create a new one")

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			lb := NewGenesisBlock()

			err = b.Put(lb.Hash, lb.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), lb.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = lb.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{tip, db}
}

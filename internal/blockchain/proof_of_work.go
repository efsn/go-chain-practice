package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var maxNonce = math.MaxInt64

const targetBits = 16

// ProofOfWork POW
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork say something
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

func (r *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			r.block.PreBlockHash,
			r.block.Data,
			IntToByte(r.block.Timestamp),
			IntToByte(targetBits),
			IntToByte(int64(nonce)),
		}, []byte{})
	return data
}

// Run calc POW
func (r *ProofOfWork) Run() (int, []byte) {
	var (
		hashInt big.Int
		hash    [32]byte
	)
	nonce := 0

	fmt.Printf("mining the block contain %s\n", r.block.Data)
	for nonce < maxNonce {
		data := r.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(r.target) < 0 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

// Valid check POW
func (r *ProofOfWork) Valid() bool {
	var hashInt big.Int
	data := r.prepareData(r.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(r.target) < 0
}

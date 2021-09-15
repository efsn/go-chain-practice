package blockchain

type BlockChain struct {
	Blocks []*Block
}

func (r *BlockChain) AddBlock(data string) {
	preBlock := r.Blocks[len(r.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	r.Blocks = append(r.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

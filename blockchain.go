package main

// Blockchain keeps a sequence of Blocks
type BlockChain struct {
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockChain creates a new BlockChain with genesis Block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

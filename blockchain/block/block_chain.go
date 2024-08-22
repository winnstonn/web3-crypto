package block

import (
	"time"
)

type Blockchain struct {
	genesisBlock *Block
	chain        []*Block
	difficulty   int
}

func (bc *Blockchain) CreateBlockchain(difficulty int) *Blockchain {
	genesisBlock := &Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return &Blockchain{
		genesisBlock,
		[]*Block{genesisBlock},
		difficulty,
	}
}

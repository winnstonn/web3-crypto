package block

import "time"

type Block struct {
	data      map[string]interface{}
	hash      string
	prevHash  string
	timestamp time.Time
	pow       int
}

type Blockchain struct {
	genesisBlock *Block
	chain        []*Block
	difficulty   int
}

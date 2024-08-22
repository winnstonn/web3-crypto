package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data      map[string]interface{}
	hash      string
	prevHash  string
	timestamp time.Time
	pow       int
}

func (b *Block) calculateHash() string {
	data, err := json.Marshal(b.data)
	if err != nil {
		fmt.Printf("Invalid JSON data on block, error: %v\n", err)
	}

	blockData := b.prevHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

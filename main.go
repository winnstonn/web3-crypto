package main

import (
	"fmt"

	"github.com/web3-crypto/blockchain/merkle"
)

func main() {
	fmt.Println("Welcome to Crypto World...")

	data := []string{"A", "B", "C", "D"}
	rootIter := merkle.BuildMerkleTreeIter(data)
	merkle.PrintTree(rootIter, "")
	fmt.Println("Root Iter Hash:", rootIter.Hash)
	fmt.Println()

	rootRec := merkle.BuildMerkleTreeRec(data)
	merkle.PrintTree(rootRec, "")
	fmt.Println("Root Rec Hash:", rootRec.Hash)
}

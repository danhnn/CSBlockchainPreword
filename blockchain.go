package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
	"errors"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	// Fill me in, noble warrior.
	genesisBlock := Block{"Genesis Block", time.Now().Unix(), []byte{}, []byte{}}
	Blockchain = []Block{genesisBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	newBlock := Block{data, time.Now().Unix(),oldBlock.Hash, []byte{} }
	hash := newBlock.calculateHash()
	newBlock.Hash = hash
	return newBlock
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	// Fill me in, brave wizard.
	lastBlock := Blockchain[len(Blockchain)-1]

	if (!bytes.Equal(lastBlock.Hash, b.PrevHash)) {
		return errors.New("invalid block")
	}

	aceptedHash := b.calculateHash()
	currentHash := b.Hash
	
	if (!bytes.Equal(aceptedHash, currentHash)) {
		return errors.New("invalid block")
	}
	
	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

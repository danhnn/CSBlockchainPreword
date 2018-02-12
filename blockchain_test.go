package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// This test just verifies that life is ok.
func TestTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}

func TestBlock(t *testing.T) {
	InitBlockchain()
	length := len(Blockchain)
	if length != 1 {
		t.Error("New Blockchain should have length 1, but has length: ", length)
	}
}

func TestNewBlock(t *testing.T) {
	testOldBlock := createTestBlock()
	testStr := "testing new block!"
	b := NewBlock(testOldBlock, testStr)
	if b.Data != "testing new block!" {
		t.Error("New Block should have data: ", testStr, "but has data: ", b.Data)
	}
	if bytes.Compare(b.PrevHash, testOldBlock.calculateHash()) != 0 {
		t.Errorf("New Block should have PrevHash: %x, but has PrevHash: %x",
			testOldBlock.calculateHash(), b.PrevHash)
	}
	if bytes.Compare(b.Hash, b.calculateHash()) != 0 {
		t.Errorf("New Block should have Hash: %x, but has Hash: %x",
			b.calculateHash(), b.Hash)
	}
}

func TestAddInvalidBlock(t *testing.T) {
	InitBlockchain()

	testOldBlock := createTestBlock()
	testStr := "testing new block!"
	b := NewBlock(testOldBlock, testStr)
	err := AddBlock(b)
	if err == nil {
		t.Error("Adding invalid block should have thrown an error.")
	}
}

func TestAddValidBlock(t *testing.T) {
	InitBlockchain()
	lastBlock := Blockchain[len(Blockchain)-1]
	testStr := "testing new block!"

	b := NewBlock(lastBlock, testStr)
	err := AddBlock(b)

	if err != nil {
		t.Error("Adding block should have not returned error, but returned", err)
	}
}

func createTestBlock() Block {
	const s = "590c9f8430c7435807df8ba9a476e3f1295d46ef210f6efae2043a4c085a569e"
	decodedTestHash, _ := hex.DecodeString(s)
	return Block{"test", 0, []byte{}, decodedTestHash}
}

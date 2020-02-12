package blockchain

import (
	"reflect"
	"testing"

	"github.com/bamada/blockchain/block"
)

func TestBlockchainStarWithGenesisBlock(t *testing.T) {
	bc := New()
	if bc.Chain[0] != *block.Genesis() {
		t.Error("Invalid chain")
	}
}

func TestBlockAdd(t *testing.T) {
	bc := New()
	data := "foo"
	bc.AddBlock(data)

	if bc.GetLastBlock().Data != data {
		t.Error("Add new block failed")
	}
}

func TestValidateChain(t *testing.T) {
	bc := New()
	data := "bar"
	bc.AddBlock(data)

	if !bc.IsValidChain() {
		t.Error("Invalid blockchain")
	}
}

func TestInvalidateCorruptChain(t *testing.T) {
	bc := New()
	data := "bar"
	bc.AddBlock(data)
	bc.Chain[1].Data = "jane"

	if bc.IsValidChain() {
		t.Error("Corrupted chain")
	}
}

func TestReplaceChainWhenReceivedChainLenghHigherThanCurrentChainLengh(t *testing.T) {
	bc1 := New()
	bc2 := New()
	data := "foo"
	bc1.AddBlock(data)
	bc1.ReplaceChain(bc2)
	if reflect.DeepEqual(bc1.Chain, bc2.Chain) {
		t.Error("Replace with invalid chain")
	}
}

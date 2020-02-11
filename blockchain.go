package blockchain

import (
	"errors"

	"github.com/bamada/blockchain/block"
)

// BlockChain blockchain struct
type BlockChain struct {
	Chain []block.Block
}

// AddBlock append new block to the chain
func (bc *BlockChain) AddBlock(data string) *block.Block {
	lb := bc.Chain[len(bc.Chain)-1]
	nb := block.MineBlock(&lb, data)
	bc.Chain = append(bc.Chain, *nb)

	return nb
}

// IsValidChain validate chain
func (bc *BlockChain) IsValidChain(ch *BlockChain) bool {
	if bc.Chain[0] != *block.Genesis() {
		return false
	}

	for i := 1; i < len(bc.Chain); i++ {
		cblock := bc.Chain[i]
		lblock := bc.Chain[i-1]
		if (cblock.LastHash != lblock.Hash) || (cblock.Hash != block.BHash(&cblock)) {
			return false
		}

	}

	return true
}

// ReplaceChain replace current chain
func (bc *BlockChain) ReplaceChain(nbc *BlockChain) error {
	if len(nbc.Chain) <= len(bc.Chain) {
		return errors.New("received chain lenght is lower than current chain")
	} else if !bc.IsValidChain(nbc) {
		return errors.New("received chain is not invalid")
	}

	bc.Chain = nbc.Chain

	return nil
}

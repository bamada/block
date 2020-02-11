package block

import (
	"strings"
	"testing"
)

func TestBlockDataMatchInput(t *testing.T) {
	gBlock := Genesis()
	data := "fake data"

	nBlock := MineBlock(gBlock, data)

	if nBlock.Data != data {
		t.Error("Block data does not match the expected.")
	}
}

func TestBlockLastHashMatchLastBlockHash(t *testing.T) {
	gBlock := Genesis()
	data := "fake data"

	nBlock := MineBlock(gBlock, data)

	if nBlock.LastHash != gBlock.Hash {
		t.Error("Block lastHast does not match lastBlock hash.")
	}
}

func TestGenerateHashThatMatchBlockDifficulty(t *testing.T) {
	gBlock := Genesis()
	data := "fake data"

	nBlock := MineBlock(gBlock, data)

	if nBlock.Hash[0:nBlock.Difficulty] != strings.Repeat("0", nBlock.Difficulty) {
		t.Error("Generated hash does not match block difficulty.")
	}
}

func TestLowersDifficultyForSlowerMinedBlock(t *testing.T) {
	gBlock := Genesis()
	data := "fake data"

	nBlock := MineBlock(gBlock, data)

	if ADifficulty(nBlock, 3600000000) != nBlock.Difficulty-1 {
		t.Error("Lowers mining difficulty failed.")
	}
}

func TestRaisesDifficultyForQuicklyMinedBlock(t *testing.T) {
	gBlock := Genesis()
	data := "fake data"

	nBlock := MineBlock(gBlock, data)

	if ADifficulty(nBlock, 3600000) != nBlock.Difficulty+1 {
		t.Error("Raises mining difficulty failed.")
	}
}

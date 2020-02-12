package block

import (
	"fmt"
	"strings"
	"time"
)

const DIFFICULTY = 2

const MINE_RATE = 3000

// Block represent a basic blockchain component
type Block struct {
	Timestamp  int64
	LastHash   string
	Hash       string
	Data       string
	Nonce      int
	Difficulty int
}

// New instanciate new  block
func New(ts int64, lHash string, hash string, data string, nonce int, diff int) *Block {
	return &Block{Timestamp: ts, LastHash: lHash, Hash: hash, Data: data, Nonce: nonce, Difficulty: diff}
}

func (b Block) String() string {
	return fmt.Sprintf("Timestamp: %d LastHash: %s Hash: %s Nonce: %d Difficulty: %d Data: %s", b.Timestamp, b.LastHash, b.Hash, b.Nonce, b.Difficulty, b.Data)
}

// Genesis generate Genesis block
// Genesis block represent the first block of the blockchain
func Genesis() *Block {
	return New(1000000000, "first-sha256", "------------", "first-data", 0, DIFFICULTY)
}

// MineBlock generate new block
func MineBlock(b *Block, data string) *Block {
	nonce := 0
	lhash := b.Hash
	diff := b.Difficulty
	var timestamp int64
	var hash string

	for {
		nonce++
		timestamp = time.Now().Unix()
		diff = ADifficulty(b, timestamp)
		hash = CHash(fmt.Sprintf("%d%s%s%d%d", timestamp, lhash, data, nonce, diff))
		if hash[0:diff] == strings.Repeat("0", diff) {
			break
		}
	}

	return New(timestamp, lhash, hash, data, nonce, diff)
}

// BHash compute hash of block b
func BHash(b *Block) string {
	return CHash(fmt.Sprintf("%d%s%s%d%d", b.Timestamp, b.LastHash, b.Data, b.Nonce, b.Difficulty))
}

// ADifficulty automatique adjust block difficulity
func ADifficulty(b *Block, ts int64) int {
	if x := (b.Timestamp + MINE_RATE); x > ts {
		return b.Difficulty + 1
	}

	return b.Difficulty - 1
}

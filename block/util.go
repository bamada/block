package block

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/google/uuid"
)

// CHash compute hash
func CHash(d string) string {
	h := sha256.New()
	h.Write([]byte(d))

	return hex.EncodeToString(h.Sum(nil))
}

// ID generate new uuid
func ID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("UUID generation failed.")
	}
	return u.String()
}

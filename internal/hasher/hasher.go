package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(data []byte) string {
	hashArr := sha256.Sum256(data)
	return hex.EncodeToString(hashArr[:])
}

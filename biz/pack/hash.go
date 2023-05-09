package pack

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(raw string) string {
	sha := sha256.New()
	hashed := hex.EncodeToString(sha.Sum([]byte(raw)))
	return hashed
}

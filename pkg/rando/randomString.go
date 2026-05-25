package rando

import (
	"crypto/rand"
	"encoding/hex"
)

func randomString() string {
	bytes := make([]byte, 2)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

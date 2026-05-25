package rand

import (
	"crypto/rand"
	"encoding/hex"
	"path/filepath"
	"strings"
	"time"
)

func GenerateRandomFileNames(fileName string) string {
	ext := filepath.Ext(fileName)
	bytes := make([]byte, 8)
	_, _ = rand.Read(bytes)
	randomHex := hex.EncodeToString(bytes)

	return time.Now().Format("20060102") + "-" + randomHex + strings.ToLower(ext)
}

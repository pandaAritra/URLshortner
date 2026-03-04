package tools

import (
	"math/rand"
	"strings"
	"time"
)

// code generator
// ---------------
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for range 6 {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}
	return sb.String()
}

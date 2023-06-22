// Issue 89
// Pseudo random number generator should
// be avoided for cryptographic purpose

package main

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"time"
)

func hashVal(val []byte) string {
	salt := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	rand.Read(salt)

	saltedVal := append(val, salt...)
	hashed := sha512.Sum512(saltedVal)
	return hex.EncodeToString(hashed[:])
}

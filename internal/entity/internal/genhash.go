package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashHMACSHA256(password, salt string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(salt))

	// Write Data to it
	h.Write([]byte(password))

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}

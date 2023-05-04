package algorithm

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HS256Compute(message, secretKey string) string {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HS256Verify(message, messageHmac, secretKey string) bool {
	return hmac.Equal([]byte(messageHmac), []byte(HS256Compute(message, secretKey)))
}

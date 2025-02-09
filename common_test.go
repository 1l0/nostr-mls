package mls

import "crypto/rand"

func randomBytes(size int) []byte {
	out := make([]byte, size)
	rand.Read(out)
	return out
}

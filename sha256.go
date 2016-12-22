package crypto

import "golang.org/x/crypto/sha3"

// Shake256 calls sha3.ShakeSum256 with 64 bytes of output.
func Shake256(data []byte) []byte {
	var hash [64]byte
	sha3.ShakeSum256(hash[:], data)
	return hash[:]
}

// Shake256 calls sha3.ShakeSum256 with 64 bytes of output.
func Shake256s(data string) []byte {
	var hash [64]byte
	sha3.ShakeSum256(hash[:], []byte(data))
	return hash[:]
}

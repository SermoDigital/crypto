// Package crypto provides basic cryptographic routines.
package crypto

import "golang.org/x/crypto/scrypt"

// Key calls a standardized scrypt.Key.
// 	N = 1 << 15 (32768)
// 	r = 8
// 	p = 1
// 	keylen = 32
// It panics if scrypt.Key panics, but that will only happen if the constants
// are invalid. (Which won't happen.)
func Key(pass, salt []byte) []byte {
	// For parameters, see https://github.com/Tarsnap/scrypt/issues/19
	b, err := scrypt.Key(pass, salt, 1<<15, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return b
}

// Package rng provides routines to create random integers.
package rng

import (
	"crypto/rand"
	"encoding/binary"
)

// Int returns a uniform random value in [0, 1 << INT_BITS - 1). It will panic
// if it cannot obtain enough random bytes from rand.Read.
func Int() int {
	x := Uint64()
	return int(x << 1 >> 1)
}

// Uint64 returns a uniform random value in [0, 1 << 64 - 1]. It will panic if
// it cannot obtain enough random bytes from rand.Read.
func Uint64() uint64 {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(b[:])
}

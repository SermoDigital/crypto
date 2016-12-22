package rng

import (
	"crypto/rand"
	"encoding/binary"
	"log"
)

// Int returns a uniform random value in [0, 1 << INT_BITS - 1).
func Int() int {
	x := Uint64()
	return int(x << 1 >> 1)
}

// Uint64 returns a uniform random value in [0, 1 << 64 - 1]
func Uint64() uint64 {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		log.Fatalln(err)
	}
	return binary.LittleEndian.Uint64(b[:])
}

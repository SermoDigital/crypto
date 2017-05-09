// Timebox is a thin wrapper around nacl/secretbox for time-based secrets.
package timebox

import (
	"crypto/rand"
	"time"

	"golang.org/x/crypto/nacl/secretbox"
)

// SealWith encrypts data the time-sensitive using nacl/secretbox.
func SealWith(data []byte, expires time.Time, nonce *[24]byte, key *[32]byte) ([]byte, error) {
	tb, err := expires.MarshalBinary()
	if err != nil {
		return nil, err
	}

	data1 := make([]byte, len(data)+len(tb)+1)
	data1[0] = byte(len(tb))
	n := copy(data1[1:], tb)
	copy(data1[n+1:], data)

	out := make([]byte, len(nonce), len(nonce)+len(data)+secretbox.Overhead)
	copy(out, nonce[:])
	return secretbox.Seal(out, data1, nonce, key), nil
}

// Seal is shorthand for calling SealWith with a nonce read via rand.Read.
func Seal(data []byte, expires time.Time, key *[32]byte) ([]byte, error) {
	var nonce [24]byte
	_, err := rand.Read(nonce[:])
	if err != nil {
		return nil, err
	}
	return SealWith(data, expires, &nonce, key)
}

// Open is shorthand for calling OpenAt with time.Now as its first argument.
func Open(data []byte, key *[32]byte) ([]byte, bool) {
	return OpenAt(time.Now(), data, key)
}

// OpenAt attempts to unseal the sealed data, returning false if the data has
// expired.
func OpenAt(when time.Time, data []byte, key *[32]byte) (out []byte, ok bool) {
	var nonce [24]byte
	copy(nonce[:], data)

	data, ok = secretbox.Open(out, data[len(nonce):], &nonce, key)
	if !ok {
		return nil, false
	}

	len := int(data[0]) + 1

	var t time.Time
	err := t.UnmarshalBinary(data[1:len])
	if err != nil || !when.Before(t) {
		return nil, false
	}
	return data[len:], true
}

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

	// The data is encoded as follows.
	//
	// [0, 1):      Length of the following binary encoded time.Time value. (N)
	// [1, N]:      Through the prefixed length (byte 0) is the binary encoded
	// 		        time.Time value.
	// (N, N+24]:   Nonce
	// (N+24, ...]: Sealed data

	outlen := 1 + len(tb) + len(nonce)

	out := make([]byte, outlen, outlen+len(data)+secretbox.Overhead)

	// Encode the length of the marshaled time.Time value in the first byte
	// since time.Time.UnmarshalBinary panics if its argument isn't the correct
	// length.
	out[0] = byte(len(tb))
	copy(out[1:], tb)
	copy(out[1+len(tb):], nonce[:])

	return secretbox.Seal(out, data, nonce, key), nil
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
	end := int(data[0]) + 1

	var t time.Time
	err := t.UnmarshalBinary(data[1:end])
	if err != nil || !when.Before(t) {
		return nil, false
	}
	data = data[end:]

	var nonce [24]byte
	n := copy(nonce[:], data)
	return secretbox.Open(out, data[n:], &nonce, key)
}

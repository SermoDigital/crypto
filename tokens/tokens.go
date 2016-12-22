// Tokens is a group of useful functions for creating cryptographically-secure
// tokens. All functions will panic if they cannot gather the required number
// of bytes via rand.Read.
//
// Sizes of the tokens are from Thomas Ptacek's gist:
// https://gist.github.com/tqbf/be58d2d39690c3b366ad
//
// To use the tokens with the gogo/protobuf library, build this with the build
// tag "-protobuf"
package tokens

import (
	"crypto/rand"
	"encoding/base64"
)

const (
	sessionIDSizeBytes    = 512 / 8
	authSizeBytes         = 512 / 8
	saltSizeBytes         = 512 / 8
	randPasswordSizeBytes = 2048 / 8
	urlTokenSizeBytes     = 512 / 8
)

// (n + 2) / 3 * 4 == EncodedLen with padding

const (
	// AuthTokenLength is the length of an encoded AuthToken.
	AuthTokenLength = (authSizeBytes + 2) / 3 * 4

	// SessionIDLength is the length of an encoded SessionID.
	SessionIDLength = (sessionIDSizeBytes + 2) / 3 * 4

	// SaltLength is the length of encoded Salt.
	SaltLength = (saltSizeBytes + 2) / 3 * 4

	// URLTokenLength is the length of an encoded pasword ResetToken.
	URLTokenLength = (urlTokenSizeBytes + 2) / 3 * 4
)

func mustFill(buf []byte) {
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
}

func encode(dst, src []byte) {
	mustFill(src)
	base64.URLEncoding.Encode(dst, src)
}

// AuthToken is a base-64 encoded array suitable for use as an API key or
// authentication token.
type AuthToken [AuthTokenLength]byte

// NewAuthToken generates a cryptographically-secure AuthToken.
func NewAuthToken() AuthToken {
	var tok AuthToken
	var buf [authSizeBytes]byte
	encode(tok[:], buf[:])
	return tok
}

// Salt is an array used in cryptographic functions.
type Salt [SaltLength]byte

// NewSalt generates a cryptographically-secure Salt.
func NewSalt() Salt {
	var salt Salt
	mustFill(salt[:])
	return salt
}

// RandomPassword generates a cryptographically-secure slice of bytes suitable
// for use as a password. The slice may contain null bytes.
func RandomPassword() []byte {
	var buf [randPasswordSizeBytes]byte
	mustFill(buf[:])
	return buf[:]
}

// NewURLToken generates a cryptographically-secure random slice of base-64
// encoded bytes suitable for use in a URL.
func NewURLToken() []byte {
	var buf [urlTokenSizeBytes]byte
	var enc [URLTokenLength]byte
	encode(enc[:], buf[:])
	return enc[:]
}

// SessionID is a base-64 encoded array used to identify a particular session.
type SessionID [SessionIDLength]byte

// NewSessionID creates a cryptographically-secure SessionID.
func NewSessionID() SessionID {
	var buf [sessionIDSizeBytes]byte
	var sid SessionID
	encode(sid[:], buf[:])
	return sid
}

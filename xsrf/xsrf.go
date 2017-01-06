// Package xsrf provides simple routines for creating and retrieving XSRF
// tokens with the Angular framework.
package xsrf

import (
	"errors"
	"net/http"

	"github.com/sermodigital/crypto"
	"github.com/sermodigital/crypto/tokens"
)

// Token is an XSRF token.
type Token struct{ tokens.AuthToken }

// NewToken generates a new Token.
func NewToken() Token {
	return Token{AuthToken: tokens.NewAuthToken()}
}

// Hash returns the SHAKE-256 hash of the token. Hash should be used when
// storing the XSRF token on the backend and hash comparisons should use a
// constant-time comparison function, like those found in crypto/subtle.
func (t Token) Hash() []byte {
	return crypto.Shake256(t.AuthToken[:])
}

// CookieName is the name of the cookie angular checks to determine whether
// it should send an "X-XSRF-TOKEN" header.
const CookieName = "XSRF-TOKEN"

// Cookie returns an http.Cookie with the Name and Value fields set.
func (t Token) Cookie() *http.Cookie {
	return &http.Cookie{Name: CookieName, Value: string(t.AuthToken[:])}
}

// Retrieve fetches the Token from the http.Request, if it exists.
func Retrieve(r *http.Request) (Token, error) {
	var t Token
	c, err := r.Cookie(CookieName)
	if err != nil {
		return t, err
	}
	if len(c.Value) != len(t.AuthToken) {
		return t, errors.New(CookieName + " cookie's value is an incorrect length")
	}
	copy(t.AuthToken[:], c.Value)
	return t, nil
}

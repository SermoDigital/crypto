// Package xsrf provides simple routines for creating and retrieving XSRF
// tokens with the Angular framework.
package xsrf

import (
	"errors"
	"net/http"

	"github.com/sermodigital/crypto"
	"github.com/sermodigital/crypto/tokens"
)

// CookieName is the name of the cookie angular checks to determine whether
// it should send an "X-XSRF-TOKEN" header.
const CookieName = "XSRF-TOKEN"

// Hash returns the SHAKE-256 hash of the token. Hash should be used when
// storing the XSRF token on the backend and hash comparisons should use a
// constant-time comparison function, like those found in crypto/subtle.
func Hash(t tokens.AuthToken) []byte {
	return crypto.Shake256(t[:])
}

// Cookie returns an http.Cookie with the Name and Value fields set.
func Cookie(t tokens.AuthToken) *http.Cookie {
	return &http.Cookie{Name: CookieName, Value: string(t[:])}
}

// Retrieve fetches the Token from the http.Request, if it exists.
func Retrieve(r *http.Request) (t tokens.AuthToken, err error) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		return t, err
	}
	if len(c.Value) != len(t) {
		return t, errors.New(CookieName + " cookie's value is an incorrect length")
	}
	copy(t[:], c.Value)
	return t, nil
}

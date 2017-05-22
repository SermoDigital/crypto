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
// storing the XSRF token on the backend and comparisons should be done in
// constant-time function, using a function like those found in crypto/subtle.
func Hash(t tokens.AuthToken) []byte {
	return crypto.Shake256(t[:])
}

// Cookie returns an http.Cookie with the Name and Value fields set.
func Cookie(t tokens.AuthToken) *http.Cookie {
	return &http.Cookie{Name: CookieName, Value: string(t[:])}
}

// FromRequest fetches the Token from the X-XSRF-TOKEN header, if it exists.
func FromRequest(r *http.Request) (t tokens.AuthToken, err error) {
	return FromHeader(r.Header)
}

// FromHeader fetches the Token from the X-XSRF-TOKEN header, if it exists.
func FromHeader(h http.Header) (t tokens.AuthToken, err error) {
	val := h.Get("X-" + CookieName)
	if len(val) != len(t) {
		return t, errors.New("X-" + CookieName + " header is an incorrect length")
	}
	copy(t[:], val)
	return t, nil
}

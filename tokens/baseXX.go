package tokens

import "encoding/base64"

// Base64Encode encodes the provided []byte using the padded and URL-safe
// base64 scheme.
func Base64Encode(b []byte) []byte {
	enc := make([]byte, base64.URLEncoding.EncodedLen(len(b)))
	base64.URLEncoding.Encode(enc, b)
	return enc
}

// Base64Decode decodes the provided []byte using the padded and URL-safe
// base64 scheme.
func Base64Decode(b []byte) ([]byte, error) {
	dec := make([]byte, base64.URLEncoding.EncodedLen(len(b)))
	n, err := base64.URLEncoding.Decode(dec, b)
	return dec[:n], err
}

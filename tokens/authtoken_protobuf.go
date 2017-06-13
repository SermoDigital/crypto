package tokens

import (
	"encoding/json"
	"errors"
)

func (a AuthToken) Equal(a2 AuthToken) bool {
	return a == a2
}

func (a AuthToken) Size() int {
	return len(a)
}

func (a AuthToken) Marshal() ([]byte, error) {
	return a[:], nil
}

func (a AuthToken) MarshalTo(data []byte) (int, error) {
	return copy(data, a[:]), nil
}

func (a *AuthToken) Unmarshal(data []byte) error {
	if copy(a[:], data) != len(a) {
		return errors.New("tokens.AuthToken.Unmarshal: not enough data")
	}
	return nil
}

var nilToken AuthToken

func (a AuthToken) MarshalJSON() ([]byte, error) {
	if a.Equal(nilToken) {
		return nil, nil
	}
	var buf [len(a) + 2]byte
	buf[0] = '"'
	copy(buf[1:], a[:])
	buf[len(buf)-1] = '"'
	return buf[:], nil
}

func (a *AuthToken) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("invalid JSON passed tok tokens.AuthToken.UnmarshalJSON")
	}
	return a.Unmarshal(data[1 : len(data)-1])
}

var (
	_ json.Marshaler   = AuthToken{}
	_ json.Unmarshaler = (*AuthToken)(nil)
)

type randy interface {
	Intn(n int) int
}

func NewPopulatedAuthToken(_ randy, easy ...bool) *AuthToken {
	a := NewAuthToken()
	return &a
}

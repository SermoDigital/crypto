// +build protobuf
package tokens

import (
	"encoding/json"

	"github.com/sermodigital/errors"
)

func (s SessionID) Equal(s2 SessionID) bool {
	return s == s2
}

func (s SessionID) Size() int {
	return len(s)
}

func (s SessionID) Marshal() ([]byte, error) {
	return s[:], nil
}

func (s SessionID) MarshalTo(data []byte) (int, error) {
	return copy(data, s[:]), nil
}

func (s *SessionID) Unmarshal(data []byte) error {
	if copy(s[:], data) != len(s) {
		return errors.New("tokens.SessionID.Unmarshal: not enough data")
	}
	return nil
}

func (s *SessionID) unmarshal(data string) error {
	if copy(s[:], data) != len(s) {
		return errors.New("tokens.SessionID.Unmarshal: not enough data")
	}
	return nil
}

var nilID SessionID

func (s SessionID) MarshalJSON() ([]byte, error) {
	if s.Equal(nilID) {
		return nil, nil
	}
	var buf [len(s) + 2]byte
	buf[0] = '"'
	copy(buf[1:], s[:])
	buf[len(buf)-1] = '"'
	return buf[:], nil
}

func (s *SessionID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("invalid JSON passed tok tokens.SessionID.UnmarshalJSON")
	}
	return s.Unmarshal(data[1 : len(data)-1])
}

var (
	_ json.Marshaler   = SessionID{}
	_ json.Unmarshaler = (*SessionID)(nil)
)

func NewPopulatedSessionID(_ randy, easy ...bool) *SessionID {
	s := NewSessionID()
	return &s
}

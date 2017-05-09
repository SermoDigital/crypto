package timebox

import (
	"bytes"
	"testing"
	"time"
)

func TestSealOpen(t *testing.T) {
	key := new([32]byte)
	now := time.Now()

	for i, test := range [...]struct {
		data    []byte
		expires time.Duration
		openat  time.Duration
		ok      bool
	}{
		0: {[]byte("foobarbaz"), +5, 0, true},
		1: {[]byte("helloworld!"), -5, 0, false},
		2: {[]byte(""), +1, 0, true},
	} {
		expires := now.Add(test.expires * time.Minute)
		openat := now.Add(test.openat * time.Minute)

		sealed, err := Seal(test.data, expires, key)
		if err != nil {
			t.Fatalf("#%d: %v", i, err)
		}

		opened, ok := OpenAt(openat, sealed, key)
		if ok != test.ok {
			t.Fatalf("#%d: wanted %t, got %t", i, test.ok, ok)
		}

		if ok && !bytes.Equal(test.data, opened) {
			t.Fatalf(`#%d:
wanted: %v
got:    %v`, i, test.data, opened)
		}
	}
}

package crypto

import (
	"crypto/subtle"
	"sync"
)

// Clearer allows deferred clearing of multiple slices while still allowing
// the slices to be used inside goroutines. Add should be called to add another
// reader to the Clearer. Once the slice is finished being used, Done should be
// called.
type Clearer struct {
	mu      sync.RWMutex
	cleared bool
	slices  [][]byte
}

func (c *Clearer) Add()  { c.mu.RLock() }
func (c *Clearer) Done() { c.mu.RUnlock() }

// NewClearer creates a new Clearer initialized with the provided byte slices.
func NewClearer(buf ...[]byte) *Clearer {
	return &Clearer{slices: buf}
}

// Clear locks the Clearer and, if Clear has not already been executed, clears
// each slice it contains. Clear does not block so it need not be run in its
// own goroutine.
func (c *Clearer) Clear() {
	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		if !c.cleared {
			ClearSlice(c.slices...)
		}
	}()
}

// ClearSlice will zero each provided slice.
func ClearSlice(buf ...[]byte) {
	for _, s := range buf {
		// With Go 1.5+ this will simply be a memclr command instead of looping
		// over every byte.
		for i := range s {
			s[i] = 0
		}
	}
}

// Equal returns true if a == b using constant time comparison.
func Equal(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

// Equal returns true if x == y using constant time comparison.
func EqualString(x, y string) bool {
	// Copied from https://golang.org/src/crypto/subtle/constant_time.go?s=490:531#L2
	if len(x) != len(y) {
		return false
	}

	var v byte
	for i := 0; i < len(x); i++ {
		v |= x[i] ^ y[i]
	}
	return subtle.ConstantTimeByteEq(v, 0) == 1
}

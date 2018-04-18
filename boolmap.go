// Package boolmap creates a map of bools using bytes for efficiency (needs benchmarking for memory)
package boolmap

// Map is the default boolmap
type Map map[uint64]byte

// NewMap returns a new, initialised Map
func NewMap() Map {
	return make(Map)
}

// Get returns a bool, represented by a byte, for the specified position
func (m Map) Get(p uint64) byte {
	if m.GetBool(p) {
		return 1
	}
	return 0
}

// GetBool returns a bool for the specified position
func (m Map) GetBool(p uint64) bool {
	return m[p>>3]&(1<<(p&7)) != 0
}

// Set sets a bool, represented by a byte, at the specified position
func (m Map) Set(p uint64, d byte) {
	m.SetBool(p, d != 0)
}

// SetBool sets a bool at the specified position
func (m Map) SetBool(p uint64, d bool) {
	shift := byte(1 << (p & 7))
	pos := p >> 3
	var (
		c  byte
		ok bool
	)
	if c, ok = m[pos]; !ok && !d {
		return
	}
	if d {
		c |= shift
	} else {
		c &^= shift
	}
	if c == 0 {
		delete(m, pos)
	} else {
		m[pos] = c
	}
}

// Slice is a slice of bytes representing bools
type Slice []byte

// NewSlice returnns a new, initialised Slice
func NewSlice() *Slice {
	return NewSliceSize(1)
}

// NewSliceSize returns a new Slice, intitialised to the size given
func NewSliceSize(size uint) *Slice {
	sliceSize := size >> 3
	if size&7 != 0 {
		sliceSize++
	}
	s := make(Slice, sliceSize)
	return &s
}

// Get returns a byte, representing a bool, at the specified position
func (s *Slice) Get(p uint) byte {
	if s.GetBool(p) {
		return 1
	}
	return 0
}

// GetBool returns a bool for the specified position
func (s *Slice) GetBool(p uint) bool {
	pos := p >> 3
	if pos > uint(len(*s)) {
		return false
	}
	return (*s)[pos]&(1<<(p&7)) != 0
}

// Set sets a bool, given as a byte, at the specified position
func (s *Slice) Set(p uint, d byte) {
	s.SetBool(p, d != 0)
}

// SetBool sets a bool at the specified position
func (s *Slice) SetBool(p uint, d bool) {
	pos := p >> 3
	if pos >= uint(len(*s)) {
		if !d {
			return
		}
		if pos < uint(cap(*s)) {
			*s = (*s)[:cap(*s)]
		} else {
			var newData Slice
			if pos < 512 {
				newData = make(Slice, pos<<1)
			} else {
				newData = make(Slice, pos+(pos>>2))
			}
			copy(newData, *s)
			*s = newData
		}
	}
	shift := byte(1 << (p & 7))
	if d {
		(*s)[pos] |= shift
	} else {
		(*s)[pos] &^= shift
	}
}

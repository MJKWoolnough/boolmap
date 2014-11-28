// Package boolmap creates a map of bools using bytes for efficiency (needs benchmarking for memory)
package boolmap

// Map is the default boolmap
type Map struct {
	data map[uint64]byte
}

// NewMap returns a new, initialised Map
func NewMap() Map {
	return Map{make(map[uint64]byte)}
}

// Get returns a bool for the specified position
func (m *Map) Get(p uint64) bool {
	return m.data[p>>3]&(1<<(p&7)) != 0
}

// Set sets a bool at the specified position
func (m *Map) Set(p uint64, d bool) {
	shift := byte(1 << (p & 7))
	pos := p >> 3
	var (
		c  byte
		ok bool
	)
	if c, ok = m.data[pos]; !ok && !d {
		return
	}
	if d {
		c |= shift
	} else {
		c &^= shift
	}
	if c == 0 {
		delete(m.data, pos)
	} else {
		m.data[pos] = c
	}
}

// Slice is a slice of bytes representing bools
type Slice struct {
	data []byte
}

// NewSlice returnns a new, initialised Slice
func NewSlice() *Slice {
	return &Slice{make([]byte, 1)}
}

func NewSliceSize(size uint) *Slice {
	return &Slice{make([]byte, size)}
}

// Get returns a bool for the specified position
func (s *Slice) Get(p uint) bool {
	pos := p >> 3
	if pos > uint(len(s.data)) {
		return false
	}
	return s.data[pos]&(1<<(p&7)) != 0
}

// Set sets a bool at the specified position
func (s *Slice) Set(p uint, d bool) {
	pos := p >> 3
	if pos >= uint(len(s.data)) {
		if !d {
			return
		}
		if pos < uint(cap(s.data)) {
			s.data = s.data[:cap(s.data)]
		} else {
			var newData []byte
			if pos < 512 {
				newData = make([]byte, pos<<1)
			} else {
				newData = make([]byte, pos+(pos>>2))
			}
			copy(newData, s.data)
			s.data = newData
		}
	}
	shift := byte(1 << (p & 7))
	if d {
		s.data[pos] |= shift
	} else {
		s.data[pos] &^= shift
	}
}

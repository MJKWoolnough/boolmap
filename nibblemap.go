package boolmap

// NibbleMap is a map of Nibbles (4-bits, values 0-15)
type NibbleMap map[uint64]byte

// NewNibbleMap return a new, initialised, NibbleMap
func NewNibbleMap() NibbleMap {
	return make(NibbleMap)
}

// Get returns a crumb from the given position
func (n NibbleMap) Get(p uint64) byte {
	d := n[p>>1]
	if p&1 == 0 {
		return d & 15
	}
	return d >> 4
}

// Set sets the crumb at the given position
func (n NibbleMap) Set(p uint64, d byte) {
	pos := p >> 1
	oldData, ok := n[pos]
	if !ok && d == 0 {
		return
	}
	if p&1 == 0 {
		d = oldData&240 | d&15
	} else {
		d = oldData&15 | d<<4
	}
	if d == 0 {
		delete(n, pos)
	} else {
		n[pos] = d
	}
}

// NibbleSlice is a slice of bytes representing nibbles (4-bits)
type NibbleSlice []byte

// NewNibbleSlice returns a new, initialised, CrumbSlice
func NewNibbleSlice() NibbleSlice {
	return NewNibbleSliceSize(1)
}

// NewNibbleSliceSize returns a new NibbleSlice, initialised to the given size
func NewNibbleSliceSize(size uint) NibbleSlice {
	sliceSize := size >> 1
	if size&1 != 0 {
		sliceSize++
	}
	return make(NibbleSlice, sliceSize)
}

// Get returns a crumb from the given position
func (n *NibbleSlice) Get(p uint) byte {
	pos := p >> 1
	if pos >= uint(len(*n)) {
		return 0
	}
	d := (*n)[pos]
	if p&1 == 0 {
		return d & 15
	}
	return d >> 4
}

// Set sets the crumb at the given position
func (n *NibbleSlice) Set(p uint, d byte) {
	pos := p >> 1
	if pos >= uint(len(*n)) {
		if d == 0 {
			return
		}
		if pos < uint(cap(*n)) {
			*n = (*n)[:cap(*n)]
		} else {
			var newData NibbleSlice
			if pos < 512 {
				newData = make([]byte, pos<<1)
			} else {
				newData = make([]byte, pos+(pos>>2))
			}
			copy(newData, *n)
			*n = newData
		}
	}
	oldData := (*n)[pos]
	if p&1 == 0 {
		d = oldData&240 | d&15
	} else {
		d = oldData&15 | d<<4
	}
	(*n)[pos] = d
}

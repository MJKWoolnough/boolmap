package boolmap

// NibbleMap is a map of Nibbles (4-bits, values 0-15)
type NibbleMap struct {
	data map[uint64]byte
}

// NewNibbleMap return a new, initialised, NibbleMap
func NewNibbleMap() *NibbleMap {
	return &NibbleMap{make(map[uint64]byte)}
}

// Get returns a crumb from the given position
func (n *NibbleMap) Get(p uint64) byte {
	d := n.data[p>>1]
	if p&1 == 0 {
		return d & 15
	}
	return d >> 4
}

// Set sets the crumb at the given position
func (n *NibbleMap) Set(p uint64, d byte) {
	pos := p >> 1
	oldData, ok := n.data[pos]
	if !ok && d == 0 {
		return
	}
	if p&1 == 0 {
		d = oldData&240 | d&15
	} else {
		d = oldData&15 | d<<4
	}
	if d == 0 {
		delete(n.data, pos)
	} else {
		n.data[pos] = d
	}
}

// NibbleSlice is a slice of bytes representing nibbles (4-bits)
type NibbleSlice struct {
	data []byte
}

// NewNibbleSlice returns a new, initialised, CrumbSlice
func NewNibbleSlice() *NibbleSlice {
	return &NibbleSlice{make([]byte, 1)}
}

// NewNibbleSliceSize returns a new Crumbslice, initialised to the given size
func NewNibbleSliceSize(size uint) *NibbleSlice {
	sliceSize := size >> 1
	if size&1 != 0 {
		sliceSize++
	}
	return &NibbleSlice{make([]byte, sliceSize)}
}

// Get returns a crumb from the given position
func (n *NibbleSlice) Get(p uint) byte {
	pos := p >> 1
	if pos >= uint(len(n.data)) {
		return 0
	}
	d := n.data[pos]
	if p&1 == 0 {
		return d & 15
	}
	return d >> 4
}

// Set sets the crumb at the given position
func (n *NibbleSlice) Set(p uint, d byte) {
	pos := p >> 1
	if pos >= uint(len(n.data)) {
		if d == 0 {
			return
		}
		if pos < uint(cap(n.data)) {
			n.data = n.data[:cap(n.data)]
		} else {
			var newData []byte
			if pos < 512 {
				newData = make([]byte, pos<<1)
			} else {
				newData = make([]byte, pos+(pos>>2))
			}
			copy(newData, n.data)
			n.data = newData
		}
	}
	oldData := n.data[pos]
	if p&1 == 0 {
		d = oldData&240 | d&15
	} else {
		d = oldData&15 | d<<4
	}
	n.data[pos] = d
}

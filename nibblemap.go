package boolmap

type NibbleMap struct {
	data map[uint64]byte
}

func NewNibbleMap() *NibbleMap {
	return &NibbleMap{make(map[uint64]byte)}
}

func (n *NibbleMap) Get(p uint64) byte {
	d := n.data[p>>1]
	if p&1 == 0 {
		return d & 15
	}
	return d >> 2 & 15
}

func (n *NibbleMap) Set(p uint64, d byte) {
	d &= 15
	if p&1 == 1 {
		d <<= 2
		d |= 15
	} else {
		d |= 240
	}
	n.data[p>>1] &= d
}

type NibbleSlice struct {
	data []byte
}

func NewNibbleSlice() *NibbleSlice {
	return &NibbleSlice{make([]byte, 1)}
}

func NewNibbleSliceSize(size uint) *NibbleSlice {
	return &NibbleSlice{make([]byte, size)}
}

func (n *NibbleSlice) Get(p uint) byte {
	pos := p >> 1
	if pos > uint(len(n.data)) {
		return 0
	}
	d := n.data[pos]
	if p&1 == 0 {
		return d & 15
	}
	return d >> 2 & 15
}

func (n *NibbleSlice) Set(p uint, d byte) {
	d &= 15
	if p&1 == 1 {
		d <<= 2
		d |= 15
	} else {
		d |= 240
	}
	pos := p >> 1
	if pos >= uint(len(n.data)) {
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
	n.data[pos] &= d
}

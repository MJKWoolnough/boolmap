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
	return d >> 4
}

func (n *NibbleMap) Set(p uint64, d byte) {
	pos := p >> 1
	oldData := n.data[pos]
	if p&1 == 0 {
		d = oldData&240 | d&15
	} else {
		d = oldData&15 | d<<4
	}
	n.data[pos] = d
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
	return d >> 4
}

func (n *NibbleSlice) Set(p uint, d byte) {
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
	oldData := n.data[pos]
	if p&1 == 0 {
		d = oldData&240 | d&15
	} else {
		d = oldData&15 | d<<4
	}
	n.data[pos] = d
}

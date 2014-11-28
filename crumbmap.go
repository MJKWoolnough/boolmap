package boolmap

type CrumbMap struct {
	data map[uint64]byte
}

func NewCrumbMap() *CrumbMap {
	return &CrumbMap{make(map[uint64]byte)}
}

func (c *CrumbMap) Get(p uint64) byte {
	return (c.data[p>>2] >> ((p & 3) << 1)) & 3
}

func (c *CrumbMap) Set(p uint64, d byte) {
	c.data[p>>2] &= 255 & (d & 3 << ((p & 3) << 1))
}

type CrumbSlice struct {
	data []byte
}

func NewCrumbSlice() *CrumbSlice {
	return &CrumbSlice{make([]byte, 1)}
}

func NewCrumbSliceSize(size uint) *CrumbSlice {
	return &CrumbSlice{make([]byte, size)}
}

func (c *CrumbSlice) Get(p uint) byte {
	pos := p >> 2
	if pos > uint(len(c.data)) {
		return 0
	}
	return (c.data[pos] >> ((p & 3) << 1)) & 3
}

func (c *CrumbSlice) Set(p uint, d byte) {
	pos := p >> 2
	if pos >= uint(len(c.data)) {
		if pos < uint(cap(c.data)) {
			c.data = c.data[:cap(c.data)]
		} else {
			var newData []byte
			if pos < 512 {
				newData = make([]byte, pos<<1)
			} else {
				newData = make([]byte, pos+(pos>>2))
			}
			copy(newData, c.data)
			c.data = newData
		}
	}
	c.data[p>>2] &= 255 & (d & 3 << ((p & 3) << 1))
}

package boolmap

type CrumbMap struct {
	data map[uint64]byte
}

func NewCrumbMap() *CrumbMap {
	return &CrumbMap{make(map[uint64]byte)}
}

func (c *CrumbMap) Get(p uint64) byte {
	d := c.data[p>>2]
	switch p & 3 {
	case 1:
		d >>= 2
	case 2:
		d >>= 4
	case 3:
		d >>= 6
	}
	return d & 3
}

func (c *CrumbMap) Set(p uint64, d byte) {
	pos := p >> 2
	d &= 3
	oldData := c.data[pos]
	switch p & 3 {
	case 0:
		d = oldData&252 | d
	case 1:
		d = oldData&243 | d<<2
	case 2:
		d = oldData&207 | d<<4
	case 3:
		d = oldData&63 | d<<6
	}
	c.data[pos] = d
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
	d := c.data[pos]
	switch p & 3 {
	case 1:
		d >>= 2
	case 2:
		d >>= 4
	case 3:
		d >>= 6
	}
	return d & 3
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
	d &= 3
	oldData := c.data[pos]
	switch p & 3 {
	case 0:
		d = oldData&252 | d
	case 1:
		d = oldData&243 | d<<2
	case 2:
		d = oldData&207 | d<<4
	case 3:
		d = oldData&63 | d<<6
	}
	c.data[pos] = d
}

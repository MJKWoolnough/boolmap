package boolmap

// CrumbMap is a map of Crumbs (2-bits, values 0, 1, 2, 3)
type CrumbMap struct {
	data map[uint64]byte
}

// NewCrumbMap returns a new, initialised, CrumbMap
func NewCrumbMap() *CrumbMap {
	return &CrumbMap{make(map[uint64]byte)}
}

// Get returns a crumb from the given position
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

// Set sets the crumb at the given position
func (c *CrumbMap) Set(p uint64, d byte) {
	pos := p >> 2
	d &= 3
	oldData, ok := c.data[pos]
	if !ok && d == 0 {
		return
	}
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
	if d == 0 {
		delete(c.data, pos)
	} else {
		c.data[pos] = d
	}
}

// CrumbSlice is a slice of bytes, representing crumbs (2-bits)
type CrumbSlice struct {
	data []byte
}

// NewCrumbSlice returns a new, initialised, CrumbSlice
func NewCrumbSlice() *CrumbSlice {
	return &CrumbSlice{make([]byte, 1)}
}

// NewCrumbSliceSize returns a new Crumbslice, initialised to the given size
func NewCrumbSliceSize(size uint) *CrumbSlice {
	return &CrumbSlice{make([]byte, size)}
}

// Get returns a crumb from the given position
func (c *CrumbSlice) Get(p uint) byte {
	pos := p >> 2
	if pos >= uint(len(c.data)) {
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

// Set sets the crumb at the given position
func (c *CrumbSlice) Set(p uint, d byte) {
	pos := p >> 2
	if pos >= uint(len(c.data)) {
		if d == 0 {
			return
		}
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

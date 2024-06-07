package boolmap

// CrumbMap is a map of Crumbs (2-bits, values 0, 1, 2, 3).
type CrumbMap map[uint64]byte

// NewCrumbMap returns a new, initialised, CrumbMap.
func NewCrumbMap() CrumbMap {
	return make(CrumbMap)
}

// Get returns a crumb from the given position.
func (c CrumbMap) Get(p uint64) byte {
	return (c[p>>2] >> ((p & 3) << 1)) & 3
}

// Set sets the crumb at the given position.
func (c CrumbMap) Set(p uint64, d byte) {
	pos := p >> 2
	d &= 3

	oldData, ok := c[pos]
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
		delete(c, pos)
	} else {
		c[pos] = d
	}
}

// CrumbSlice is a slice of bytes, representing crumbs (2-bits).
type CrumbSlice []byte

// NewCrumbSlice returns a new, initialised, CrumbSlice.
func NewCrumbSlice() *CrumbSlice {
	return NewCrumbSliceSize(1)
}

// NewCrumbSliceSize returns a new Crumbslice, initialised to the given size.
func NewCrumbSliceSize(size uint) *CrumbSlice {
	sliceSize := size >> 2

	if size&3 != 0 {
		sliceSize++
	}

	c := make(CrumbSlice, sliceSize)

	return &c
}

// Get returns a crumb from the given position.
func (c CrumbSlice) Get(p uint) byte {
	pos := p >> 2

	if pos >= uint(len(c)) {
		return 0
	}

	d := c[pos]

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

// Set sets the crumb at the given position.
func (c *CrumbSlice) Set(p uint, d byte) {
	pos := p >> 2

	if pos >= uint(len(*c)) {
		if d == 0 {
			return
		}

		if pos < uint(cap(*c)) {
			(*c) = (*c)[:cap(*c)]
		} else {
			var newData CrumbSlice

			if pos < 512 {
				newData = make(CrumbSlice, pos<<1)
			} else {
				newData = make(CrumbSlice, pos+(pos>>2))
			}

			copy(newData, *c)
			*c = newData
		}
	}

	d &= 3
	oldData := (*c)[pos]

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

	(*c)[pos] = d
}

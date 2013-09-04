// boolmap creates a map of bools using bytes for efficiency (needs benchmarking for memory)
package boolmap

type Map struct {
	data map[uint]byte
}

func NewMap() Map {
	return Map{make(map[uint]byte)}
}

func (m Map) Get(p uint) bool {
	return m.data[p>>3]&(1<<(p&7)) > 0
}

func (m Map) Set(p uint, d bool) {
	if d {
		m.data[p>>3] |= 1 << (p & 7)
	} else {
		m.data[p>>3] &= 0xFF ^ (1 << (p & 7))
	}
	if m.data[p>>3] == 0 {
		delete(m.data, p>>3)
	}
}

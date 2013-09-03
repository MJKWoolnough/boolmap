// boolmap creates a map of bools using uint64s for efficiency (needs benchmarking for memory)
package boolmap

type Map struct {
	data map[uint64]uint64
}

func NewMap() Map {
	return Map{make(map[uint64]uint64)}
}

func (m Map) Get(p uint64) bool {
	return m.data[p>>6]&(1<<(p&63)) > 0
}

func (m Map) Set(p uint64, d bool) {
	if d {
		m.data[p>>6] |= 1 << (p & 63)
	} else {
		m.data[p>>6] &= 0xFFFFFFFFFFFFFFFF ^ (1 << (p & 63))
	}
	if m.data[p>>6] == 0 {
		delete(m.data, p>>6)
	}
}

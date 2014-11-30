package boolmap

import "testing"

type testMap interface {
	Get(p uint64) byte
	Set(p uint64, d byte)
}

type testSlice interface {
	Get(p uint) byte
	Set(p uint, d byte)
}

var (
	_ testMap   = &Map{}
	_ testSlice = &Slice{}
)

func TestBoolMap(t *testing.T) {
	m := NewMap()
	tests := []struct {
		position    uint64
		value       bool
		mapPosition uint64
		mapValue    byte
	}{
		{0, true, 0, 1},
		{1, true, 0, 3},
		{2, true, 0, 7},
		{7, true, 0, 135},
		{8, true, 1, 1},
		{8, false, 1, 0},
		{1, false, 0, 133},
	}
	for n, test := range tests {
		m.SetBool(test.position, test.value)
		if m.data[test.mapPosition] != test.mapValue {
			t.Errorf("test %d: expecting value %d, got %d", n+1, test.mapValue, m.data[test.mapPosition])
		} else if test.mapValue == 0 {
			if _, ok := m.data[test.mapPosition]; ok {
				t.Errorf("test %d: map entry should have been removed on a zero value", n+1)
			}
		}
	}
}

func TestBoolSlice(t *testing.T) {
	s := NewSlice()
	tests := []struct {
		position      uint
		value         bool
		slicePosition uint
		mapValue      byte
	}{
		{0, true, 0, 1},
		{1, true, 0, 3},
		{2, true, 0, 7},
		{7, true, 0, 135},
		{8, true, 1, 1},
		{8, false, 1, 0},
		{1, false, 0, 133},
	}
	for n, test := range tests {
		s.SetBool(test.position, test.value)
		if s.data[test.slicePosition] != test.mapValue {
			t.Errorf("test %d: expecting value %d, got %d", n+1, test.mapValue, s.data[test.slicePosition])
		}
	}
}

func BenchmarkBoolMap(b *testing.B) {
	m := NewMap()
	for n := 0; n < b.N; n++ {
		for i := uint64(0); i < 100; i++ {
			m.SetBool(i, true)
		}
	}
}

func BenchmarkNonBoolMap(b *testing.B) {
	m := make(map[uint64]bool)
	for n := 0; n < b.N; n++ {
		for i := uint64(0); i < 100; i++ {
			m[i] = true
		}
	}
}

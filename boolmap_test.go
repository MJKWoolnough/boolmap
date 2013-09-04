package boolmap

import "testing"

func TestBoolMap(t *testing.T) {
	m := NewMap()
	tests := []struct {
		position    uint
		value       bool
		mapPosition uint
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
		m.Set(test.position, test.value)
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
		s.Set(test.position, test.value)
		if s.data[test.slicePosition] != test.mapValue {
			t.Errorf("test %d: expecting value %d, got %d", n+1, test.mapValue, s.data[test.slicePosition])
		}
	}
}

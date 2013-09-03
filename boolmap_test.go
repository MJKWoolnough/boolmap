package boolmap

import "testing"

func TestBoolMap(t *testing.T) {
	m := NewMap()
	tests := []struct {
		position    uint64
		value       bool
		mapPosition uint64
		mapValue    uint64
	}{
		{0, true, 0, 1},
		{1, true, 0, 3},
		{2, true, 0, 7},
		{63, true, 0, 9223372036854775815},
		{64, true, 1, 1},
		{64, false, 1, 0},
		{1, false, 0, 9223372036854775813},
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

package boolmap

import "testing"

func TestNibbleMap(t *testing.T) {
	tests := []struct {
		val    byte
		pos    uint64
		repeat bool
	}{
		{1, 0, false},
		{0, 0, false},
		{15, 0, false},
		{15, 1, true},
		{6, 0, true},
		{12, 3, true},
		{5, 4, true},
		{14, 33, true},
	}

	m := NewNibbleMap()

	for n, test := range tests {
		m.Set(test.pos, test.val)
		if g := m.Get(test.pos); g != test.val {
			t.Errorf("test %d-1: expecting %d, got %d", n+1, test.val, g)
		}
	}
	for n, test := range tests {
		if !test.repeat {
			continue
		}
		if g := m.Get(test.pos); g != test.val {
			t.Errorf("test %d-2: expecting %d, got %d", n+1, test.val, g)
		}
	}
}

func TestNibbleSlice(t *testing.T) {
	tests := []struct {
		val    byte
		pos    uint
		repeat bool
	}{
		{1, 0, false},
		{0, 0, false},
		{15, 0, false},
		{15, 1, true},
		{6, 0, true},
		{12, 3, true},
		{5, 4, true},
		{14, 33, true},
	}

	s := NewNibbleSlice()

	for n, test := range tests {
		s.Set(test.pos, test.val)
		if g := s.Get(test.pos); g != test.val {
			t.Errorf("test %d-1: expecting %d, got %d", n+1, test.val, g)
		}
	}
	for n, test := range tests {
		if !test.repeat {
			continue
		}
		if g := s.Get(test.pos); g != test.val {
			t.Errorf("test %d-2: expecting %d, got %d", n+1, test.val, g)
		}
	}
}

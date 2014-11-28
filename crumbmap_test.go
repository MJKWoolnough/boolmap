package boolmap

import "testing"

func TestCrumbMap(t *testing.T) {
	tests := []struct {
		val    byte
		pos    uint64
		repeat bool
	}{
		{3, 0, false},
		{2, 0, false},
		{1, 0, false},
		{0, 0, false},
		{3, 0, true},
		{2, 1, true},
		{1, 2, true},
		{0, 3, true},
		{0, 4, true},
		{1, 5, true},
		{2, 6, true},
		{3, 7, true},
	}

	m := NewCrumbMap()

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

func TestCrumbSlice(t *testing.T) {
	tests := []struct {
		val    byte
		pos    uint
		repeat bool
	}{
		{3, 0, false},
		{2, 0, false},
		{1, 0, false},
		{0, 0, false},
		{3, 0, true},
		{2, 1, true},
		{1, 2, true},
		{0, 3, true},
		{0, 4, true},
		{1, 5, true},
		{2, 6, true},
		{3, 7, true},
	}

	s := NewCrumbSlice()

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

func BenchmarkCrumbMap(b *testing.B) {
	m := NewCrumbMap()
	for n := 0; n < b.N; n++ {
		for i := uint64(0); i < 100; i++ {
			m.Set(i, byte(i&3))
		}
	}
}

func BenchmarkNonCrumbMap(b *testing.B) {
	m := make(map[uint64]byte)
	for n := 0; n < b.N; n++ {
		for i := uint64(0); i < 100; i++ {
			m[i] = byte(i & 3)
		}
	}
}

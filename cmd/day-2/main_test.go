package main

import "testing"

func TestDuplicateNumber(t *testing.T) {
	tts := []struct {
		name string
		d    int
		dup  bool
	}{
		//	{"valid 1", 1, false},
		//	{"valid 123", 123, false},
		{"invalid 446446", 446446, true},
		{"invalid 2121212121", 2121212121, true},
		{"invalid 824824824", 824824824, true},
		{"valid 825824824", 825824824, false},
		{"valid 565655", 565655, false},
		{"invalid 565656", 565656, true},
		{"valid 3131313186", 3131313186, false},
	}
	for _, tt := range tts {
		if tt.dup != isRepeatedMultipleTimes(tt.d) {
			t.Errorf("%s wanted:%v, got:%v", tt.name, tt.dup, !tt.dup)
		}
	}
}

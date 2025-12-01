package main

import "testing"

func TestRotate(t *testing.T) {
	tts := []struct {
		name  string
		curr  int64
		dir   direction
		steps int64
		res   int64
	}{
		{"left move 1 from 0, should be 99", 0, left, 1, 99},
		{"right move from 99, should be 0", 99, right, 1, 0},
		{"left move 5 from 5, should be 0", 5, left, 5, 0},
		{"right move 5 from 95, should be 0", 95, right, 5, 0},
		{"right move 110 from 95, should be 5", 95, right, 110, 5},
		{"right move 123 from 95, should be 5", 95, right, 123, 18},
		{"left move 123 from 95, should be 71", 95, left, 123, 72},
		{"left move 68 from 50, should be 82", 50, left, 68, 82},
		{"right move 60 from 95, should be 55", 95, right, 60, 55},
		{"Left move 936, from 10, should be 74", 10, left, 936, 74},
	}

	for _, tt := range tts {
		r := rotateDial(tt.dir, tt.curr, tt.steps)
		if r != tt.res {
			t.Errorf("%s: got %d expected:%d", tt.name, r, tt.res)
		}
	}
}

func TestRotateCountZeroes(t *testing.T) {
	tts := []struct {
		name   string
		curr   int64
		dir    direction
		steps  int64
		res    int64
		clicks int64
	}{
		{"start at 0, move left 5, res 95, 0 clicks", 0, left, 5, 95, 0},
		{"start at 10, move left 15, res 95, 1 clicks", 10, left, 15, 95, 1},
		{"start at 10, move left 115, res 95, 2 clicks", 10, left, 115, 95, 2},
		{"start at 10, move left 215, res 95, 3 clicks", 10, left, 215, 95, 3},
		{"start at 10, move right 90, res 0, 1 clicks", 10, right, 90, 0, 0},
		{"start at 10, move right 190, res 0, 2 clicks", 10, right, 190, 0, 1},
		{"start at 10, move right 200, res 10, 2 clicks", 10, right, 200, 10, 2},
	}

	for _, tt := range tts {
		r, c := rotateDialCountZeroes(tt.dir, tt.curr, tt.steps)
		if r != tt.res {
			t.Errorf("%s: got %d expected:%d", tt.name, r, tt.res)
		}
		if c != tt.clicks {
			t.Errorf("%s: got %d expected:%d", tt.name, c, tt.clicks)

		}
	}
}

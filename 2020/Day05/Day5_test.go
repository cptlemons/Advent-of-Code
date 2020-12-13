package main

import "testing"

func TestParsePass(t *testing.T) {
	inps := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	wants := []int64{357, 567, 119, 820}
	for i, inp := range inps {
		if got, want := parsePass(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	inp := loadInput("input.txt")
	if got, want := part1(inp), int64(838); got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}

}

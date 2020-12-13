package main

import "testing"

func TestPart1(t *testing.T) {
	inp := loadInput("input_test.txt")
	if got, want := part1(inp), 2; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	inp = loadInput("input.txt")
	if got, want := part1(inp), 603; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	inp := loadInput("input_test.txt")
	if got, want := part2(inp), 1; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	inp = loadInput("input.txt")
	if got, want := part2(inp), 404; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

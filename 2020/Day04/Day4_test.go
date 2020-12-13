package main

import "testing"

func TestPart1(t *testing.T) {
	inp := loadInput("input_test.txt")
	if got, want := part1(inp), 2; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	inpReal := loadInput("input.txt")
	if got, want := part1(inpReal), 182; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	inp := loadInput("input_test2.txt")
	if got, want := part2(inp), 4; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	inpReal := loadInput("input.txt")
	if got, want := part2(inpReal), 109; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

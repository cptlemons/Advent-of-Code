package main

import "testing"

func TestPart1(t *testing.T) {
	inp := loadInput("input_test.txt")
	if got, want := part1(inp, 3, 1), 7; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	inpReal := loadInput("input.txt")
	if got, want := part1(inpReal, 3, 1), 257; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	inp := loadInput("input_test.txt")
	if got, want := part2(inp), 336; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	inpReal := loadInput("input.txt")
	if got, want := part2(inpReal), 1744787392; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

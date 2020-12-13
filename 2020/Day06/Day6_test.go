package main

import "testing"

func TestPart1(t *testing.T) {
	inps := []string{"input_test.txt", "input.txt"}
	wants := []int{11, 6686}
	for i, inp := range inps {
		if got, want := part1(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	inps := []string{"input_test.txt", "input.txt"}
	wants := []int{6, 3476}
	for i, inp := range inps {
		if got, want := part2(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

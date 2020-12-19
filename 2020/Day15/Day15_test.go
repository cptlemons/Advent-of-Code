package main

import "testing"

func TestPart1(t *testing.T) {
	inps := [][]int{{0, 3, 6}, {1, 3, 2}, {2, 1, 3}, {1, 2, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}, {9, 19, 1, 6, 0, 5, 4}}
	wants := []int{436, 1, 10, 27, 78, 438, 1836, 1522}
	for i, inp := range inps {
		if got, want := part1(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	inps := [][]int{{0, 3, 6}, {1, 3, 2}, {2, 1, 3}, {1, 2, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}, {9, 19, 1, 6, 0, 5, 4}}
	wants := []int{175594, 2578, 3544142, 261214, 6895259, 18, 362, 18234}
	for i, inp := range inps {
		if got, want := part2(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

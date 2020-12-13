package main

import "testing"

func TestPart1(t *testing.T) {
	files := []string{"input_test.txt", "input_test2.txt", "input.txt"}
	wants := []int{35, 220, 2414}
	for i, file := range files {
		inp := loadInput(file)
		if got, want := part1(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	files := []string{"input_test.txt", "input_test2.txt", "input.txt"}
	wants := []int{8, 19208, 21156911906816}
	for i, file := range files {
		inp := loadInput(file)
		if got, want := part2(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

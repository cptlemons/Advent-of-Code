package main

import "testing"

func TestPart1(t *testing.T) {
	files := []string{"input_test.txt", "input.txt"}
	wants := []int{4, 316}
	for i, file := range files {
		inp := loadInput(file)
		if got, want := part1(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	files := []string{"input_test.txt", "input_test2.txt", "input.txt"}
	wants := []int{32, 126, 11310}
	for i, file := range files {
		inp := loadInput(file)
		if got, want := part2(inp, "shiny gold"), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

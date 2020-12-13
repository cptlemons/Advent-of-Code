package main

import "testing"

func TestPart1(t *testing.T) {
	files := []string{"input_test.txt", "input.txt"}
	pre := []int{5, 25}
	wants := []int{127, 257342611}
	for i, file := range files {
		inp := loadInput(file)
		if got, want := part1(inp, pre[i]), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	files := []string{"input_test.txt", "input.txt"}
	key := []int{127, 257342611}
	wants := []int{62, 35602097}
	for i, file := range files {
		inp := loadInput(file)
		if got, want := part2(inp, key[i]), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

package main

import "testing"

func TestPart1(t *testing.T) {
	files := []string{"input_test.txt", "input.txt"}
	wants := []int{295, 153}
	for i, file := range files {
		inp1, inp2 := loadInputp1(file)
		if got, want := part1(inp1, inp2), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	files := []string{"input_test.txt", "input.txt"}
	wants := []int{1068781, 471793476184394}
	for i, file := range files {
		inp := loadInputp2(file)
		if got, want := part2(inp), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

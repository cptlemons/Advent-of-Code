package main

import "testing"

func TestPart1(t *testing.T) {
	files := [][]string{{"input_test_rules.txt", "input_test_tickets.txt"}, {"input_rules.txt", "input_tickets.txt"}}
	wants := []int{71, 27898}
	for i, file := range files {
		inp1, inp2 := loadInput(file[0], file[1])
		if got, want := part1(inp1, inp2), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	files := [][]string{{"input_test2_rules.txt", "input_test2_tickets.txt"}, {"input_rules.txt", "input_tickets.txt"}}
	myTicks := [][]int{{11, 12, 13}, {181, 131, 61, 67, 151, 59, 113, 101, 79, 53, 71, 193, 179, 103, 149, 157, 127, 97, 73, 191}}
	wants := []int{-1, 2766491048287}
	for i, file := range files {
		inp1, inp2 := loadInput(file[0], file[1])
		if got, want := part2(inp1, inp2, myTicks[i]), wants[i]; got != want {
			t.Errorf("got = %d, want = %d", got, want)
		}
	}
}

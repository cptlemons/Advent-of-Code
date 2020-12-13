package main

import "testing"

func TestPart1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	if got, want := part1(input), 514579; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	realInput := loadInput()
	if got, want := part1(realInput), 712075; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	if got, want := part2(input), 241861950; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
	realInput := loadInput()
	if got, want := part2(realInput), 145245270; got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

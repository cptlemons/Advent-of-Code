package year2024

import (
	"github.com/mattkoler/Advent-of-Code/utils"
	"strconv"
	"strings"
	"testing"
)

func Day01p1() int {
	lines := utils.InputByLines("day01_input.txt")

	var left, right utils.MinHeap[int]

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		split := strings.Fields(line)

		numl, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		numr, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		left.Push(numl)
		right.Push(numr)
	}

	var distance int
	for left.Len() > 0 {
		distance += utils.Abs(right.Pop() - left.Pop())
	}

	return distance
}

func Day01p2() int {
	lines := utils.InputByLines("day01_input.txt")

	left, right := make(map[int]int), make(map[int]int)

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		split := strings.Fields(line)

		numl, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		numr, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		left[numl]++
		right[numr]++
	}

	var total int
	for k, v := range left {
		if rv, ok := right[k]; ok {
			total += k * v * rv
		}
	}
	return total
}

func TestDay1p1(t *testing.T) {
	want := 2031679
	if got := Day01p1(); got != want {
		t.Errorf("Day01p1() = %v, want %v", got, want)
	}
}

func TestDay1p2(t *testing.T) {
	want := 19678534
	if got := Day01p2(); got != want {
		t.Errorf("Day01p2() = %v, want %v", got, want)
	}
}

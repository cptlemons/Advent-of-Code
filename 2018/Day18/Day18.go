package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}

func getInput() (input []string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s\n", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		input = append(input, scn.Text())
	}

	return input
}

func part1(initial []string) (ans int) {
	for i := 0; i < 10; i++ {
		initial = nextMinute(initial)
	}
	return getP1Ans(initial)
}

func part2(initial []string) (ans int) {
	// too many to brute force, need to save state and check for loops
	prevstates := make(map[string]int)
	var cycles int
	var found bool
	for i := 0; i < 1000000000; i++ {
		initial = nextMinute(initial)
		prevstates, cycles, found = checkState(initial, i, prevstates)
		// fast forward cycles if we have seen a state before and reset the map
		if found {
			cycles = i - cycles
			ff := ((1000000000 - i) / cycles)
			i += ff * cycles
			prevstates = make(map[string]int)
		}
	}

	return getP1Ans(initial)
}

func nextMinute(prev []string) (new []string) {
	for y, line := range prev {
		new = append(new, "")
		for x, char := range line {
			trees, lumber := checkAdj(prev, x, y)
			var newchar string
			switch char {
			case '.':
				if trees >= 3 {
					newchar = "|"
				} else {
					newchar = "."
				}
			case '|':
				if lumber >= 3 {
					newchar = "#"
				} else {
					newchar = "|"
				}
			case '#':
				if trees >= 1 && lumber >= 1 {
					newchar = "#"
				} else {
					newchar = "."
				}
			}
			new[y] += newchar
		}
	}
	return new
}

func checkAdj(layout []string, x, y int) (trees, lumber int) {
	// generate the 9 possible offsets
	for ay := -1; ay <= 1; ay++ {
		for ax := -1; ax <= 1; ax++ {
			// skip the point with no offset
			if ay == 0 && ax == 0 {
				continue
			}
			// check bounds
			cy, cx := ay+y, ax+x
			if cy < 0 || cy > len(layout)-1 || cx < 0 || cx > len(layout[0])-1 {
				continue
			}
			if layout[cy][cx] == '|' {
				trees++
			}
			if layout[cy][cx] == '#' {
				lumber++
			}
		}
	}
	return trees, lumber
}

func getP1Ans(state []string) (ans int) {
	var trees, lumber int
	for _, line := range state {
		for _, char := range line {
			if char == '|' {
				trees++
			}
			if char == '#' {
				lumber++
			}
		}
	}
	return trees * lumber
}

func checkState(state []string, cycle int, prevstates map[string]int) (map[string]int, int, bool) {
	var stateString string
	for _, line := range state {
		for _, char := range line {
			stateString += string(char)
		}
	}
	if cycle, ok := prevstates[stateString]; ok {
		return prevstates, cycle, true
	}

	prevstates[stateString] = cycle
	return prevstates, -1, false
}

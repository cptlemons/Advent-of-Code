package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	config, state := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(config, state))
	fmt.Printf("Part 2 answer: %d\n", part2(config, state))

}

func getInput() (config map[string]string, state string) {
	config = make(map[string]string)
	state = "###..#...####.#..###.....####.######.....##.#####.##.##..###....#....##...##...##.#..###..#.#...#..#"

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	getMappings := regexp.MustCompile(`([.#]+) => ([.#])`)

	for scn.Scan() {
		line := getMappings.FindStringSubmatch(scn.Text())
		config[line[1]] = line[2]
	}
	return config, state
}

func part1(config map[string]string, state string) (ans int) {
	var offset int
	for gen := 0; gen < 20; gen++ {
		// adding to front and back to avoid edge equations
		state = "...." + state + "...."
		offset -= 4

		// seed the beginning nulls for offset calculations
		newstate := ".."
		for i := 2; i < len(state)-2; i++ {
			if pot, ok := config[state[i-2:i+3]]; ok {
				newstate += pot
			} else {
				newstate += "."
			}
		}
		for newstate[0] == '.' {
			offset++
			newstate = newstate[1:]
		}
		state = strings.TrimRight(newstate, ".")
	}

	for i, c := range state {
		if c == '#' {
			ans += i + offset
		}
	}
	return ans

}

func part2(config map[string]string, state string) (ans int) {
	var offset int
	// keep track of previous plant configs and their [generation, offset]
	potMemory := make(map[string][2]int)
	for gen := 0; gen < 50000000000; gen++ {
		// detect if we have seen this state before
		if n, ok := potMemory[state]; ok {
			genElapse := gen - n[0]
			offElapse := offset - n[1]

			// fast forward until a cycle before 50BN, -1 because we want to run the last partial cycle below
			cycles := (50000000000 - gen - 1) / genElapse
			gen += cycles * genElapse
			offset += cycles * offElapse
		}
		potMemory[state] = [2]int{gen, offset}
		// adding to front and back to avoid edge equations
		state = "...." + state + "...."
		offset -= 4

		// seed the beginning nulls for offset calculations
		newstate := ".."
		for i := 2; i < len(state)-2; i++ {
			if pot, ok := config[state[i-2:i+3]]; ok {
				newstate += pot
			} else {
				newstate += "."
			}
		}
		for newstate[0] == '.' {
			offset++
			newstate = newstate[1:]
		}
		state = strings.TrimRight(newstate, ".")

	}

	for i, c := range state {
		if c == '#' {
			ans += i + offset
		}
	}
	return ans

}

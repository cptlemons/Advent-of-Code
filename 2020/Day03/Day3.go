package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

}

func loadInput(file string) (inp []string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		inp = append(inp, line)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Scanning error: %s", err)
		os.Exit(1)
	}
	return inp
}

func part1(biome []string, right int, down int) (trees int) {
	var col int
	for y := down; y < len(biome); y += down {
		line := biome[y]
		col += right
		if col >= len(line) {
			col = col - len(line)
		}
		if line[col] == byte('#') {
			trees++
		}

	}
	return trees
}

func part2(biome []string) (total int) {
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var trees []int
	for _, slope := range slopes {
		trees = append(trees, part1(biome, slope[0], slope[1]))
	}
	total = 1
	for _, i := range trees {
		total *= i
	}
	return total
}

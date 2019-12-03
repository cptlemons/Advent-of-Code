package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}

func getInput() (input []int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		i, err := strconv.Atoi(scn.Text())
		if err != nil {
			fmt.Printf("Bad input detected: %s\n", err)
		}
		input = append(input, i)
	}
	return input
}

func part1(input []int) (ans int) {
	for _, v := range input {
		ans += v
	}
	return ans
}

func part2(input []int) (ans int) {
	freqs := make(map[int]bool)
	var freq int
	freqs[0] = true
	for {
		for _, v := range input {
			freq += v
			if ok := freqs[freq]; ok {
				return freq
			} else {
				freqs[freq] = true
			}
		}
	}
}

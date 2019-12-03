package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}
func getInput() (input string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		input = scn.Text()
	}
	return input
}

func part1(input string) (ans int) {
	for i := 0; i < len(input)-1; i++ {
		c := input[i : i+1]
		switch {
		case c >= "a" && c <= "z":
			if strings.ToUpper(c) == input[i+1:i+2] {
				input = input[:i] + input[i+2:]
				i -= 2
			}
		case c >= "A" && c <= "Z":
			if strings.ToLower(c) == input[i+1:i+2] {
				input = input[:i] + input[i+2:]
				i -= 2
			}
		default:
			fmt.Printf("Unexpected char %s at index %d\n", c, i)
		}
		// check to make sure we don't go too far back
		if i < -1 {
			i = -1
		}
	}

	return len(input)
}

func part2(input string) (ans int) {
	rem := 'a'
	minLen := 100000000
	for rem <= 'z' {
		testLen := part1(strings.ReplaceAll(strings.ReplaceAll(input, string(rem-32), ""), string(rem), ""))
		if testLen < minLen {
			minLen = testLen
		}
		rem++
	}
	return minLen
}

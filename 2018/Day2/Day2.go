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
	fmt.Printf("Part 2 answer: %s\n", part2(input))
}

func getInput() (input []string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		input = append(input, scn.Text())
	}
	return input
}

func part1(input []string) (ans int) {
	var countTwos, countThrees int
	for _, line := range input {
		dict := make(map[rune]int)
		for _, char := range line {
			dict[char]++
		}
		var two, three bool
		for _, v := range dict {
			if two && three {
				break
			}
			if v == 2 && !two {
				countTwos++
				two = true
				continue
			}
			if v == 3 && !three {
				countThrees++
				three = true
			}
		}
	}
	return countTwos * countThrees
}

func part2(input []string) (ans string) {
	for i, line := range input {
		for _, comp := range input[i+1:] {
			errors := 1
			var badidx int
			for i := 0; i < len(comp); i++ {
				if line[i] != comp[i] {
					errors--
					badidx = i
				}
				if errors < 0 {
					break
				}
			}
			if errors == 0 {
				return line[:badidx] + line[badidx+1:]
			}
		}
	}
	return "-1"
}

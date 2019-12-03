package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}

type claim struct {
	id, gapL, gapT, width, height int
}

func getInput() (input []claim) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)
	// #1 @ 151,671: 11x15
	match := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	for scn.Scan() {
		line := match.FindStringSubmatch(scn.Text())
		c := claim{
			id:     atoi(line[1]),
			gapL:   atoi(line[2]),
			gapT:   atoi(line[3]),
			width:  atoi(line[4]),
			height: atoi(line[5]),
		}
		input = append(input, c)
	}
	return input
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

type coords struct{ x, y int }

func part1(input []claim) (ans int) {
	claimLocs := make(map[coords]int)
	for _, inp := range input {
		startx := inp.gapL
		starty := inp.gapT
		for w := 0; w < inp.width; w++ {
			for h := 0; h < inp.height; h++ {
				loc := coords{x: startx + w, y: starty + h}
				claimLocs[loc]++
			}
		}
	}
	for _, v := range claimLocs {
		if v > 1 {
			ans++
		}
	}
	return ans
}

func part2(input []claim) (ans int) {
	claimLocs := make(map[coords]int)
	for _, inp := range input {
		startx := inp.gapL
		starty := inp.gapT
		for w := 0; w < inp.width; w++ {
			for h := 0; h < inp.height; h++ {
				loc := coords{x: startx + w, y: starty + h}
				claimLocs[loc]++
			}
		}
	}
	for _, inp := range input {
		startx := inp.gapL
		starty := inp.gapT
		var overlap bool
		for w := 0; w < inp.width; w++ {
			for h := 0; h < inp.height; h++ {
				loc := coords{x: startx + w, y: starty + h}
				if claimLocs[loc] > 1 {
					overlap = true
				}
			}
		}
		if !overlap {
			return inp.id
		}
	}
	return -1
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1 answer: %d\n", part1(getInput()))
	fmt.Printf("Part 2 answer: %d\n", part2(getInput()))
}

func getInput() (input []int) {
	f, err := os.Open("sampleinput.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		nodes := strings.Split(line, " ")
		for _, n := range nodes {
			i, err := strconv.Atoi(n)
			if err != nil {
				fmt.Printf("Unable to conver %s: %s", n, err)
				continue
			}
			input = append(input, i)
		}
	}
	return input
}

func part1(input []int) (ans int) {
	_, ans = findMetadata(input[2:], input[0], input[1])
	return ans
}

func findMetadata(in []int, child, meta int) (out []int, ans int) {
	if child == 0 {
		for i := 0; i < meta; i++ {
			ans += in[i]
		}
		return in[meta:], ans
	}
	for i := child; i > 0; i-- {
		var ansadd int
		in, ansadd = findMetadata(in[2:], in[0], in[1])
		ans += ansadd
	}
	for i := 0; i < meta; i++ {
		ans += in[i]
	}
	return in[meta:], ans
}

func part2(input []int) (ans int) {
	_, ans = findRootVal(input[2:], input[0], input[1])
	return ans
}

func findRootVal(in []int, child, meta int) (out []int, ans int) {
	if child == 0 {
		for i := 0; i < meta; i++ {
			ans += in[i]
		}
		return in[meta:], ans
	}
	var cvalues []int
	for i := child; i > 0; i-- {
		var ansadd int
		in, ansadd = findMetadata(in[2:], in[0], in[1])
		cvalues = append(cvalues, ansadd)
	}
	for i := 0; i < meta; i++ {
		if in[i] < len(cvalues) {
			ans += cvalues[in[i]]
		}
	}
	return in[meta:], ans
}

// 33717 too high

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	passes := loadInput("input.txt")
	fmt.Println(part1(passes))
	fmt.Println(part2(passes))
}

func loadInput(file string) (passes []string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		passes = append(passes, line)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Scanning error: %s", err)
		os.Exit(1)
	}
	return passes
}

func part1(passes []string) (max int64) {
	for _, pass := range passes {
		val := parsePass(pass)
		if val > max {
			max = val
		}
	}
	return max
}

func parsePass(pass string) (id int64) {
	var rstr, cstr string
	for _, char := range pass {
		switch char {
		case 'F':
			rstr += "0"
		case 'B':
			rstr += "1"
		case 'L':
			cstr += "0"
		case 'R':
			cstr += "1"
		}
	}
	row, err := strconv.ParseInt(rstr, 2, 64)
	if err != nil {
		fmt.Printf("Unable to parse row string: %s", err)
	}
	col, err := strconv.ParseInt(cstr, 2, 64)
	if err != nil {
		fmt.Printf("Unable to parse col string: %s", err)
	}
	return row*8 + col
}

func part2(passes []string) int {
	var ids []int
	for _, pass := range passes {
		id := parsePass(pass)
		ids = append(ids, int(id))
	}
	sort.Ints(ids)
	for i, seat := range ids {
		if ids[i+1]-seat == 2 {
			return seat + 1
		}
	}
	return -1
}

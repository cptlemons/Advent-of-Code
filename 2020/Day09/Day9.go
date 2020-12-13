package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

}

func loadInput(file string) (inp []int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Unable to convert %s to int: %s\n", line, err)
		}
		inp = append(inp, val)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return inp
}

func part1(inp []int, pre int) (bad int) {
	for i := pre + 1; i < len(inp); i++ {
		check := inp[i]
		valid := false
		for j, n1 := range inp[i-pre : i] {
			for _, n2 := range inp[i-pre+j : i] {
				if n1+n2 == check {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if !valid {
			return check
		}
	}
	return -1
}

func part2(inp []int, key int) (sum int) {
	var start, offset int
	for i, n1 := range inp {
		sum := n1
		for j, n2 := range inp[i+1:] {
			sum += n2
			if sum == key {
				start = i
				offset = j + 2
				goto solve
			} else if sum > key {
				break
			}
		}
	}
solve:
	slice := inp[start : start+offset]
	sort.Ints(slice)
	return slice[0] + slice[len(slice)-1]
}

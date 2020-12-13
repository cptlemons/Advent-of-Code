package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := loadInput()
	fmt.Println(part1(inp))
	fmt.Println(part2(inp))
}

func loadInput() (inp []int) {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error converting line to int: %s\nLine: %s", err, line)
			os.Exit(1)
		}
		inp = append(inp, n)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Scanning error: %s", err)
		os.Exit(1)
	}
	return inp
}

func part1(nums []int) int {
	for i, n1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n1+nums[j] == 2020 {
				return n1 * nums[j]
			}
		}
	}
	return -1
}

func part2(nums []int) int {
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 2020 {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			n2 := nums[j]
			if n1+n2 > 2020 {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				n3 := nums[k]
				if n1+n2+n3 == 2020 {
					return n1 * n2 * n3
				}
			}
		}
	}
	return -1
}

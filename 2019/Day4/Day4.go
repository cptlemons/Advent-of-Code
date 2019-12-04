package main

import (
	"fmt"
	"strconv"
)

func main() {
	min := 264360
	max := 746325
	fmt.Printf("Part 1 answer: %d\n", part1(min, max))
	fmt.Printf("Part 2 answer: %d\n", part2(min, max))
}

func part1(min, max int) (ans int) {
	for i := min; i <= max; i++ {
		var adj bool
		pw := strconv.Itoa(i)
		valid := true
		for i, c := range pw[1:] {
			prev := rune(pw[i])
			if !adj {
				adj = c == prev
			}
			if prev > c {
				valid = false
				break
			}
		}
		if valid && adj {
			ans++
		}
	}
	return ans
}

func part2(min, max int) (ans int) {
next:
	for i := min; i <= max; i++ {
		pw := strconv.Itoa(i)
		set := make(map[rune]int)
		for i, c := range pw[1:] {
			prev := rune(pw[i])
			set[c]++
			if prev > c {
				continue next
			}
		}
		set[rune(pw[0])]++
		for _, count := range set {
			if count == 2 {
				ans++
				continue next
			}
		}

	}
	return ans
}

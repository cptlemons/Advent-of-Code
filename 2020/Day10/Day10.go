package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {}

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

func part1(inp []int) (ans int) {
	var diff1, diff2, diff3 int
	inp = append(inp, 0)
	sort.Ints(inp)
	inp = append(inp, inp[len(inp)-1]+3)
	for i, n := range inp[:len(inp)-1] {
		n2 := inp[i+1]
		switch n2 - n {
		case 1:
			diff1++
		case 2:
			diff2++
		case 3:
			diff3++
		}
	}
	return diff1 * diff3
}

func part2(inp []int) (ans int) {
	inp = append(inp, 0)
	sort.Ints(inp)
	inp = append(inp, inp[len(inp)-1]+3)
	sets := [][]int{}
	ll := 0
	for i, n := range inp[:len(inp)-1] {
		n2 := inp[i+1]
		if n2-n == 3 {
			sets = append(sets, inp[ll:i+1])
			ll = i + 1
		}
	}
	ans = 1
	for _, set := range sets {
		combos := findCombos(set)
		ans *= combos
	}
	return ans
}

func findCombos(sub []int) (n int) {
	first := sub[0]
	last := sub[len(sub)-1]
	// first and last are fixed in place, get all combos of ints between them
	if len(sub) <= 2 {
		return 1
	}
	subsets := getCombos(sub[1 : len(sub)-1])
	for _, set := range subsets {
		current := first
		valid := true
		for _, num := range set {
			if num-current > 3 {
				valid = false
				break
			}
			current = num
		}
		if last-current > 3 {
			valid = false
		}
		if valid {
			n++
		}
	}
	return n
}

func getCombos(set []int) (sets [][]int) {
	length := int(len(set))
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int
		for object := int(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		sets = append(sets, subset)
	}
	sets = append(sets, []int{})
	return sets
}

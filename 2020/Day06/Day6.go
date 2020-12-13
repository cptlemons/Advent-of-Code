package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

}

func part1(file string) (sum int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	grpCount := make(map[rune]bool)
	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			for range grpCount {
				sum++
			}
			grpCount = make(map[rune]bool)
			continue
		}
		for _, char := range line {
			grpCount[char] = true
		}
	}
	for range grpCount {
		sum++
	}
	return sum
}

func part2(file string) (sum int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	grpCount := make(map[rune]bool)
	newGrp := true
	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			for range grpCount {
				sum++
			}
			grpCount = make(map[rune]bool)
			newGrp = true
			continue
		}
		if newGrp {
			for _, char := range line {
				grpCount[char] = true
			}
			newGrp = false
			continue
		}
		lnGroup := make(map[rune]bool)
		for _, char := range line {
			lnGroup[char] = true
		}
		for k := range grpCount {
			if _, ok := lnGroup[k]; !ok {
				delete(grpCount, k)
			}
		}
	}
	for range grpCount {
		sum++
	}
	return sum
}

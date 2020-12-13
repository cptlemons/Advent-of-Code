package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func loadInput(file string) map[string]map[string]int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	rules := make(map[string]map[string]int)

	for scn.Scan() {
		line := scn.Text()
		split := strings.Split(line, " bags contain ")
		outer := split[0]
		rules[outer] = make(map[string]int)
		if split[1] == "no other bags." {
			rules[outer][""] = 0
			continue
		}
		inners := strings.Split(split[1], ",")
		for _, inner := range inners {
			inner := strings.TrimSpace(inner)
			isplit := strings.Split(inner, " ")
			n, err := strconv.Atoi(isplit[0])
			if err != nil {
				fmt.Printf("Unable to convert %s to a number: %s\n", isplit[0], err)
			}
			bagType := strings.Join(isplit[1:3], " ")
			rules[outer][bagType] = n
		}
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}

	return rules
}

func part1(rules map[string]map[string]int) int {
	outer := make(map[string]bool)
	for k, v := range rules {
		if _, ok := v["shiny gold"]; ok {
			outer[k] = false
		}
	}
	new := true
	for new {
		new = false
		for name, seen := range outer {
			if seen {
				continue
			}
			outer[name] = true
			for k, v := range rules {
				if _, ok := v[name]; ok {
					if _, ok := outer[k]; !ok {
						outer[k] = false
						new = true
					}
				}
			}
		}
	}
	return len(outer)
}

func part2(rules map[string]map[string]int, name string) (total int) {
	inners := rules[name]
	for name, n := range inners {
		if n == 0 {
			return 0
		}
		total += n + n*part2(rules, name)
	}
	return total
}

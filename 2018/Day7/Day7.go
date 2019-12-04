package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %s\n", part1(input))
	input = getInput()
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}

func getInput() map[string][]string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	getSteps := regexp.MustCompile(`Step (.) must be finished before step (.) can begin.`)
	stepUnlocks := make(map[string][]string)
	masterSteps := make(map[string]bool)
	for scn.Scan() {
		steps := getSteps.FindStringSubmatch(scn.Text())
		stepUnlocks[steps[1]] = append(stepUnlocks[steps[1]], steps[2])
		masterSteps[steps[2]] = true
	}

	// add steps that are present but do not unblock anything else
	for k := range masterSteps {
		if _, ok := stepUnlocks[k]; !ok {
			stepUnlocks[k] = nil
		}
	}

	return stepUnlocks
}

func part1(input map[string][]string) (ans string) {
	available := make(map[string]bool)
	findAvailable(input, available, "")
	for len(available) > 0 {
		var keys []string
		for k := range available {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		ans += keys[0]
		findAvailable(input, available, keys[0])
	}
	return ans
}

// given the input map, figure out which steps are currently available removing the optional value r from the requirements
func findAvailable(input map[string][]string, available map[string]bool, r string) {
	delete(input, r)
	delete(available, r)
	if r != "" {
		for k, v := range input {
			for i, sv := range v {
				if r == sv {
					input[k] = append(v[:i], v[i+1:]...)
				}
			}
		}
	}
	stepAvail := make(map[string]bool)
	for k, v := range input {
		// if our key exists, we don't want to override it's value
		if _, ok := stepAvail[k]; !ok {
			stepAvail[k] = true
		}
		// these are always dependent on a prior step so are not initially available
		for _, sv := range v {
			stepAvail[sv] = false
		}
	}
	for k, v := range stepAvail {
		if v {
			available[k] = true
		}
	}
}

func part2(input map[string][]string) (ans int) {
	workers := make(map[string]int)
	available := make(map[string]bool)
	findAvailable(input, available, "")
	for timeNow := 0; timeNow < 10000000; timeNow++ {
		// go through the workers and decrement them, then removing ones at 0 to free them up and remove the items from the map
		for k := range workers {
			workers[k]--
			if workers[k] == 0 {
				findAvailable(input, available, k)
				delete(workers, k)
			}
		}
		// if there are no more steps left, return
		if len(input) <= 0 {
			return timeNow
		}
		// launch available jobs based on the numbers of workers we have available
		for len(available) > 0 {
			if len(workers) >= 5 {
				break
			}
			var keys []string
			for k := range available {
				// skip over available keys that are being worked on
				if _, ok := workers[k]; ok {
					delete(available, k)
					continue
				}
				keys = append(keys, k)

			}
			if len(keys) > 0 {
				sort.Strings(keys)
				workers[keys[0]] = int(60 + rune(keys[0][0]) - 'A' + 1)
				delete(available, keys[0])
			}
		}
	}
	return -1
}

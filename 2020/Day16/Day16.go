package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {}

type rule struct {
	name      string
	nrange    [][]int
	validInts map[int]bool
}

func loadInput(rulesFile, ticksFile string) (rules []*rule, ticks [][]int) {
	f, err := os.Open(rulesFile)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		split1 := strings.Split(line, ": ")
		name := split1[0]
		split2 := strings.Split(split1[1], " or ")
		var nrange [][]int
		for _, pair := range split2 {
			ints := strings.Split(pair, "-")
			var intList []int
			for _, sn := range ints {
				n, err := strconv.Atoi(sn)
				if err != nil {
					fmt.Printf("Unable to convert %s into a int: %s\n", sn, err)
				}
				intList = append(intList, n)
			}
			nrange = append(nrange, intList)
		}
		rules = append(rules, &rule{name: name, nrange: nrange})
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}

	f, err = os.Open(ticksFile)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn = bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		split := strings.Split(line, ",")
		var ints []int
		for _, sn := range split {
			n, err := strconv.Atoi(sn)
			if err != nil {
				fmt.Printf("Unable to convert %s into a int: %s\n", sn, err)
			}
			ints = append(ints, n)
		}
		ticks = append(ticks, ints)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return rules, ticks
}

func part1(rules []*rule, nearTicks [][]int) int {
	validInts := make(map[int]bool)
	for _, rule := range rules {
		for _, nrange := range rule.nrange {
			min, max := nrange[0], nrange[1]
			for i := min; i <= max; i++ {
				validInts[i] = true
			}
		}
	}
	errorRate := 0
	for _, tick := range nearTicks {
		for _, n := range tick {
			if _, ok := validInts[n]; !ok {
				errorRate += n
			}
		}
	}
	return errorRate
}

func part2(rules []*rule, nearTicks [][]int, myTick []int) int {
	validInts := make(map[int]bool)
	for _, rule := range rules {
		rule.validInts = make(map[int]bool)
		for _, nrange := range rule.nrange {
			min, max := nrange[0], nrange[1]
			for i := min; i <= max; i++ {
				validInts[i] = true
				rule.validInts[i] = true
			}
		}
	}

	var validTicks [][]int
	validTicks = append(validTicks, myTick)
	for _, tick := range nearTicks {
		valid := true
		for _, n := range tick {
			if _, ok := validInts[n]; !ok {
				valid = false
				break
			}
		}
		if valid {
			validTicks = append(validTicks, tick)
		}
	}
	// map rule to a list of bools denoting which fields are valid for that rule
	truthMap := make(map[string]map[int]bool)
	for _, rule := range rules {
		truthMap[rule.name] = make(map[int]bool)
		for i := 0; i < len(myTick); i++ {
			valid := true
			for _, tick := range validTicks {
				if _, ok := rule.validInts[tick[i]]; !ok {
					valid = false
					break
				}
			}
			if valid {
				truthMap[rule.name][i] = valid
			}
		}
	}
	if len(truthMap) <= 3 {
		solveTruthTest(truthMap)
		return -1
	}
	return solveTruth(truthMap, myTick)
}

type truthStruct struct {
	name        string
	mappings    map[int]bool
	solvedIndex int
}

func (tm *truthStruct) mappingLength() int {
	return len(tm.mappings)
}

func (tm *truthStruct) deleteMapping(n int) {
	delete(tm.mappings, n)
}

func solveTruthTest(truthMaps map[string]map[int]bool) {
	var truthStructs []*truthStruct
	for name, truthMap := range truthMaps {
		truthStructs = append(truthStructs, &truthStruct{
			name:     name,
			mappings: truthMap,
		})
	}

	for {
		solvedIndex := -1
		for _, ts := range truthStructs {
			if len(ts.mappings) == 1 {
				for k := range ts.mappings {
					solvedIndex = k
					ts.deleteMapping(k)
					ts.solvedIndex = k
				}
			}
		}
		if solvedIndex >= 0 {
			for _, ts := range truthStructs {
				ts.deleteMapping(solvedIndex)
			}
		} else {
			break
		}
	}
	for _, ts := range truthStructs {
		fmt.Println(ts.name, ts.solvedIndex)
	}
}

func solveTruth(truthMaps map[string]map[int]bool, myTick []int) (ans int) {
	var truthStructs []*truthStruct
	for name, truthMap := range truthMaps {
		truthStructs = append(truthStructs, &truthStruct{
			name:     name,
			mappings: truthMap,
		})
	}

	for {
		solvedIndex := -1
		for _, ts := range truthStructs {
			if len(ts.mappings) == 1 {
				for k := range ts.mappings {
					solvedIndex = k
					ts.deleteMapping(k)
					ts.solvedIndex = k
				}
			}
		}
		if solvedIndex >= 0 {
			for _, ts := range truthStructs {
				ts.deleteMapping(solvedIndex)
			}
		} else {
			break
		}
	}
	ans = 1
	for _, ts := range truthStructs {
		if strings.HasPrefix(ts.name, "departure") {
			ans *= myTick[ts.solvedIndex]
		}
	}
	return ans
}

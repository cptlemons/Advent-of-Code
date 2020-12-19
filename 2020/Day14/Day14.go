package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {}

func loadInput(file string) (inp []string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		inp = append(inp, line)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return inp
}

type memMap struct {
	mm      map[int]int
	mask    map[int]int
	floater map[int]bool
}

func part1(input []string) (ans int) {
	memMap := &memMap{
		mm:   make(map[int]int),
		mask: make(map[int]int),
	}
	for _, line := range input {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			memMap.newMask(split[1])
			continue
		}
		loc, err := strconv.Atoi(split[0][4 : len(split[0])-1])
		if err != nil {
			fmt.Printf("Unable to convert %s to int: %s", split[0][4:len(split[0])-1], err)
		}
		val, err := strconv.Atoi(split[1])
		memMap.setValue(loc, val)
	}

	return memMap.sumVals()
}

func (mm *memMap) newMask(mask string) {
	overrideMap := make(map[int]int)
	for i, char := range mask {
		if char == 'X' {
			continue
		}
		if char == '0' {
			overrideMap[35-i] = 0
		}
		if char == '1' {
			overrideMap[35-i] = 1
		}
	}
	mm.mask = overrideMap
}

func (mm *memMap) setValue(loc, val int) {
	valMap := make(map[int]int)
	for i := 35; i >= 0; i-- {
		if val >= rexp(2, i) {
			valMap[i] = 1
			val -= rexp(2, i)
		}
	}
	for k, v := range mm.mask {
		valMap[k] = v
	}
	var sum int
	for k, v := range valMap {
		if v == 1 {
			sum += rexp(2, k)
		}
	}
	mm.mm[loc] = sum
}

func (mm *memMap) sumVals() (sum int) {
	for _, v := range mm.mm {
		sum += v
	}
	return sum
}

func rexp(base, exp int) (n int) {
	if exp == 0 {
		return 1
	}
	return (base * rexp(base, exp-1))
}

func part2(inp []string) (ans int) {
	memMap := &memMap{
		mm:      make(map[int]int),
		mask:    make(map[int]int),
		floater: make(map[int]bool),
	}
	for _, line := range inp {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			memMap.newMaskv2(split[1])
			continue
		}
		loc, err := strconv.Atoi(split[0][4 : len(split[0])-1])
		if err != nil {
			fmt.Printf("Unable to convert %s to int: %s", split[0][4:len(split[0])-1], err)
		}
		val, err := strconv.Atoi(split[1])
		memMap.setValuev2(loc, val)
	}
	return memMap.sumVals()
}

func (mm *memMap) newMaskv2(mask string) {
	overrideMap := make(map[int]int)
	floaterMap := make(map[int]bool)
	for i, char := range mask {
		switch char {
		case 'X':
			floaterMap[35-i] = true
		case '0':
			overrideMap[35-i] = 0
		case '1':
			overrideMap[35-i] = 1
		}
		mm.mask = overrideMap
		mm.floater = floaterMap
	}
}

func (mm *memMap) setValuev2(loc, val int) {
	locMap := make(map[int]int)
	for i := 35; i >= 0; i-- {
		if loc >= rexp(2, i) {
			locMap[i] = 1
			loc -= rexp(2, i)
		}
	}
	for k, v := range mm.mask {
		if v == 1 {
			locMap[k] = v
		}
	}
	var floaters []int
	for k, v := range mm.floater {
		if v {
			floaters = append(floaters, k)
		}
	}
	mm.floaterSetVal(floaters, locMap, val)
}

func (mm *memMap) floaterSetVal(floaters []int, locMap map[int]int, val int) {
	if len(floaters) == 0 {
		var loc int
		for k, v := range locMap {
			if v == 1 {
				loc += rexp(2, k)
			}
		}
		mm.mm[loc] = val
		return
	}
	pop := floaters[0]
	floaters = floaters[1:]
	locMap[pop] = 0
	mm.floaterSetVal(floaters, locMap, val)
	locMap[pop] = 1
	mm.floaterSetVal(floaters, locMap, val)
}

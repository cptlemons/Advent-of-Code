package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {}

func loadInputp1(file string) (earliest int, buses []int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		vals := strings.Split(line, ",")
		if len(vals) == 1 {
			n, err := strconv.Atoi(vals[0])
			if err != nil {
				fmt.Printf("Unable to convert %s to int: %s", vals[0], err)
			}
			earliest = n
			continue
		}
		for i, n := range vals {
			if n == "x" {
				continue
			}
			n, err := strconv.Atoi(vals[i])
			if err != nil {
				fmt.Printf("Unable to convert %s to int: %s", vals[i], err)
			}
			buses = append(buses, n)
		}
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return earliest, buses
}

func part1(earliest int, buses []int) (ans int) {
	var soonest, bb int
	for _, bus := range buses {
		mod, quo := earliest%bus, earliest/bus
		if mod != 0 {
			quo++
		}
		depart := quo * bus
		if soonest == 0 || depart < soonest {
			soonest = depart
			bb = bus
		}
		fmt.Println(soonest, bb, bus)
	}
	return (soonest - earliest) * bb
}

func loadInputp2(file string) (buses []string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	skip := true
	for scn.Scan() {
		line := scn.Text()
		if skip {
			skip = false
			continue
		}
		buses = strings.Split(line, ",")
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return buses
}

func part2(buses []string) (ans int) {
	var busToOffset [][]int
	for i, str := range buses {
		if str == "x" {
			continue
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Unable to convert %s to int: %s\n", str, err)
		}
		busToOffset = append(busToOffset, []int{n, i})
	}
	start, every := 1, 1
	n1 := busToOffset[0][0]
	for i := 1; i < len(busToOffset); i++ {
		n2, off := busToOffset[i][0], busToOffset[i][1]
		start, every = solveTwo(n1, n2, off, start, every)
	}
	return start
}

func solveTwo(n1, n2, noffset, istart, ioffset int) (start, every int) {
	var first int
	i := istart
	for {
		if i%n1 != 0 {
			i += ioffset
			continue
		}
		if (i+noffset)%n2 != 0 {
			i += ioffset
			continue
		}
		if first == 0 {
			first = i
			i += ioffset
			continue
		}
		return first, i - first
	}
}

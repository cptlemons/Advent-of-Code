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

type inst struct {
	inst string
	val  int
}

func loadInput(file string) (inp []inst) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		split := strings.Split(line, " ")
		val, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Unable to convert %s to int: %s\n", split[1], err)
		}
		inp = append(inp, inst{inst: split[0], val: val})
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return inp
}

func part1(insts []inst) (acc int) {
	seen := make(map[int]bool)
	currentLoc := 0
	for {
		if _, ok := seen[currentLoc]; ok {
			break
		}
		seen[currentLoc] = true
		inst := insts[currentLoc]
		switch inst.inst {
		case "nop":
			currentLoc++
		case "acc":
			acc += inst.val
			currentLoc++
		case "jmp":
			currentLoc += inst.val
		default:
			fmt.Printf("Unknown instruction: %s", inst.inst)
			os.Exit(1)
		}
	}
	return acc
}

func part2(insts []inst) (acc int) {
	for i, inst := range insts {
		if inst.inst != "acc" {
			if acc, exit := instSwapCheck(insts, i); exit {
				return acc
			}
		}
	}
	return -1
}

func instSwapCheck(insts []inst, i int) (acc int, exit bool) {
	seen := make(map[int]bool)
	currentLoc := 0
	for currentLoc < len(insts) {
		if _, ok := seen[currentLoc]; ok {
			return -1, false
		}
		seen[currentLoc] = true
		inst := insts[currentLoc]
		if currentLoc == i {
			if inst.inst == "nop" {
				inst.inst = "acc"
			} else {
				inst.inst = "nop"
			}
		}
		switch inst.inst {
		case "nop":
			currentLoc++
		case "acc":
			acc += inst.val
			currentLoc++
		case "jmp":
			currentLoc += inst.val
		default:
			fmt.Printf("Unknown instruction: %s", inst.inst)
			os.Exit(1)
		}
	}
	return acc, true
}

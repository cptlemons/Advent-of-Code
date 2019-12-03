package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 answer:", part1())
	noun, verb := part2()
	fmt.Println("Part 2 answer:", noun, verb)

}

func loadInput() (line string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		line = scn.Text()
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}

func part1() int {
	inp := loadInput()

	split := strings.Split(inp, ",")
	var splt []int
	for _, s := range split {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Bad conversion: %s %s", s, err)
		}
		splt = append(splt, i)
	}
	splt[1] = 12
	splt[2] = 2
	var i int
	for i < len(splt) {
		switch splt[i] {
		case 1:
			splt[splt[i+3]] = splt[splt[i+1]] + splt[splt[i+2]]
			i += 4
		case 2:
			splt[splt[i+3]] = splt[splt[i+1]] * splt[splt[i+2]]
			i += 4
		case 99:
			i = 10000000000000000
		default:
			fmt.Printf("Got %s at %d\n", split[i], i)
		}
	}
	return splt[0]
}

func part2() (noun, verb int) {
	inp := loadInput()

	split := strings.Split(inp, ",")
	var splt []int
	for _, s := range split {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Bad conversion: %s %s", s, err)
		}
		splt = append(splt, i)
	}
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if try2(noun, verb, splt) == 19690720 {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func try2(noun, verb int, splt []int) (ans int) {
	mem := append([]int{}, splt...)
	mem[1] = noun
	mem[2] = verb

	var i int
	for i < len(mem) {
		switch mem[i] {
		case 1:
			mem[mem[i+3]] = mem[mem[i+1]] + mem[mem[i+2]]
			i += 4
		case 2:
			mem[mem[i+3]] = mem[mem[i+1]] * mem[mem[i+2]]
			i += 4
		case 99:
			i = 10000000000000000
		default:
			fmt.Printf("Got %d at %d\n", mem[i], i)
			i = 10000000000000000
			mem[0] = -1
		}
	}
	return mem[0]
}

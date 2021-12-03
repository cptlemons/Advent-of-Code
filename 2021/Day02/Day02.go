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
	inp := loadInput("input.txt")
	fmt.Println(part1(inp))
	fmt.Println(part2(inp))
}

func loadInput(name string) (inp []string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		inp = append(inp, scn.Text())
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Scanning error: %s", err)
		os.Exit(1)
	}
	return inp
}

func part1(cmds []string) (ans int) {
	var h, v int
	for _, cmd := range cmds {
		splt := strings.Split(cmd, " ")
		num, err := strconv.Atoi(splt[1])
		if err != nil {
			log.Fatalf("%s", err)
		}
		switch splt[0] {
		case "forward":
			h += num
		case "up":
			v -= num
		case "down":
			v += num
		}
	}
	return h * v
}

func part2(cmds []string) (ans int) {
	var h, v, aim int
	for _, cmd := range cmds {
		splt := strings.Split(cmd, " ")
		num, err := strconv.Atoi(splt[1])
		if err != nil {
			log.Fatalf("%s", err)
		}
		switch splt[0] {
		case "forward":
			h += num
			v += (aim * num)
		case "up":
			aim -= num
		case "down":
			aim += num
		}
	}
	return h * v
}

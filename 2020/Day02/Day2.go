package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := loadInput("input.txt")
	fmt.Println(part1(inp))
}

type password struct {
	min int
	max int
	req byte
	pwd string
}

func loadInput(file string) (inp []*password) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		splits := strings.Split(line, " ")
		rints := strings.Split(splits[0], "-")
		int1, err := strconv.Atoi(rints[0])
		if err != nil {
			fmt.Printf("Unable to convert first expect int in %s: %s", rints, err)
		}
		int2, err := strconv.Atoi(rints[1])
		if err != nil {
			fmt.Printf("Unable to convert first expect int in %s: %s", rints, err)
		}
		inp = append(inp, &password{
			min: int1,
			max: int2,
			req: splits[1][0],
			pwd: splits[2],
		})
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Scanning error: %s", err)
		os.Exit(1)
	}
	return inp
}

func part1(pwds []*password) (valid int) {
	for _, pwd := range pwds {
		var count int
		for _, char := range pwd.pwd {
			if byte(char) == pwd.req {
				count++
			}
		}
		if count >= pwd.min && count <= pwd.max {
			valid++
		}
	}
	return valid
}

func part2(pwds []*password) (valid int) {
	for _, pwd := range pwds {
		first := pwd.min - 1
		second := pwd.max - 1
		var valid1, valid2 bool
		if first > len(pwd.pwd) {
			continue
		}
		if pwd.pwd[first] == pwd.req {
			valid1 = true
		}
		if second < len(pwd.pwd) {
			if pwd.pwd[second] == pwd.req {
				valid2 = true
			}
		}
		if valid1 != valid2 {
			valid++
		}
	}
	return valid
}

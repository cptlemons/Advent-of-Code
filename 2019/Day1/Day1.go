package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1 answer:", part1())
	fmt.Println("Part 2 answer:", part2())
}

func part1() (fuel int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		t := scn.Text()
		if i, err := strconv.Atoi(t); err == nil {
			fuel += fuelCounter(i)
		}
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return fuel
}

func fuelCounter(mass int) (fuel int) {
	return mass/3 - 2
}

func part2() (total int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		t := scn.Text()
		if i, err := strconv.Atoi(t); err == nil {
			fuel := fuelCounter(i)
			total += fuel
			for true {
				newfuel := fuelCounter(fuel)
				if newfuel < 1 {
					break
				}
				total += newfuel
				fuel = newfuel
			}
		}
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

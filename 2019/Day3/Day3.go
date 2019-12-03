package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type path struct {
	direction string
	distance  int
}

type loc struct {
	x int
	y int
}

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Printf("Part 2 answer: %d\n", part2(input))

}

func getInput() (paths [][]path) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input: %s", err)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		var vals []path
		for _, text := range strings.Split(line, ",") {
			dir := text[0:1]
			dis, err := strconv.Atoi(text[1:])
			if err != nil {
				fmt.Printf("Bad value in %s can't convert to int\n", text)
				continue
			}
			vals = append(vals, path{direction: dir, distance: dis})
		}
		paths = append(paths, vals)
	}
	return paths
}

func part1(input [][]path) (ans int) {
	locations := make(map[loc]bool)
	closest := 100000000000
	for i, line := range input {
		currentloc := loc{x: 0, y: 0}
		for _, vect := range line {
			for vect.distance > 0 {
				switch vect.direction {
				case "U":
					currentloc.y++
				case "D":
					currentloc.y--
				case "R":
					currentloc.x++
				case "L":
					currentloc.x--
				default:
					fmt.Printf("Got a bad vector: %s%d\n", vect.direction, vect.distance)
				}
				if i == 0 {
					locations[currentloc] = true
				} else if ok := locations[currentloc]; ok {
					intersection := abs(currentloc.x) + abs(currentloc.y)
					if intersection < closest {
						closest = intersection
					}
				}
				vect.distance--
			}
		}
	}
	return closest
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func part2(input [][]path) (ans int) {
	locations := make(map[loc]int)
	closest := 100000000000
	for i, line := range input {
		currentloc := loc{x: 0, y: 0}
		var steps int
		for _, vect := range line {
			for vect.distance > 0 {
				switch vect.direction {
				case "U":
					currentloc.y++
				case "D":
					currentloc.y--
				case "R":
					currentloc.x++
				case "L":
					currentloc.x--
				default:
					fmt.Printf("Got a bad vector: %s%d\n", vect.direction, vect.distance)
				}
				steps++
				if i == 0 {
					if _, ok := locations[currentloc]; !ok {
						locations[currentloc] = steps
					}
				} else if steps1, ok := locations[currentloc]; ok {
					distance := steps1 + steps
					if distance < closest {
						closest = distance
					}
				}
				vect.distance--
			}
		}
	}
	return closest
}

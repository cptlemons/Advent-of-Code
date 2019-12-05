package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	input := getInput()
	part1(input)
}

type star struct {
	pos xy
	vel xy
}

type xy struct {
	x int
	y int
}

func getInput() (input []star) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	getStar := regexp.MustCompile(`position=< *(-?\d*), *(-?\d*)> velocity=< *(-?\d*), *(-?\d*)>`)

	for scn.Scan() {
		input = append(input, starToStruct(getStar.FindStringSubmatch(scn.Text())))
	}

	return input
}

func starToStruct(slc []string) (st star) {
	st.pos = xy{x: atoi(slc[1]), y: atoi(slc[2])}
	st.vel = xy{x: atoi(slc[3]), y: atoi(slc[4])}
	return st
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func part1(input []star) {
	smallest := 100000000000000
	var time int
	var bestGrid grid
	for i := 0; i < 100000; i++ {
		testGrid := getGrid(input)
		if testGrid.size < smallest {
			bestGrid = testGrid
			smallest = bestGrid.size
			time = i
		}
		input = timeStep(input)
	}
	printGrid(bestGrid)
	fmt.Println("Time taken:", time)
}

func timeStep(input []star) (output []star) {
	for _, st := range input {
		output = append(output, star{
			pos: xy{
				x: st.pos.x + st.vel.x,
				y: st.pos.y + st.vel.y,
			},
			vel: st.vel})
	}
	return output
}

type grid struct {
	size, lowx, lowy, highx, highy int
	points                         map[int][]int
}

func getGrid(input []star) grid {
	positions := make(map[int][]int)
	lowx, lowy := 1000000, 1000000
	highx, highy := -lowx, -lowy
	for _, st := range input {
		x, y := st.pos.x, st.pos.y
		if x < lowx {
			lowx = x
		} else if x > highx {
			highx = x
		}
		if y < lowy {
			lowy = y
		} else if y > highy {
			highy = y
		}
		positions[st.pos.y] = append(positions[st.pos.y], st.pos.x)
	}
	return grid{
		size:   (highx - lowx) * (highy - lowy),
		lowx:   lowx,
		lowy:   lowy,
		highx:  highx,
		highy:  highy,
		points: positions,
	}
}

func printGrid(input grid) {
	for y := input.lowy; y <= input.highy; y++ {
		var row []string
		points := input.points[y]
		sort.Ints(points)
		for x := input.lowx; x <= input.highx; x++ {
			if len(points) > 0 && x == points[0] {
				row = append(row, "#")
				// some points may overlap
				for points[0] == x && len(points) > 1 {
					points = points[1:]
				}
			} else {
				row = append(row, ".")
			}
		}
		fmt.Println(row)
	}
	fmt.Println()
}

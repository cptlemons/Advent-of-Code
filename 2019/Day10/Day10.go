package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := loadInput()
	ans := part1(input)
	fmt.Printf("Part 1 answer: %d\n", ans)
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}

func loadInput() (input []string) {
	f, err := os.Open("sampleinput.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		input = append(input, scn.Text())

	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

type asteroid struct {
	x, y int
}

func part1(input []string) (ans int) {
	var asteroids []asteroid
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				asteroids = append(asteroids, asteroid{x: x, y: y})
			}
		}
	}
	for i, ast := range asteroids {
		offsets := make(map[[2]int]bool)
		for j, view := range asteroids {
			if j == i {
				continue
			}
			xoff := ast.x - view.x
			yoff := ast.y - view.y
			var xlow, ylow int
			for i := 1; i < 100; i++ {
				if xoff%i == 0 && yoff%i == 0 {
					xlow, ylow = xoff/i, yoff/i
				}
			}
			offsets[[2]int{xlow, ylow}] = true
		}
		if len(offsets) > ans {
			ans = len(offsets)
		}
	}
	return ans
}

func part2(input []string) (ans int) {
	var asteroids []asteroid
	var x, y int
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				asteroids = append(asteroids, asteroid{x: x, y: y})
			}
		}
	}
	var origOffset map[[2]int]bool
	for i, ast := range asteroids {
		offsets := make(map[[2]int]bool)
		for j, view := range asteroids {
			if j == i {
				continue
			}
			xoff := ast.x - view.x
			yoff := ast.y - view.y
			var xlow, ylow int
			for i := 1; i < 100; i++ {
				if xoff%i == 0 && yoff%i == 0 {
					xlow, ylow = xoff/i, yoff/i
				}
			}
			offsets[[2]int{xlow, ylow}] = true
		}
		if len(offsets) > ans {
			ans = len(offsets)
			origOffset = offsets
			x = ast.x
			y = ast.y
		}
	}

	var lastDest [2]int
	var count int
	var swap bool
	vector := 0.0
	for len(origOffset) > 198 {
		if dest, ok := nextHighest(origOffset, vector); !swap && ok {
			fmt.Printf("Destroying asteroid %d at %d, %d with offsets %d, %d\n", count, dest[0]+x, y-dest[1], dest[0], dest[1])
			delete(origOffset, dest)
			vector = float64(dest[0]) / float64(dest[1])
			lastDest = dest
			count++
		} else if !swap {
			fmt.Println("No suitable asteroid found, swapping funcs")
			swap = true
		}
		if swap {
			if dest, ok := nextLowest(origOffset, vector); ok {
				fmt.Printf("Destroying asteroid %d at %d, %d with offsets %d, %d\n", count, dest[0]+x, y-dest[1], dest[0], dest[1])
				delete(origOffset, dest)
				vector = float64(dest[0]) / float64(dest[1])
				lastDest = dest
				count++
			} else {
				fmt.Println("No suitable asteroid found after swapping funcs")
				return -1
			}
		}
	}
	return (lastDest[0]+x)*100 + y - lastDest[1]

	/*
		var asteroids []asteroid
		for y, line := range input {
			for x, char := range line {
				if char == '#' {
					asteroids = append(asteroids, asteroid{x: x, y: y})
				}
			}
		}
		offsets := make(map[[2]int]bool)
		for _, ast := range asteroids {
			xoff := ast.x - x
			yoff := y - ast.y
			offsets[[2]int{xoff, yoff}] = true
		}

		vector := -0.0000000001
		xpos, ypos := true, true
		count := 1
		for i := 0; i < 20; i++ {
			if dest, ok := nextHighest(offsets, vector, xpos, ypos); ok {
				fmt.Printf("Destroying asteroid %d at %d, %d with offsets %d, %d\n", count, dest[0]+x, y-dest[1], dest[0], dest[1])
				delete(offsets, dest)
				vector = float64(dest[0]) / float64(dest[1])
				count++
			} else {
				fmt.Println("No suitable asteroid found")
			}
		}
	*/
}

func nextHighest(offsets map[[2]int]bool, val float64) (next [2]int, found bool) {
	var candidate float64
	for k := range offsets {
		// only looking for negative values of x and positive values of y
		if k[0] > 0 || k[1] < 0 {
			continue
		}
		if float64(k[0])/float64(k[1]) < val {
			if !found {
				candidate = float64(k[0]) / float64(k[1])
				next = k
				found = true
			} else if float64(k[0])/float64(k[1]) > candidate {
				candidate = float64(k[0]) / float64(k[1])
				next = k
			} else if float64(k[0])/float64(k[1]) == candidate {
				if k[0] > next[0] {
					next = k
				} else if k[0] == next[0] && k[1] > next[1] {
					next = k
				}
			}
		}
	}
	return next, found
}

func nextLowest(offsets map[[2]int]bool, val float64) (next [2]int, found bool) {
	var candidate float64
	for k := range offsets {
		// only looking for negative values of x and negative or zero values of y
		if k[0] > 0 || k[1] > 0 {
			continue
		}
		if float64(k[0])/float64(k[1]) > val {
			if !found {
				candidate = float64(k[0]) / float64(k[1])
				next = k
				found = true
			} else if float64(k[0])/float64(k[1]) < candidate {
				candidate = float64(k[0]) / float64(k[1])
				next = k
			} else if float64(k[0])/float64(k[1]) == candidate {
				if k[0] < next[0] {
					next = k
				} else if k[0] == next[0] && k[1] < next[1] {
					next = k
				}
			}
		}
	}
	return next, found
}

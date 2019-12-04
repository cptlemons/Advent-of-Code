package main

import "fmt"

func main() {
	x, y := part1(5034)
	fmt.Printf("Part 1 answer: %d,%d\n", x, y)
	x, y, s := part2(5034)
	fmt.Printf("Part 2 answer: %d,%d,%d\n", x, y, s)
}

func part1(serial int) (x, y int) {
	grid := initGrid(serial)
	return find3x3(grid)
}

func part2(serial int) (x, y, s int) {
	grid := initGrid(serial)
	return findAnysize(grid)
}

// note grid here is y, x
func initGrid(serial int) (grid [][]int) {
	for y := 0; y < 300; y++ {
		grid = append(grid, []int{})
		for x := 0; x < 300; x++ {
			id := (x + 1) + 10
			power := id * (y + 1)
			power += serial
			power *= id
			if power >= 100 {
				power = ((power % 1000) / 100) - 5
			} else {
				power = -5
			}
			grid[y] = append(grid[y], power)
		}
	}
	return grid
}

func find3x3(grid [][]int) (bx, by int) {
	highpower := -10000000
	for y := 0; y < 298; y++ {
		for x := 0; x < 298; x++ {
			power := grid[y][x] + grid[y][x+1] + grid[y][x+2] + grid[y+1][x] + grid[y+1][x+1] + grid[y+1][x+2] + grid[y+2][x] + grid[y+2][x+1] + grid[y+2][x+2]
			if power > highpower {
				highpower = power
				bx = x + 1
				by = y + 1
			}
		}
	}
	return bx, by
}

func findAnysize(grid [][]int) (bx, by, bs int) {
	highpower := -10000000
	size := 1
nextSize:
	for size <= 300 {
	nextY:
		for y := 0; y < 300; y++ {
			if y+size >= 300 {
				size++
				continue nextSize
			}
			for x := 0; x < 300; x++ {
				if x+size >= 300 {
					continue nextY
				}
				var power int
				for s := 0; s < size; s++ {
					slc := grid[y+s][x : x+size]
					for _, i := range slc {
						power += i
					}
				}
				if power > highpower {
					highpower = power
					bs = size
					bx = x + 1
					by = y + 1
				}
			}
		}
		size++
	}
	return bx, by, bs
}

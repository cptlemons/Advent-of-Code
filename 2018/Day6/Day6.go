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
	points, botR := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(points, botR))
	fmt.Printf("Part 2 answer: %d\n", part2(points, botR))
}

type coords struct{ x, y int }

func getInput() (polishedCoords []coords, botR coords) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)
	var roughCoords []coords
	lowestX, lowestY := 1000000, 1000000
	var highestX, highestY int
	for scn.Scan() {
		xy := strings.Split(scn.Text(), ", ")
		xi := atoi(xy[0])
		if xi < lowestX {
			lowestX = xi
		} else if xi > highestX {
			highestX = xi
		}
		yi := atoi(xy[1])
		if yi < lowestY {
			lowestY = yi
		} else if yi > highestY {
			highestY = yi
		}
		roughCoords = append(roughCoords, coords{x: xi, y: yi})
	}

	// reducing the area to the furthest left and up points to help with calculating infinity
	for _, coord := range roughCoords {
		coord.x -= lowestX
		coord.y -= lowestY
		polishedCoords = append(polishedCoords, coord)
	}
	botR.x = highestX
	botR.y = highestY
	return polishedCoords, botR
}

func atoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Bad atoi input: %s", s)
	}
	return i
}

type closest struct {
	point    coords
	distance int
}

func part1(points []coords, botR coords) (ans int) {
	// grid is a map of coords that points to each point and how far away it is from that coord
	grid := make(map[coords]closest)

	// populate the grid with a max distance that is not attainable
	for x := 0; x <= botR.x; x++ {
		for y := 0; y <= botR.y; y++ {
			loc := coords{x: x, y: y}
			grid[loc] = closest{distance: 10000000}
		}
	}

	// for each point we have, calculate its distance to each point on the grid
	for _, p := range points {
		for x := 0; x <= botR.x; x++ {
			for y := 0; y <= botR.y; y++ {
				loc := coords{x: x, y: y}
				distance := absInt(x-p.x) + absInt(y-p.y)
				// if it is shorter, replace the point with this
				if distance < grid[loc].distance {
					grid[loc] = closest{point: p, distance: distance}
					// if distance is the same, keep track of distance but erase the point as it counts for nobody
				} else if distance == grid[loc].distance {
					grid[loc] = closest{distance: distance}
				}
			}
		}
	}
	//figure out which points go into infinity and calculate each areas points
	inifitePoints := make(map[coords]bool)
	pointScore := make(map[coords]int)
	for x := 0; x <= botR.x; x++ {
		for y := 0; y <= botR.y; y++ {
			// if on the edge
			closestPoint := grid[coords{x: x, y: y}].point
			if x == 0 || y == 0 || x == botR.x || y == botR.y {
				inifitePoints[closestPoint] = true
			}
			pointScore[closestPoint]++
		}
	}

	// find the highest point score that is not infinite
	var score int
	for p, v := range pointScore {
		if inifitePoints[p] {
			continue
		}
		if v > score {
			score = v
		}
	}

	return score
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func part2(points []coords, botR coords) (ans int) {
	// grid is a map of coords that points to each point and how far away it is from that coord
	grid := make(map[coords]closest)

	// for each point we have, calculate its distance to each point on the grid
	for _, p := range points {
		for x := 0; x <= botR.x; x++ {
			for y := 0; y <= botR.y; y++ {
				loc := coords{x: x, y: y}
				dist := absInt(x-p.x) + absInt(y-p.y)
				grid[loc] = closest{distance: dist + grid[loc].distance}
			}
		}
	}
	// figure out which points have <10000 distance
	var size int
	for x := 0; x <= botR.x; x++ {
		for y := 0; y <= botR.y; y++ {
			loc := coords{x: x, y: y}
			if grid[loc].distance < 10000 {
				size++
			}
		}
	}
	return size
}

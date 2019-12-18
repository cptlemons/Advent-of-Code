package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	for _, row := range input.grid {
		fmt.Println(row)
	}
	fmt.Println(part1(input))
}

type layout struct {
	grid        []string
	keys, doors map[rune]coord
	start       coord
}

type coord struct {
	x, y int
}

func getInput() (input layout) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	input.keys = make(map[rune]coord)
	input.doors = make(map[rune]coord)

	for scn.Scan() {
		input.grid = append(input.grid, scn.Text())
	}

	for r, row := range input.grid {
		for c, char := range row {
			if char >= 65 && char <= 90 {
				input.doors[char] = coord{c, r}
			} else if char >= 97 && char <= 122 {
				input.keys[char] = coord{c, r}
			} else if char == 64 {
				input.start = coord{c, r}
				input.grid[r] = strings.ReplaceAll(input.grid[r], "@", ".")
			}
		}
	}

	return input
}

type path struct {
	steps int
	loc   coord
	keys  string
}

func copyKeys(old map[byte]bool) (new map[byte]bool) {
	new = make(map[byte]bool)
	for k, v := range old {
		new[k] = v
	}
	return new
}

func copyCoords(old map[coord]bool) (new map[coord]bool) {
	new = make(map[coord]bool)
	for k, v := range old {
		new[k] = v
	}
	return new
}

func checkDoorKey(door byte, keys string) bool {
	for i := range keys {
		if door == keys[i]-32 {
			return true
		}
	}
	return false
}

func checkKey(door byte, keys string) bool {
	for i := range keys {
		if door == keys[i] {
			return true
		}
	}
	return false
}

func part1(l layout) (ans int) {
	starting := path{loc: l.start}
	hits := bfs(l, &starting)
	leastSteps := 10000000
	var mostKeys int
	bestKeySteps := make(map[string]int)

	for i := 0; len(hits) > 0; i++ {
		var newhits []*path
		for _, hit := range hits {
			if hit.steps >= leastSteps {
				continue
			}
			if len(hit.keys) > 0 {
				if steps, ok := bestKeySteps[hit.keys]; ok {
					if steps <= hit.steps {
						continue
					}
				}
				bestKeySteps[hit.keys] = hit.steps
			}
			if len(hit.keys) > mostKeys {
				mostKeys = len(hit.keys)
			}
			if len(hit.keys) >= len(l.keys) && leastSteps > hit.steps {
				leastSteps = hit.steps
			}
			newhits = append(newhits, bfs(l, hit)...)
		}
		hits = newhits
	}
	return leastSteps
}

func bfs(l layout, p *path) (hits []*path) {
	queue := []*path{p}
	visited := make(map[coord]bool)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		for s := -1; s <= 1; s += 2 {
			keys := next.keys
			loc := next.loc
			steps := next.steps
			nx := s + next.loc.x
			if nx >= 0 && nx < len(l.grid[0]) && !visited[coord{x: nx, y: next.loc.y}] {
				xnext := path{keys: keys, loc: loc, steps: steps}
				xnext.steps++
				xnext.loc.x += s
				visited[xnext.loc] = true
				char := l.grid[next.loc.y][nx]
				switch {
				// wall
				case char == '#':
				// open space
				case char == '.':
					queue = append(queue, &xnext)
				// door
				case char >= 65 && char <= 90:
					if checkDoorKey(char, xnext.keys) {
						queue = append(queue, &xnext)
					}
				// key
				case char >= 97 && char <= 122:
					if !checkKey(char, xnext.keys) {
						xnext.keys += string(char)
						hits = append(hits, &xnext)
					}
					queue = append(queue, &xnext)
				default:
					fmt.Println("Unknown: ", char)
				}
			}

			ny := s + next.loc.y
			if ny >= 0 && ny < len(l.grid) && !visited[coord{x: next.loc.x, y: ny}] {
				ynext := path{keys: keys, loc: loc, steps: steps}
				ynext.steps++
				ynext.loc.y += s
				visited[ynext.loc] = true
				char := l.grid[ny][next.loc.x]
				switch {
				// wall
				case char == '#':
				// open space
				case char == '.':
					queue = append(queue, &ynext)
				// door
				case char >= 65 && char <= 90:
					if checkDoorKey(char, ynext.keys) {
						queue = append(queue, &ynext)
					}
				// key
				case char >= 97 && char <= 122:
					if !checkKey(char, ynext.keys) {
						ynext.keys += string(char)
						hits = append(hits, &ynext)
					}
					queue = append(queue, &ynext)
				default:
					fmt.Println("Unknown: ", char)
				}
			}
		}
	}
	return hits
}

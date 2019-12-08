package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	c := getInput()
	winner, ans := part1(c)
	fmt.Printf("Part 1 answer: %s wins with %d points\n", winner, ans)

	for i := 3; i < 201; i++ {
		c = getInput()
		if winner, ans := part2(c, i); winner == "E" && !c.elfDied {
			fmt.Printf("Part 2 answer: E wins with %d points", ans)
			break
		}
	}
}

type cave struct {
	layout  [][]object
	turn    int
	elfDied bool
}

type object struct {
	symbol              string
	hp, atk, x, y, turn int
}

func getInput() *cave {
	c := new(cave)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	var line int
	for scn.Scan() {
		c.layout = append(c.layout, []object{})
		for i, char := range scn.Text() {
			switch char {
			case '#':
				c.layout[line] = append(c.layout[line], object{
					symbol: "#",
				})
			case '.':
				c.layout[line] = append(c.layout[line], object{
					symbol: ".",
				})
			case 'E':
				c.layout[line] = append(c.layout[line], object{
					symbol: "E",
					hp:     200,
					atk:    3,
					x:      i,
					y:      line,
				})
			case 'G':
				c.layout[line] = append(c.layout[line], object{
					symbol: "G",
					hp:     200,
					atk:    3,
					x:      i,
					y:      line,
				})
			default:
				fmt.Printf("Unknown character: %c", char)
			}
		}
		line++
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return c
}

func (c *cave) printCave() {
	for _, row := range c.layout {
		var line string
		var hp string
		for _, char := range row {
			line += char.symbol
			if char.hp != 0 {
				hp += char.symbol + ":" + strconv.Itoa(char.hp) + " "
			}
		}
		fmt.Println(line, hp)
	}
}

func (c *cave) takeTurn() (string, int, bool) {
	for _, row := range c.layout {
		for _, char := range row {
			switch char.symbol {
			case "#":
				continue
			case ".":
				continue
			case "E", "G":
				if winner, ans, done := c.unitTurn(char); done {
					return winner, ans, done
				}
			default:
				fmt.Printf("Unknown object: %#v\n", char)
			}
		}
	}
	c.turn++
	return "nobody", 0, false
}

// ordering is in reading order for tiebreakers
const (
	N int = iota
	W
	E
	S
)

func (c *cave) unitTurn(u object) (winner string, ans int, done bool) {
	// check to see if the unit has already taken a turn
	if u.turn > c.turn {
		return
	}
	if winner, hp, done := c.done(); done {
		return winner, c.turn * hp, true
	}
	u = c.move(u)
	c.attack(u) // TODO (kevlar): Kyle made a mistake here

	return "nobody", -1, false
}

func (c *cave) done() (winner string, hp int, done bool) {
	var unit string
	for _, row := range c.layout {
		for _, char := range row {
			switch char.symbol {
			case "E", "G":
				if unit != "" && char.symbol != unit {
					return "nobody", -1, false
				}
				unit = char.symbol
				hp += char.hp
			}
		}
	}
	return unit, hp, true
}

type step struct {
	path    string
	x, y    int
	visited map[[2]int]bool
}

func (c *cave) move(u object) object {
	// implement bfs to figure out a step
	stepDir := c.findPath(u)
	switch stepDir {
	case "N":
		u.y--
		c.layout[u.y][u.x] = u
		c.layout[u.y+1][u.x] = object{symbol: "."}
	case "W":
		u.x--
		c.layout[u.y][u.x] = u
		c.layout[u.y][u.x+1] = object{symbol: "."}
	case "E":
		u.x++
		c.layout[u.y][u.x] = u
		c.layout[u.y][u.x-1] = object{symbol: "."}
	case "S":
		u.y++
		c.layout[u.y][u.x] = u
		c.layout[u.y-1][u.x] = object{symbol: "."}
	}
	return u
}

func (c *cave) findPath(u object) (stepDir string) {
	var tracking []step
	visited := make(map[[2]int]bool)
	tracking = append(tracking, step{
		path: "",
		x:    u.x,
		y:    u.y,
	})

	visited[[2]int{u.x, u.y}] = true

	for len(tracking) > 0 {
		current := tracking[0]
		tracking = tracking[1:]
		for i := N; i <= S; i++ {
			loc := [2]int{current.x, current.y}
			var dir string
			switch i {
			case N:
				dir = "N"
				loc[1]--
			case S:
				dir = "S"
				loc[1]++
			case W:
				dir = "W"
				loc[0]--
			case E:
				dir = "E"
				loc[0]++
			}
			// skip if we have been here before
			if _, ok := visited[loc]; ok {
				continue
			}

			switch symbol := c.layout[loc[1]][loc[0]].symbol; symbol {
			// found a wall
			case "#":
				continue
			// found a friendly unit
			case "G", "E":
				if u.symbol == symbol {
					continue
				}
				// found an enemy unit
				if len(current.path) > 0 {
					return current.path[0:1]
				}
				return ""
			}

			tracking = append(tracking, step{
				path: current.path + dir,
				x:    loc[0],
				y:    loc[1],
			})

			visited[loc] = true
		}
	}
	return ""
}

func (c *cave) attack(u object) {
	// check adjacent units
	// tiebreaks: lowest hp then reading order

	enemy := object{hp: 10000}
	for i := N; i <= S; i++ {
		loc := [2]int{u.x, u.y}
		switch i {
		case N:
			loc[1]--
		case S:
			loc[1]++
		case W:
			loc[0]--
		case E:
			loc[0]++
		}
		poten := c.layout[loc[1]][loc[0]]
		switch symbol := poten.symbol; symbol {
		case "G", "E":
			if u.symbol == symbol {
				continue
			}
			if poten.hp < enemy.hp {
				enemy = poten
			}
		}
	}
	// if we found an enemy to attack
	if enemy.symbol != "" {
		c.layout[enemy.y][enemy.x].hp -= u.atk
		if c.layout[enemy.y][enemy.x].hp <= 0 {
			if enemy.symbol == "E" {
				c.elfDied = true
			}
			c.layout[enemy.y][enemy.x] = object{symbol: "."}
		}
	}
	c.layout[u.y][u.x].turn++
}

func part1(c *cave) (winner string, ans int) {
	for i := 0; i < 100000; i++ {
		if winner, ans, done := c.takeTurn(); done {
			return winner, ans
		}
	}
	return "nobody", -1
}

func part2(c *cave, atk int) (winner string, ans int) {
	c.newAttack(atk)
	for i := 0; i < 100000; i++ {
		if winner, ans, done := c.takeTurn(); winner == "E" && done {
			return winner, ans
		}
	}
	return "", -1
}

func (c *cave) newAttack(atk int) {
	// todo: modify attack values
	for y, row := range c.layout {
		for x, char := range row {
			switch char.symbol {
			case "E":
				char.atk = atk
				c.layout[y][x] = char
			}
		}
	}
}

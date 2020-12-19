package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {}

type action struct {
	action string
	val    int
}

func loadInput(file string) (actions []action) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		acn := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Unable to convert %s to int: %s", line[1:], err)
			os.Exit(1)
		}
		actions = append(actions, action{action: string(acn), val: val})
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return actions
}

type ship struct {
	curx   int
	cury   int
	facing int
	wayx   int
	wayy   int
}

func part1(actions []action) (dist int) {
	boat := ship{facing: 90} // facing east
	for _, act := range actions {
		switch act.action {
		case "N":
			boat.cury += act.val
		case "S":
			boat.cury -= act.val
		case "E":
			boat.curx += act.val
		case "W":
			boat.curx -= act.val
		case "L":
			boat.facing -= act.val
			for boat.facing < 0 {
				boat.facing += 360
			}
		case "R":
			boat.facing += act.val
			for boat.facing >= 360 {
				boat.facing -= 360
			}
		case "F":
			switch boat.facing {
			case 0:
				boat.cury += act.val
			case 90:
				boat.curx += act.val
			case 180:
				boat.cury -= act.val
			case 270:
				boat.curx -= act.val
			}

		}
	}

	return boat.absCurx() + boat.absCury()
}

func (s ship) absCurx() int {
	if s.curx < 0 {
		return -s.curx
	}
	return s.curx
}

func (s ship) absCury() int {
	if s.cury < 0 {
		return -s.cury
	}
	return s.cury
}

func part2(actions []action) (dist int) {
	boat := ship{wayx: 10, wayy: 1} // facing east
	for _, act := range actions {
		switch act.action {
		case "N":
			boat.wayy += act.val
		case "S":
			boat.wayy -= act.val
		case "E":
			boat.wayx += act.val
		case "W":
			boat.wayx -= act.val
		case "L":
			switch act.val {
			case 90:
				boat.wayx, boat.wayy = -boat.wayy, boat.wayx
			case 180:
				boat.wayx, boat.wayy = -boat.wayx, -boat.wayy
			case 270:
				boat.wayx, boat.wayy = boat.wayy, -boat.wayx
			}
		case "R":
			switch act.val {
			case 90:
				boat.wayx, boat.wayy = boat.wayy, -boat.wayx
			case 180:
				boat.wayx, boat.wayy = -boat.wayx, -boat.wayy
			case 270:
				boat.wayx, boat.wayy = -boat.wayy, boat.wayx
			}
		case "F":
			boat.curx += boat.wayx * act.val
			boat.cury += boat.wayy * act.val
		}
	}

	return boat.absCurx() + boat.absCury()
}

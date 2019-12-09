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
	r := loadInput()
	r.inputs = append(r.inputs, 1)
	r.d5p1exec()
	fmt.Printf("Part 1 answer: %d\n", r.outputs[len(r.outputs)-1])

	r = loadInput()
	r.inputs = append(r.inputs, 5)
	r.d5p2exec()
	fmt.Printf("Part 2 answer: %d\n", r.outputs[len(r.outputs)-1])
}

func loadInput() (r *commands) {
	r = new(commands)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	var line string
	for scn.Scan() {
		line = scn.Text()
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	splt := strings.Split(line, ",")
	for _, cmd := range splt {
		r.cmds = append(r.cmds, atoi(cmd))
	}
	return r
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

type commands struct {
	cmds, inputs, outputs []int
	currentloc            int
}

func (r *commands) d5p1exec() {
	for r.currentloc < len(r.cmds) {
		cmd := r.cmds[r.currentloc]
		opcode := cmd % 100
		modes := []int{(cmd / 100) % 10, (cmd / 1000) % 10, (cmd / 10000) % 10}
		switch opcode {
		case 1: // 1st + 2param put in 3rd param, move ahead 4
			var first, second int
			if modes[0] == 0 {
				first = r.cmds[r.cmds[r.currentloc+1]]
			} else {
				first = r.cmds[r.currentloc+1]
			}
			if modes[1] == 0 {
				second = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				second = r.cmds[r.currentloc+2]
			}
			if modes[2] == 0 {
				r.cmds[r.cmds[r.currentloc+3]] = first + second
			} else {
				r.cmds[r.currentloc+3] = first + second
			}
			r.currentloc += 4
		case 2: // 1st * 2param put in 3rd param, move ahead 4
			var first, second int
			if modes[0] == 0 {
				first = r.cmds[r.cmds[r.currentloc+1]]
			} else {
				first = r.cmds[r.currentloc+1]
			}
			if modes[1] == 0 {
				second = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				second = r.cmds[r.currentloc+2]
			}
			if modes[2] == 0 {
				r.cmds[r.cmds[r.currentloc+3]] = first * second
			} else {
				r.cmds[r.currentloc+3] = first * second
			}
			r.currentloc += 4
		case 3: // input
			if modes[0] == 0 {
				r.cmds[r.cmds[r.currentloc+1]] = r.inputs[0]
			} else {
				r.cmds[r.currentloc+1] = r.inputs[0]
			}
			r.currentloc += 2
			r.inputs = r.inputs[1:]
		case 4: // output
			if modes[0] == 0 {
				r.outputs = append(r.outputs, r.cmds[r.cmds[r.currentloc+1]])
			} else {
				r.outputs = append(r.outputs, r.cmds[r.currentloc+1])
			}
			r.currentloc += 2
		case 99: // halt
			r.currentloc = 1000000
			return
		default:
			fmt.Printf("Unhandled cmd case %d\n", opcode)
			return
		}
	}
}

func (r *commands) d5p2exec() {
	for r.currentloc < len(r.cmds) {
		cmd := r.cmds[r.currentloc]
		opcode := cmd % 100
		modes := []int{(cmd / 100) % 10, (cmd / 1000) % 10, (cmd / 10000) % 10}
		switch opcode {
		case 1: // 1st + 2param put in 3rd param, move ahead 4
			var first, second int
			if modes[0] == 0 {
				first = r.cmds[r.cmds[r.currentloc+1]]
			} else {
				first = r.cmds[r.currentloc+1]
			}
			if modes[1] == 0 {
				second = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				second = r.cmds[r.currentloc+2]
			}
			if modes[2] == 0 {
				r.cmds[r.cmds[r.currentloc+3]] = first + second
			} else {
				r.cmds[r.currentloc+3] = first + second
			}
			r.currentloc += 4
		case 2: // 1st * 2param put in 3rd param, move ahead 4
			var first, second int
			if modes[0] == 0 {
				first = r.cmds[r.cmds[r.currentloc+1]]
			} else {
				first = r.cmds[r.currentloc+1]
			}
			if modes[1] == 0 {
				second = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				second = r.cmds[r.currentloc+2]
			}
			if modes[2] == 0 {
				r.cmds[r.cmds[r.currentloc+3]] = first * second
			} else {
				r.cmds[r.currentloc+3] = first * second
			}
			r.currentloc += 4
		case 3: // input
			if modes[0] == 0 {
				r.cmds[r.cmds[r.currentloc+1]] = r.inputs[0]
			} else {
				r.cmds[r.currentloc+1] = r.inputs[0]
			}
			r.currentloc += 2
			r.inputs = r.inputs[1:]
		case 4: // output
			if modes[0] == 0 {
				r.outputs = append(r.outputs, r.cmds[r.cmds[r.currentloc+1]])
			} else {
				r.outputs = append(r.outputs, r.cmds[r.currentloc+1])
			}
			r.currentloc += 2
		case 5: // jump if first param is non-zero
			var jump bool
			if modes[0] == 0 {
				jump = r.cmds[r.cmds[r.currentloc+1]] != 0
			} else {
				jump = r.cmds[r.currentloc+1] != 0
			}
			if !jump {
				r.currentloc += 3
				continue
			}
			if modes[1] == 0 {
				r.currentloc = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				r.currentloc = r.cmds[r.currentloc+2]
			}

		case 6: // jump if first param is zero
			var jump bool
			if modes[0] == 0 {
				jump = r.cmds[r.cmds[r.currentloc+1]] == 0
			} else {
				jump = r.cmds[r.currentloc+1] == 0
			}
			if !jump {
				r.currentloc += 3
				continue
			}
			if modes[1] == 0 {
				r.currentloc = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				r.currentloc = r.cmds[r.currentloc+2]
			}
		case 7: // store 1 if first < second else 0
			var first, second int
			if modes[0] == 0 {
				first = r.cmds[r.cmds[r.currentloc+1]]
			} else {
				first = r.cmds[r.currentloc+1]
			}
			if modes[1] == 0 {
				second = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				second = r.cmds[r.currentloc+2]
			}
			var val int
			if first < second {
				val = 1
			}
			if modes[2] == 0 {
				r.cmds[r.cmds[r.currentloc+3]] = val
			} else {
				r.cmds[r.currentloc+3] = val
			}
			r.currentloc += 4
		case 8: // store 1 if first == second else zero
			var first, second int
			if modes[0] == 0 {
				first = r.cmds[r.cmds[r.currentloc+1]]
			} else {
				first = r.cmds[r.currentloc+1]
			}
			if modes[1] == 0 {
				second = r.cmds[r.cmds[r.currentloc+2]]
			} else {
				second = r.cmds[r.currentloc+2]
			}
			var val int
			if first == second {
				val = 1
			}
			if modes[2] == 0 {
				r.cmds[r.cmds[r.currentloc+3]] = val
			} else {
				r.cmds[r.currentloc+3] = val
			}
			r.currentloc += 4
		case 99: // halt
			r.currentloc = 1000000
			return
		default:
			fmt.Printf("Unhandled cmd case %d\n", opcode)
			return
		}
	}
}

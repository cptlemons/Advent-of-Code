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
	part1()
}

func part1() {
	var max int
	for one := 0; one < 5; one++ {
		for two := 0; two < 5; two++ {
			if two == one {
				continue
			}
			for three := 0; three < 5; three++ {
				if three == two || three == one {
					continue
				}
				for four := 0; four < 5; four++ {
					if four == three || four == two || four == one {
						continue
					}
					for five := 0; five < 5; five++ {
						if five == four || five == three || five == two || five == one {
							continue
						}
						var input int
						phases := []int{one, two, three, four, five}
						for i := 0; i < 5; i++ {
							r := loadInput()
							r.inputs = append(r.inputs, phases[i], input)
							r.execOpcodes()
							input = r.outputs[len(r.outputs)-1]
						}
						if input > max {
							max = input
						}
					}
				}
			}
		}
	}
	fmt.Printf("Part 1 answer: %d\n", max)
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

func (r *commands) execOpcodes() {
	for r.currentloc < len(r.cmds) {
		cmd := r.cmds[r.currentloc]
		opcode := cmd % 100
		modes := []int{(cmd / 100) % 10, (cmd / 1000) % 10, (cmd / 10000) % 10}
		switch opcode {
		case 1: // 1st + 2param put in 3rd param, move ahead 4
			first := r.getVal(modes[0], r.currentloc+1)
			second := r.getVal(modes[1], r.currentloc+2)
			r.setVal(modes[2], r.currentloc+3, first+second)
			r.currentloc += 4
		case 2: // 1st * 2param put in 3rd param, move ahead 4
			first := r.getVal(modes[0], r.currentloc+1)
			second := r.getVal(modes[1], r.currentloc+2)
			r.setVal(modes[2], r.currentloc+3, first*second)
			r.currentloc += 4
		case 3: // input
			r.setVal(modes[0], r.currentloc+1, r.inputs[0])
			r.currentloc += 2
			r.inputs = r.inputs[1:]
		case 4: // output
			r.outputs = append(r.outputs, r.getVal(modes[0], r.currentloc+1))
			r.currentloc += 2
		case 5: // jump if first param is non-zero
			if r.getVal(modes[0], r.currentloc+1) == 0 {
				r.currentloc += 3
				continue
			}
			r.currentloc = r.getVal(modes[1], r.currentloc+2)
		case 6: // jump if first param is zero
			if r.getVal(modes[0], r.currentloc+1) != 0 {
				r.currentloc += 3
				continue
			}
			r.currentloc = r.getVal(modes[1], r.currentloc+2)
		case 7: // store 1 if first < second else 0
			first := r.getVal(modes[0], r.currentloc+1)
			second := r.getVal(modes[1], r.currentloc+2)
			var val int
			if first < second {
				val = 1
			}
			r.setVal(modes[2], r.currentloc+3, val)
			r.currentloc += 4
		case 8: // store 1 if first == second else zero
			first := r.getVal(modes[0], r.currentloc+1)
			second := r.getVal(modes[1], r.currentloc+2)
			var val int
			if first == second {
				val = 1
			}
			r.setVal(modes[2], r.currentloc+3, val)
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

func (r *commands) getVal(mode, loc int) (val int) {
	switch mode {
	case 0:
		return r.cmds[r.cmds[loc]]
	case 1:
		return r.cmds[loc]
	default:
		log.Fatalf("Unhandled mode %d\n", mode)
		return -1
	}
}

func (r *commands) setVal(mode, loc, val int) {
	switch mode {
	case 0:
		r.cmds[r.cmds[loc]] = val
	case 1:
		r.cmds[loc] = val
	default:
		log.Fatalf("Unhandled mode %d\n", mode)
	}
}

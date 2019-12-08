package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input1 := getInput1()
	o := new(opCodes)
	o.regs = make([]int, 4)
	o.funcmap = make(map[int]func([]int))
	o.opmap = make(map[int]func([]int))
	o.counts = make(map[int]map[int]int)
	// initialize the inner maps for each opcode
	for i := 0; i < 16; i++ {
		o.counts[i] = make(map[int]int)
	}
	fmt.Printf("Part 1 answer: %d\n", o.part1(input1))
	input2 := getInput2()
	fmt.Printf("Part 2 answer: %d\n", o.part2(input2))
}

type testCode struct {
	before, command, after []int
}

func getInput1() (input1 []*testCode) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	getInts := regexp.MustCompile(`(\d+).+(\d+).+(\d+).+(\d+)`)

	var lines int
	test := new(testCode)
	for scn.Scan() {
		lines++
		switch lines % 4 {
		case 1:
			for _, i := range getInts.FindStringSubmatch(scn.Text())[1:] {
				test.before = append(test.before, atoi(i))
			}
		case 2:
			for _, i := range getInts.FindStringSubmatch(scn.Text())[1:] {
				test.command = append(test.command, atoi(i))
			}
		case 3:
			for _, i := range getInts.FindStringSubmatch(scn.Text())[1:] {
				test.after = append(test.after, atoi(i))
			}
		case 0:
			input1 = append(input1, test)
			test = new(testCode)
		}
	}
	return input1
}

func getInput2() (input2 [][]int) {
	f, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	getInts := regexp.MustCompile(`(\d+).+(\d+).+(\d+).+(\d+)`)

	for scn.Scan() {
		var line []int
		matches := getInts.FindStringSubmatch(scn.Text())[1:]
		for _, i := range matches {
			line = append(line, atoi(i))
		}
		input2 = append(input2, line)
	}
	return input2
}

func (o *opCodes) part1(inp []*testCode) (ans int) {
	for _, test := range inp {
		ans += o.testOpcode(test)
	}
	return ans
}

func (o *opCodes) part2(inp2 [][]int) (ans int) {
	for len(o.counts) > 0 {
		for k, mAp := range o.counts {
			if len(mAp) > 1 {
				continue
			}
			var id int
			for funcid := range mAp {
				id = funcid
			}
			opfunc := o.funcmap[id]
			if _, ok := o.opmap[k]; ok {
				fmt.Printf("Error opfunc %p already allocated %+v\n", opfunc, o.opmap)
			}
			o.opmap[k] = opfunc
			delete(o.counts, k)
			for _, cleanup := range o.counts {
				delete(cleanup, id)
			}
		}
	}
	o.regs = make([]int, 4)
	for _, cmd := range inp2 {
		f := o.opmap[cmd[0]]
		f(cmd)
	}
	return o.regs[0]
}

func (o *opCodes) testOpcode(test *testCode) (ans int) {
	var pass int
	for i, testOp := range []func([]int){
		o.addr,
		o.addi,
		o.mulr,
		o.muli,
		o.banr,
		o.bani,
		o.borr,
		o.bori,
		o.setr,
		o.seti,
		o.gtir,
		o.gtri,
		o.gtrr,
		o.eqir,
		o.eqri,
		o.eqrr,
	} {
		copy(o.regs, test.before)
		testOp(test.command)
		if fmt.Sprint(test.after) == fmt.Sprint(o.regs) {
			o.counts[test.command[0]][i]++
			o.funcmap[i] = testOp
			pass++
		}
	}

	if pass >= 3 {
		return 1
	}
	return 0
}

// unsafe but gaurenteed fine by regexp
func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

type opCodes struct {
	regs []int
	// counts maps an opcode to a map of func by id to count
	counts map[int]map[int]int
	// funcmap maps a func id to a func
	funcmap map[int]func([]int)
	opmap   map[int]func([]int)
}

func (o *opCodes) addr(t []int) {
	o.regs[t[3]] = o.regs[t[1]] + o.regs[t[2]]
}

func (o *opCodes) addi(t []int) {
	o.regs[t[3]] = o.regs[t[1]] + t[2]
}

func (o *opCodes) mulr(t []int) {
	o.regs[t[3]] = o.regs[t[1]] * o.regs[t[2]]
}

func (o *opCodes) muli(t []int) {
	o.regs[t[3]] = o.regs[t[1]] * t[2]
}

func (o *opCodes) banr(t []int) {
	o.regs[t[3]] = o.regs[t[1]] & o.regs[t[2]]
}

func (o *opCodes) bani(t []int) {
	o.regs[t[3]] = o.regs[t[1]] & t[2]
}

func (o *opCodes) borr(t []int) {
	o.regs[t[3]] = o.regs[t[1]] | o.regs[t[2]]
}

func (o *opCodes) bori(t []int) {
	o.regs[t[3]] = o.regs[t[1]] | t[2]
}

func (o *opCodes) setr(t []int) {
	o.regs[t[3]] = o.regs[t[1]]
}

func (o *opCodes) seti(t []int) {
	o.regs[t[3]] = t[1]
}

func (o *opCodes) gtir(t []int) {
	if t[1] > o.regs[t[2]] {
		o.regs[t[3]] = 1
	} else {
		o.regs[t[3]] = 0
	}
}

func (o *opCodes) gtri(t []int) {
	if o.regs[t[1]] > t[2] {
		o.regs[t[3]] = 1
	} else {
		o.regs[t[3]] = 0
	}
}

func (o *opCodes) gtrr(t []int) {
	if o.regs[t[1]] > o.regs[t[2]] {
		o.regs[t[3]] = 1
	} else {
		o.regs[t[3]] = 0
	}
}

func (o *opCodes) eqir(t []int) {
	if t[1] == o.regs[t[2]] {
		o.regs[t[3]] = 1
	} else {
		o.regs[t[3]] = 0
	}
}

func (o *opCodes) eqri(t []int) {
	if o.regs[t[1]] == t[2] {
		o.regs[t[3]] = 1
	} else {
		o.regs[t[3]] = 0
	}
}

func (o *opCodes) eqrr(t []int) {
	if o.regs[t[1]] == o.regs[t[2]] {
		o.regs[t[3]] = 1
	} else {
		o.regs[t[3]] = 0
	}
}

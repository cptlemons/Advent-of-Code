package Day11

import (
	"strconv"
	"strings"
	"testing"
)

type Octopi struct {
	Grid    [][]int
	Flashes int
}

func parseInput(input string) *Octopi {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			grid[i] = append(grid[i], n)
		}
	}
	return &Octopi{Grid: grid}
}

func (o *Octopi) Print() {
	for _, row := range o.Grid {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}

func (o *Octopi) Step() {
	for i, row := range o.Grid {
		for j := range row {
			o.Grid[i][j]++
			if o.Grid[i][j] > 9 {
				o.Grid[i][j] = -1000
				o.Flashes++
				o.PropagateFlash(i, j)
			}
		}
	}
	for i, row := range o.Grid {
		for j := range row {
			if o.Grid[i][j] < 0 {
				o.Grid[i][j] = 0
			}
		}
	}
	return
}

func (o *Octopi) PropagateFlash(i, j int) {
	if j-1 >= 0 {
		o.Grid[i][j-1]++
		if o.Grid[i][j-1] > 9 {
			o.Grid[i][j-1] = -1000
			o.Flashes++
			o.PropagateFlash(i, j-1)
		}
	}
	if j+1 < len(o.Grid[i]) {
		o.Grid[i][j+1]++
		if o.Grid[i][j+1] > 9 {
			o.Grid[i][j+1] = -1000
			o.Flashes++
			o.PropagateFlash(i, j+1)
		}
	}
	if i-1 >= 0 {
		if j-1 >= 0 {
			o.Grid[i-1][j-1]++
			if o.Grid[i-1][j-1] > 9 {
				o.Grid[i-1][j-1] = -1000
				o.Flashes++
				o.PropagateFlash(i-1, j-1)
			}
		}
		if j+1 < len(o.Grid[i]) {
			o.Grid[i-1][j+1]++
			if o.Grid[i-1][j+1] > 9 {
				o.Grid[i-1][j+1] = -1000
				o.Flashes++
				o.PropagateFlash(i-1, j+1)
			}
		}
		o.Grid[i-1][j]++
		if o.Grid[i-1][j] > 9 {
			o.Grid[i-1][j] = -1000
			o.Flashes++
			o.PropagateFlash(i-1, j)
		}
	}
	if i+1 < len(o.Grid) {
		if j-1 >= 0 {
			o.Grid[i+1][j-1]++
			if o.Grid[i+1][j-1] > 9 {
				o.Grid[i+1][j-1] = -1000
				o.Flashes++
				o.PropagateFlash(i+1, j-1)
			}
		}
		if j+1 < len(o.Grid[i]) {
			o.Grid[i+1][j+1]++
			if o.Grid[i+1][j+1] > 9 {
				o.Grid[i+1][j+1] = -1000
				o.Flashes++
				o.PropagateFlash(i+1, j+1)
			}
		}
		o.Grid[i+1][j]++
		if o.Grid[i+1][j] > 9 {
			o.Grid[i+1][j] = -1000
			o.Flashes++
			o.PropagateFlash(i+1, j)
		}
	}
	return
}

func Part1(input string) int {
	octopi := parseInput(input)
	octopi.Print()
	steps := 100
	for steps > 0 {
		octopi.Step()
		steps--
	}
	return octopi.Flashes
}

func Part2(input string) int {
	octopi := parseInput(input)
	octopi.Print()
	for i := 1; i < 1000; i++ {
		octopi.Step()
		if octopi.Flashes == 100 {
			return i
		}
		octopi.Flashes = 0
	}
	return -1
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
			want: 1656,
		},
		{
			name: "real",
			input: `3113284886
2851876144
2774664484
6715112578
7146272153
6256656367
3148666245
3857446528
7322422833
8152175168`,
			want: 1705,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Part1(test.input); got != test.want {
				t.Errorf("Part1() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
			want: 195,
		},
		{
			name: "real",
			input: `3113284886
2851876144
2774664484
6715112578
7146272153
6256656367
3148666245
3857446528
7322422833
8152175168`,
			want: 265,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Part2(test.input); got != test.want {
				t.Errorf("Part2() = %v, want %v", got, test.want)
			}
		})
	}
}

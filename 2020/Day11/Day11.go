package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
}

func loadInput(file string) (seats []string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		line := scn.Text()
		seats = append(seats, line)
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}
	return seats
}

func checkSeat(seats []string, y int, x int) (ocp int) {
	for c := -1; c <= 1; c++ {
		for r := -1; r <= 1; r++ {
			if y+r < 0 || y+r >= len(seats) || x+c < 0 || x+c >= len(seats[0]) {
				continue
			}
			if r == 0 && c == 0 {
				continue
			}
			if seats[y+r][x+c] == '#' {

				ocp++
			}
		}
	}
	return ocp
}

func part1(seats []string) (ocp int) {
	for {
		var change bool
		var newSeats []string
		newSeats = append(newSeats, seats...)
		for r, row := range seats {
			for c, seat := range row {
				if seat == '.' {
					continue
				}
				ocp := checkSeat(seats, r, c)
				if seat == '#' && ocp >= 4 {
					newSeats[r] = newSeats[r][:c] + "L" + newSeats[r][c+1:]
					change = true
				}
				if seat == 'L' && ocp == 0 {
					newSeats[r] = newSeats[r][:c] + "#" + newSeats[r][c+1:]
					change = true
				}
			}
		}
		seats = newSeats

		if !change {
			break
		}
	}
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				ocp++
			}
		}
	}
	return ocp
}

func checkSeat2(seats []string, y int, x int) (ocp int) {
	dirs := [][]int{{-1, -1}, {-1, 0}, {0, -1}, {1, 1}, {1, 0}, {0, 1}, {-1, 1}, {1, -1}}
	for _, dir := range dirs {
		for i := 1; i < 1000; i++ {
			r := y + i*dir[0]
			c := x + i*dir[1]
			if r < 0 || r >= len(seats) || c < 0 || c >= len(seats[0]) {
				continue
			}
			seen := seats[r][c]
			if seen == '.' {
				continue
			}
			if seen == 'L' {
				break
			}
			if seen == '#' {
				ocp++
				break
			}
		}
	}
	return ocp
}

func part2(seats []string) (ocp int) {
	for {
		var change bool
		var newSeats []string
		newSeats = append(newSeats, seats...)
		for r, row := range seats {
			for c, seat := range row {
				if seat == '.' {
					continue
				}
				ocp := checkSeat2(seats, r, c)
				if seat == '#' && ocp >= 5 {
					newSeats[r] = newSeats[r][:c] + "L" + newSeats[r][c+1:]
					change = true
				}
				if seat == 'L' && ocp == 0 {
					newSeats[r] = newSeats[r][:c] + "#" + newSeats[r][c+1:]
					change = true
				}
			}
		}
		seats = newSeats

		if !change {
			break
		}
	}
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				ocp++
			}
		}
	}
	return ocp
}

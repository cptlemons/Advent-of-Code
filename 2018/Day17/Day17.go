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
	input := loadInput()
	part1(input)
}

type layout struct {
	veins                  []*vein
	drawing                [][]string
	xmin, xmax, ymin, ymax int
	activeWater            []*coord
}

type coord struct {
	x, y int
}

type vein struct {
	xStart, xEnd, yStart, yEnd int
}

func loadInput() (input *layout) {
	f, err := os.Open("sampleinput.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s\n", err)
	}
	defer f.Close()

	// x=495, y=2..7
	// y=7, x=495..501
	getVein := regexp.MustCompile(`([xy])=(\d+), [xy]=(\d+)..(\d+)`)

	scn := bufio.NewScanner(f)
	veins := []*vein{}
	for scn.Scan() {
		text := getVein.FindStringSubmatch(scn.Text())
		if text[1] == "x" {
			veins = append(veins, &vein{
				xStart: atoi(text[2]),
				xEnd:   atoi(text[2]),
				yStart: atoi(text[3]),
				yEnd:   atoi(text[4]),
			})
		} else {
			veins = append(veins, &vein{
				yStart: atoi(text[2]),
				yEnd:   atoi(text[2]),
				xStart: atoi(text[3]),
				xEnd:   atoi(text[4]),
			})
		}
	}
	return &layout{veins: veins}
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func part1(l *layout) {
	l.getEdges()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		l.drawMap()
		l.stepWater()
	}
}

func (l *layout) getEdges() {
	xmin := 10000000
	xmax, ymax := -10000000, -10000000
	for _, vein := range l.veins {
		if vein.xStart < xmin {
			xmin = vein.xStart
		}
		if vein.xEnd > xmax {
			xmax = vein.xEnd
		}
		if vein.yEnd > ymax {
			ymax = vein.yEnd
		}
	}
	// ymin is always zero becauase the fountain starts at 500, 0
	l.xmin, l.xmax, l.ymin, l.ymax = xmin, xmax, 0, ymax
}

func (l *layout) drawMap() {
	l.drawing = [][]string{}
	for r := 0; r <= l.ymax+1; r++ {
		l.drawing = append(l.drawing, []string{})
		for c := l.xmin - 1; c <= l.xmax+1; c++ {
			if r == 0 && c == 500 {
				l.drawing[r] = append(l.drawing[r], "+")
				l.activeWater = append(l.activeWater, &coord{x: c, y: r})
			} else {
				l.drawing[r] = append(l.drawing[r], ".")
			}
		}
	}

	for _, vein := range l.veins {
		for r := vein.yStart; r <= vein.yEnd; r++ {
			for c := vein.xStart - l.xmin + 1; c <= vein.xEnd-l.xmin+1; c++ {
				l.drawing[r][c] = "#"
			}
		}
	}
	for _, r := range l.drawing {
		fmt.Println(r)
	}
}

func (l *layout) stepWater() {
	var newPoints []*coord
	for _, points := range l.activeWater {
		l.activeWater = l.activeWater[1:]
		if l.drawing[points.y+1][points.x-l.xmin+1] == "." {
			fmt.Println(points.y+1, points.x-l.xmin+1)
			l.drawing[points.y+1][points.x-l.xmin+1] = "|"
			newPoints = append(newPoints, &coord{x: points.x, y: points.y + 1})
		}
	}
	l.activeWater = newPoints
}

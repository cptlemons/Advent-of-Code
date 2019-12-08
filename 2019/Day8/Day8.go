package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Println("Part 2 answer:")
	part2(input)
}

func getInput() (input string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)

	_ = regexp.MustCompile(`position=< *(-?\d*), *(-?\d*)> velocity=< *(-?\d*), *(-?\d*)>`)

	for scn.Scan() {
		input = scn.Text()
	}

	return input
}

func part1(input string) (ans int) {
	var layers [][]string
	var layer []string
	var line string
	for i := 0; i < len(input); i++ {
		line += input[i : i+1]
		if len(line) == 25 {
			layer = append(layer, line)
			line = ""
		}
		if len(layer) == 6 {
			layers = append(layers, layer)
			layer = []string{}
		}
	}
	numCount := make(map[int]map[string]int)
	for i := 0; i < len(layers); i++ {
		numCount[i] = make(map[string]int)
		for j := 0; j < len(layers[i]); j++ {
			for l := 0; l < len(layers[i][j]); l++ {
				numCount[i][layers[i][j][l:l+1]]++
			}
		}
	}
	fmt.Println(numCount)
	minZeros := 10000
	var bestLayer int
	for k, v := range numCount {
		if v["0"] < minZeros {
			minZeros = v["0"]
			bestLayer = k
		}
	}
	return numCount[bestLayer]["1"] * numCount[bestLayer]["2"]
}

func part2(input string) {
	var layers [][]string
	var layer []string
	var line string
	for i := 0; i < len(input); i++ {
		line += input[i : i+1]
		if len(line) == 25 {
			layer = append(layer, line)
			line = ""
		}
		if len(layer) == 6 {
			layers = append(layers, layer)
			layer = []string{}
		}
	}

	var image string
	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			var stack int
		nextStack:
			switch char := layers[stack][y][x : x+1]; char {
			case "0":
				image += " "
			case "1":
				image += char
			case "2":
				stack++
				goto nextStack
			}
		}
		image += "\n"
	}
	fmt.Println(image)
}

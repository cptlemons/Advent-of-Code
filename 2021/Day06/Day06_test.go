package Day06

import (
	"testing"
)

func Part1(fish []int) int {
	daysToFish := make([]int, 9)

	for _, days := range fish {
		daysToFish[days]++
	}

	day := 0
	for day < 80 {
		newDaystoFish := make([]int, 9)
		for i, fish := range daysToFish {
			if i == 0 {
				newDaystoFish[6] = fish
				newDaystoFish[8] = fish
			} else {
				newDaystoFish[i-1] += fish
			}
		}
		daysToFish = newDaystoFish
		day++
	}
	return countFish(daysToFish)
}

func countFish(fish []int) int {
	count := 0
	for _, f := range fish {
		count += f
	}
	return count
}

func Part2(fish []int) int {
	daysToFish := make([]int, 9)

	for _, days := range fish {
		daysToFish[days]++
	}

	day := 0
	for day < 256 {
		newDaystoFish := make([]int, 9)
		for i, fish := range daysToFish {
			if i == 0 {
				newDaystoFish[6] = fish
				newDaystoFish[8] = fish
			} else {
				newDaystoFish[i-1] += fish
			}
		}
		daysToFish = newDaystoFish
		day++
	}
	return countFish(daysToFish)
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example",
			input: []int{3, 4, 3, 1, 2},
			want:  5934,
		},
		{
			name:  "Real",
			input: []int{3, 5, 2, 5, 4, 3, 2, 2, 3, 5, 2, 3, 2, 2, 2, 2, 3, 5, 3, 5, 5, 2, 2, 3, 4, 2, 3, 5, 5, 3, 3, 5, 2, 4, 5, 4, 3, 5, 3, 2, 5, 4, 1, 1, 1, 5, 1, 4, 1, 4, 3, 5, 2, 3, 2, 2, 2, 5, 2, 1, 2, 2, 2, 2, 3, 4, 5, 2, 5, 4, 1, 3, 1, 5, 5, 5, 3, 5, 3, 1, 5, 4, 2, 5, 3, 3, 5, 5, 5, 3, 2, 2, 1, 1, 3, 2, 1, 2, 2, 4, 3, 4, 1, 3, 4, 1, 2, 2, 4, 1, 3, 1, 4, 3, 3, 1, 2, 3, 1, 3, 4, 1, 1, 2, 5, 1, 2, 1, 2, 4, 1, 3, 2, 1, 1, 2, 4, 3, 5, 1, 3, 2, 1, 3, 2, 3, 4, 5, 5, 4, 1, 3, 4, 1, 2, 3, 5, 2, 3, 5, 2, 1, 1, 5, 5, 4, 4, 4, 5, 3, 3, 2, 5, 4, 4, 1, 5, 1, 5, 5, 5, 2, 2, 1, 2, 4, 5, 1, 2, 1, 4, 5, 4, 2, 4, 3, 2, 5, 2, 2, 1, 4, 3, 5, 4, 2, 1, 1, 5, 1, 4, 5, 1, 2, 5, 5, 1, 4, 1, 1, 4, 5, 2, 5, 3, 1, 4, 5, 2, 1, 3, 1, 3, 3, 5, 5, 1, 4, 1, 3, 2, 2, 3, 5, 4, 3, 2, 5, 1, 1, 1, 2, 2, 5, 3, 4, 2, 1, 3, 2, 5, 3, 2, 2, 3, 5, 2, 1, 4, 5, 4, 4, 5, 5, 3, 3, 5, 4, 5, 5, 4, 3, 5, 3, 5, 3, 1, 3, 2, 2, 1, 4, 4, 5, 2, 2, 4, 2, 1, 4},
			want:  343441,
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
		input []int
		want  int
	}{
		{
			name:  "Example",
			input: []int{3, 4, 3, 1, 2},
			want:  26984457539,
		},
		{
			name:  "Real",
			input: []int{3, 5, 2, 5, 4, 3, 2, 2, 3, 5, 2, 3, 2, 2, 2, 2, 3, 5, 3, 5, 5, 2, 2, 3, 4, 2, 3, 5, 5, 3, 3, 5, 2, 4, 5, 4, 3, 5, 3, 2, 5, 4, 1, 1, 1, 5, 1, 4, 1, 4, 3, 5, 2, 3, 2, 2, 2, 5, 2, 1, 2, 2, 2, 2, 3, 4, 5, 2, 5, 4, 1, 3, 1, 5, 5, 5, 3, 5, 3, 1, 5, 4, 2, 5, 3, 3, 5, 5, 5, 3, 2, 2, 1, 1, 3, 2, 1, 2, 2, 4, 3, 4, 1, 3, 4, 1, 2, 2, 4, 1, 3, 1, 4, 3, 3, 1, 2, 3, 1, 3, 4, 1, 1, 2, 5, 1, 2, 1, 2, 4, 1, 3, 2, 1, 1, 2, 4, 3, 5, 1, 3, 2, 1, 3, 2, 3, 4, 5, 5, 4, 1, 3, 4, 1, 2, 3, 5, 2, 3, 5, 2, 1, 1, 5, 5, 4, 4, 4, 5, 3, 3, 2, 5, 4, 4, 1, 5, 1, 5, 5, 5, 2, 2, 1, 2, 4, 5, 1, 2, 1, 4, 5, 4, 2, 4, 3, 2, 5, 2, 2, 1, 4, 3, 5, 4, 2, 1, 1, 5, 1, 4, 5, 1, 2, 5, 5, 1, 4, 1, 1, 4, 5, 2, 5, 3, 1, 4, 5, 2, 1, 3, 1, 3, 3, 5, 5, 1, 4, 1, 3, 2, 2, 3, 5, 4, 3, 2, 5, 1, 1, 1, 2, 2, 5, 3, 4, 2, 1, 3, 2, 5, 3, 2, 2, 3, 5, 2, 1, 4, 5, 4, 4, 5, 5, 3, 3, 5, 4, 5, 5, 4, 3, 5, 3, 5, 3, 1, 3, 2, 2, 1, 4, 4, 5, 2, 2, 4, 2, 1, 4},
			want:  1569108373832,
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

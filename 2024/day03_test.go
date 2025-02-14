package year2024

import (
	"github.com/mattkoler/Advent-of-Code/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func Day03p1() int {
	lines := utils.InputByLines("day03_input.txt")

	mulregex := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var total int

	for line := range lines {
		hits := mulregex.FindAllString(line, -1)
		for _, hit := range hits {
			hit = strings.TrimPrefix(hit, "mul(")
			hit = strings.TrimSuffix(hit, ")")
			nums := strings.Split(hit, ",")
			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatalf("converting %q to int: %s", nums[0], err)
			}
			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatalf("converting %q to int: %s", nums[1], err)
			}
			if num1 > 999 || num2 > 999 {
				continue
			}
			total += num1 * num2
		}
	}
	return total
}

func Day03p2() int {
	lines := utils.InputByLines("day03_input.txt")

	mulregex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	doregex := regexp.MustCompile(`do\(\)`)
	dontregex := regexp.MustCompile(`don't\(\)`)

	var total int

	var allLines string
	for line := range lines {
		allLines += line
	}

	var cleanedLines string
	for {
		r := dontregex.FindStringIndex(allLines)
		if r == nil {
			cleanedLines += allLines
			break
		}
		cleanedLines += allLines[:r[0]]
		allLines = allLines[r[1]:]

		l := doregex.FindStringIndex(allLines)
		if l == nil {
			break
		}
		allLines = allLines[l[1]:]
	}

	hits := mulregex.FindAllString(cleanedLines, -1)
	for _, hit := range hits {
		hit = strings.TrimPrefix(hit, "mul(")
		hit = strings.TrimSuffix(hit, ")")
		nums := strings.Split(hit, ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("converting %q to int: %s", nums[0], err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("converting %q to int: %s", nums[1], err)
		}
		if num1 > 999 || num2 > 999 {
			continue
		}
		total += num1 * num2
	}

	return total
}

func TestDay03(t *testing.T) {
	t.Run("Day03p1", func(t *testing.T) {
		want := 187833789
		got := Day03p1()
		if got != want {
			t.Errorf("Day03p1() = %v, want %v", got, want)
		}
	})
	t.Run("Day03p2", func(t *testing.T) {
		want := 94455185
		got := Day03p2()
		if got != want {
			t.Errorf("Day03p2() = %v, want %v", got, want)
		}
	})
}

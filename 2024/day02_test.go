package year2024

import (
	"github.com/mattkoler/Advent-of-Code/utils"
	"strings"
	"testing"
)

func Day02p1() int {
	lines := utils.InputByLines("day02_input.txt")

	var safeCount int
	for line := range lines {
		parts := strings.Fields(line)
		nums := utils.StringSliceToInts(parts)
		if safeLine(nums) {
			safeCount++
		}
	}
	return safeCount
}

func Day02p2() int {
	lines := utils.InputByLines("day02_input.txt")

	var safeCount int
	for line := range lines {
		parts := strings.Fields(line)
		nums := utils.StringSliceToInts(parts)
		safe, i1, i2 := safeLineDetails(nums)
		if safe {
			safeCount++
			continue
		}
		n1 := append([]int{}, nums[:i1]...)
		n1 = append(n1, nums[i2:]...)
		if safeLine(n1) {
			safeCount++
			continue
		}
		n2 := append([]int{}, nums[:i2]...)
		n2 = append(n2, nums[i2+1:]...)
		if safeLine(n2) {
			safeCount++
			continue
		}
		if safeLine(nums[1:]) {
			safeCount++
			continue
		}
	}
	return safeCount
}

func safeLine(nums []int) bool {
	var increasing bool
	for i, num := range nums[:len(nums)-1] {
		if i == 0 && num < nums[i+1] {
			increasing = true
		}
		// nums must differ by at least 1 and at most 3 and must always be increasing or decreasing
		if num == nums[i+1] || utils.Abs(num-nums[i+1]) > 3 || (increasing && num > nums[i+1]) || (!increasing && num < nums[i+1]) {
			return false
		}
	}
	return true
}

func safeLineDetails(nums []int) (bool, int, int) {
	var increasing bool
	for i, num := range nums[:len(nums)-1] {
		if i == 0 && num < nums[i+1] {
			increasing = true
		}
		// nums must differ by at least 1 and at most 3 and must always be increasing or decreasing
		if num == nums[i+1] || utils.Abs(num-nums[i+1]) > 3 || (increasing && num > nums[i+1]) || (!increasing && num < nums[i+1]) {
			return false, i, i + 1
		}
	}
	return true, 0, 0
}

func TestDay02p1(t *testing.T) {
	want := 686
	if got := Day02p1(); got != want {
		t.Errorf("Day02p1() = %v, want %v", got, want)
	}
}

func TestDay02p2(t *testing.T) {
	want := 717
	if got := Day02p2(); got != want {
		t.Errorf("Day02p2() = %v, want %v", got, want)
	}
}

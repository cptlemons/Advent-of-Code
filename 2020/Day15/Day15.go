package main

func main() {}

func part1(intList []int) (ans int) {
	seenMap := make(map[int][]int)
	for i, n := range intList {
		seenMap[n] = append(seenMap[n], i)
	}
	var lastn int
	for len(intList) < 2020 {
		if v, ok := seenMap[lastn]; ok && len(v) > 1 {
			lastn = seenDiff(v)
			seenMap[lastn] = append(seenMap[lastn], len(intList))
			intList = append(intList, lastn)
		} else {
			lastn = 0
			seenMap[lastn] = append(seenMap[lastn], len(intList))
			intList = append(intList, lastn)
		}
	}
	return lastn
}

func seenDiff(seen []int) (diff int) {
	return seen[len(seen)-1] - seen[len(seen)-2]
}

func part2(intList []int) (ans int) {
	seenMap := make(map[int][]int)
	for i, n := range intList {
		seenMap[n] = append(seenMap[n], i)
	}
	var lastn int
	for len(intList) < 30000000 {
		if v, ok := seenMap[lastn]; ok && len(v) > 1 {
			lastn = seenDiff(v)
			seenMap[lastn] = append(seenMap[lastn], len(intList))
			intList = append(intList, lastn)
		} else {
			lastn = 0
			seenMap[lastn] = append(seenMap[lastn], len(intList))
			intList = append(intList, lastn)
		}
	}
	return lastn
}

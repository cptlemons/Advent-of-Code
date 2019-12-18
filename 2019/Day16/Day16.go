package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "59772698208671263608240764571860866740121164692713197043172876418614411671204569068438371694198033241854293277505547521082227127768000396875825588514931816469636073669086528579846568167984238468847424310692809356588283194938312247006770713872391449523616600709476337381408155057994717671310487116607321731472193054148383351831456193884046899113727301389297433553956552888308567897333657138353770191097676986516493304731239036959591922009371079393026332649558536888902303554797360691183681625604439250088062481052510016157472847289467410561025668637527408406615316940050060474260802000437356279910335624476330375485351373298491579364732029523664108987"

	fmt.Println(part1(input))
	input = "03036732577212944063491565474664"
	fmt.Println(part2(input))
}

func part1(inp string) (ans string) {
	pattern := []int{0, 1, 0, -1}
	for r := 1; r <= 100; r++ {
		var transition string
		for p := 1; p <= len(inp); p++ {
			var sum, mult, idx int
			for i := 0; i < len(inp); i++ {
				idx = ((i + 1) / p) % 4
				if idx == 0 || idx == 2 {
					continue
				}
				mult = pattern[idx]
				sum += (int(inp[i]) - '0') * mult
			}
			transition += strconv.Itoa(abs(sum % 10))
		}
		inp = transition
	}
	return inp[:8]
}

func part2(input string) (ans string) {
	inp := []byte(strings.Repeat(input, 10000))
	for i := 0; i < 100; i++ {
		for i := len(inp) - 1; i > 0; i-- {
			sum := int(inp[i] - '0')
			if i != len(inp)-1 {
				sum += int(inp[i-1] - '0')
			}
			inp[i] = byte(sum%10 + '0')
		}
	}
	for _, char := range inp[0303673 : 0303673+8] {
		ans += string(char)
	}
	return ans
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

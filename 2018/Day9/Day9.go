package main

import (
	"container/ring"
	"fmt"
)

func main() {
	players := 423
	finalmarble := 71944
	fmt.Printf("Part 1 answer: %d\n", part1(players, finalmarble))
	fmt.Printf("Part 1 circle answer: %d\n", part2(players, finalmarble))
	fmt.Printf("Part 2 circle answer: %d\n", part2(players, finalmarble*100))
}

func part1(players, finalmarble int) (ans int) {
	turn := 2
	circle := []int{0, 1}
	currentloc := 1
	// maps a player to their score
	scores := make(map[int]int)
	for turn <= finalmarble {
		if turn%71944 == 0 {
			fmt.Println("Turn", turn)
		}
		var newcircle []int
		if turn%23 == 0 {
			removeloc := currentloc - 7
			if removeloc < 0 {
				removeloc = len(circle) + removeloc
			}
			scores[turn%players] += turn + circle[removeloc]
			if removeloc == len(circle) {
				currentloc = 0
				circle = circle[:len(circle)-1]
			} else {
				currentloc = removeloc
				newcircle = append(newcircle, circle[:removeloc]...)
				newcircle = append(newcircle, circle[removeloc+1:]...)
				circle = newcircle
			}
			turn++
			continue
		}
		startloc := (currentloc + 1) % len(circle)
		endloc := (currentloc + 2) % len(circle)
		if startloc != len(circle) {
			newcircle = append(newcircle, circle[:startloc+1]...)
		}
		currentloc = len(newcircle)
		newcircle = append(newcircle, turn)
		if endloc != 0 {
			newcircle = append(newcircle, circle[endloc:]...)
		}
		circle = newcircle
		turn++
	}

	for _, score := range scores {
		if score > ans {
			ans = score
		}
	}
	return ans
}

func part2(players, finalmarble int) (ans int) {
	circle := ring.New(1)
	circle.Value = 0
	scores := make(map[int]int)

	for turn := 1; turn <= finalmarble; turn++ {
		if turn%23 == 0 {
			circle = circle.Move(-8)
			rem := circle.Unlink(1)
			scores[turn%players] += turn + rem.Value.(int)
		} else {
			circle = circle.Next()
			circle.Link(turnToRing(turn))
		}
		circle = circle.Next()
	}
	for _, score := range scores {
		if score > ans {
			ans = score
		}
	}
	return ans
}

func turnToRing(t int) *ring.Ring {
	new := ring.New(1)
	new.Value = t
	return new
}

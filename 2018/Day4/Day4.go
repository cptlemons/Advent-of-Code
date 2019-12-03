package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	fmt.Printf("Part 1 answer: %d\n", part1(input))
	fmt.Printf("Part 2 answer: %d\n", part2(input))
}

type guard struct {
	sleepMins int
	awakeMins int
	schedule  map[string]day
}

type day struct {
	awake  [60]int
	asleep [60]int
}

func getInput() (input map[int]guard) {
	input = make(map[int]guard)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	scn := bufio.NewScanner(f)
	var events []string
	for scn.Scan() {
		events = append(events, scn.Text())
	}
	sort.Strings(events)

	getID := regexp.MustCompile(`Guard #(\d*) begins shift`)
	getDate := regexp.MustCompile(`\[(\d*-\d*-\d*)`)
	getTime := regexp.MustCompile(`:(\d*)`)
	getAction := regexp.MustCompile(`(wakes|asleep)`)
	var currentGuard, currentTime int
	var prevEvent string
	awake := true
	for _, event := range events {
		if strings.Contains(event, "Guard") {
			if currentGuard != 0 {
				g := input[currentGuard]
				if g.schedule == nil {
					g.schedule = make(map[string]day)
				}
				eventDate := getDate.FindStringSubmatch(prevEvent)[1]
				eventTime := atoi(getTime.FindStringSubmatch(prevEvent)[1])
				cday := g.schedule[eventDate]
				if awake {
					for i := currentTime; i < eventTime && i < 60; i++ {
						g.awakeMins++
						cday.awake[i] = 1
						currentTime++
					}
				} else {
					for i := currentTime; i < eventTime && i < 60; i++ {
						g.sleepMins++
						cday.asleep[i] = 1
						currentTime++
					}
				}
				g.schedule[eventDate] = cday
				input[currentGuard] = g
			}
			currentGuard = atoi(getID.FindStringSubmatch(event)[1])
			currentTime = 0
			continue
		}
		g := input[currentGuard]
		if g.schedule == nil {
			g.schedule = make(map[string]day)
		}
		eventDate := getDate.FindStringSubmatch(event)[1]
		eventTime := atoi(getTime.FindStringSubmatch(event)[1])
		action := getAction.FindStringSubmatch(event)[1]
		cday := g.schedule[eventDate]
		switch action {
		case "wakes":
			for i := currentTime; i < eventTime && i < 60; i++ {
				g.sleepMins++
				cday.asleep[i] = 1
				currentTime++
				awake = true
			}
			g.schedule[eventDate] = cday
		case "asleep":
			for i := currentTime; i < eventTime && i < 60; i++ {
				g.awakeMins++
				cday.awake[i] = 1
				currentTime++
				awake = false
			}
			g.schedule[eventDate] = cday
		default:
			fmt.Printf("Unknown action %s in event %s\n", action, event)
		}
		input[currentGuard] = g
		prevEvent = event
	}
	return input
}

func atoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func part1(input map[int]guard) (ans int) {
	var minutesAsleep, minute, guard int
	// find the guard who sleeps the most
	for id, g := range input {
		if g.sleepMins > minutesAsleep {
			guard = id
			minutesAsleep = g.sleepMins
		}
	}
	// map when the guard is sleeping
	sleepCount := make(map[int]int)
	for _, d := range input[guard].schedule {
		for i := 0; i < len(d.asleep); i++ {
			sleepCount[i] += d.asleep[i]
		}
	}
	// find out on which minute the guard slept the most
	var count int
	for min, tcount := range sleepCount {
		if tcount > count {
			minute = min
			count = tcount
		}
	}
	return minute * guard
}

func part2(input map[int]guard) (ans int) {
	var count, minute, guard int
	for id, g := range input {
		// map of minute to count of days slept on that minute
		sleepCount := make(map[int]int)
		for _, d := range g.schedule {
			for i := 0; i < len(d.asleep); i++ {
				if d.asleep[i] == 1 {
					sleepCount[i] += d.asleep[i]
				}
			}
		}
		// figure out if the guard has slept more on a minute than we have seen
		for min, tcount := range sleepCount {
			if tcount > count {
				count = tcount
				minute = min
				guard = id
			}
		}
	}
	return minute * guard
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Passport struct {
	byr   string
	iyr   string
	eyr   string
	hgt   string
	hcl   string
	ecl   string
	pid   string
	cid   string
	valid bool
}

func main() {

}

func loadInput(file string) (pprts []Passport) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
		os.Exit(1)
	}

	scn := bufio.NewScanner(f)

	var rawpprt []string
	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			pprts = append(pprts, parsePassport(rawpprt))
			rawpprt = []string{}
			continue
		}
		split := strings.Split(line, " ")
		for _, splt := range split {
			rawpprt = append(rawpprt, splt)
		}
	}

	if err := scn.Err(); err != nil {
		fmt.Printf("Scanning error: %s", err)
		os.Exit(1)
	}
	return pprts
}

func parsePassport(info []string) (pprt Passport) {
	for _, kv := range info {
		split := strings.Split(kv, ":")
		switch split[0] {
		case "byr":
			pprt.byr = split[1]
		case "iyr":
			pprt.iyr = split[1]
		case "eyr":
			pprt.eyr = split[1]
		case "hgt":
			pprt.hgt = split[1]
		case "hcl":
			pprt.hcl = split[1]
		case "ecl":
			pprt.ecl = split[1]
		case "pid":
			pprt.pid = split[1]
		case "cid":
			pprt.cid = split[1]
		case "":
		default:
			fmt.Printf("unknown key: %s\n", split[0])
		}
	}
	return pprt
}

func part1(pprts []Passport) (valid int) {
	for _, pprt := range pprts {
		if pprt.byr == "" {
			continue
		}
		if pprt.iyr == "" {
			continue
		}
		if pprt.eyr == "" {
			continue
		}
		if pprt.hgt == "" {
			continue
		}
		if pprt.hcl == "" {
			continue
		}
		if pprt.ecl == "" {
			continue
		}
		if pprt.pid == "" {
			continue
		}
		valid++
	}
	return valid
}

func part2(pprts []Passport) (valid int) {
	regbyr := regexp.MustCompile(`^19[2-9][0-9]$|^200[0-2]$`)
	regiyr := regexp.MustCompile(`^201[0-9]$|^2020$`)
	regeyr := regexp.MustCompile(`^202[0-9]$|^2030$`)
	reghgt := regexp.MustCompile(`^1[5-8]\dcm$|^19[0-3]cm$|^59in$|^6\din$|7[0-6]in$`)
	reghcl := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	regecl := regexp.MustCompile(`^amb$|^blu$|^brn$|^gr[ny]$|^hzl$|^oth$`)
	regpid := regexp.MustCompile(`^\d{9}$`)
	for _, pprt := range pprts {
		if pprt.byr == "" || !regbyr.MatchString(pprt.byr) {
			continue
		}
		if pprt.iyr == "" || !regiyr.MatchString(pprt.iyr) {
			continue
		}
		if pprt.eyr == "" || !regeyr.MatchString(pprt.eyr) {
			continue
		}
		if pprt.hgt == "" || !reghgt.MatchString(pprt.hgt) {
			continue
		}
		if pprt.hcl == "" || !reghcl.MatchString(pprt.hcl) {
			continue
		}
		if pprt.ecl == "" || !regecl.MatchString(pprt.ecl) {
			continue
		}
		if pprt.pid == "" || !regpid.MatchString(pprt.pid) {
			continue
		}
		valid++
	}
	return valid
}

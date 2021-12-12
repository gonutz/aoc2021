package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	neighbors := func(node string) []string {
		var n []string
		for _, pair := range input {
			if pair[0] == node {
				n = append(n, pair[1])
			}
			if pair[1] == node {
				n = append(n, pair[0])
			}
		}
		return n
	}

	large := func(node string) bool {
		c := node[0]
		return 'A' <= c && c <= 'Z'
	}

	visitCounts := make(map[string]int)

	visits := func(node string) int {
		return visitCounts[node]
	}

	setVisits := func(node string, count int) {
		visitCounts[node] = count
	}

	pathCount := 0
	var follow func(string)
	follow = func(start string) {
		setVisits(start, visits(start)+1)
		if start == "end" {
			pathCount++
		} else {
			for _, n := range neighbors(start) {
				if large(n) || visits(n) == 0 {
					follow(n)
				}
			}
		}
		setVisits(start, visits(start)-1)
	}
	follow("start")

	fmt.Println("part 1)", pathCount)
}
func part2() {
	neighbors := func(node string) []string {
		var n []string
		for _, pair := range input {
			if pair[0] == node {
				n = append(n, pair[1])
			}
			if pair[1] == node {
				n = append(n, pair[0])
			}
		}
		return n
	}

	large := func(node string) bool {
		c := node[0]
		return 'A' <= c && c <= 'Z'
	}

	small := func(node string) bool {
		return !large(node)
	}

	visitCounts := make(map[string]int)

	visits := func(node string) int {
		return visitCounts[node]
	}

	setVisits := func(node string, count int) {
		visitCounts[node] = count
	}

	// not 105555, too high
	noSmallVisitedTwice := func() bool {
		for name, count := range visitCounts {
			if small(name) && count == 2 {
				return false
			}
		}
		return true
	}

	pathCount := 0
	var follow func(string)
	follow = func(start string) {
		setVisits(start, visits(start)+1)
		if start == "end" {
			pathCount++
		} else {
			for _, n := range neighbors(start) {
				if large(n) || visits(n) == 0 ||
					(n != "start" && visits(n) == 1 && noSmallVisitedTwice()) {
					follow(n)
				}
			}
		}
		setVisits(start, visits(start)-1)
	}
	follow("start")

	fmt.Println("part 2)", pathCount)
}

var input = [][2]string{
	{"ey", "dv"},
	{"AL", "ms"},
	{"ey", "lx"},
	{"zw", "YT"},
	{"hm", "zw"},
	{"start", "YT"},
	{"start", "ms"},
	{"dv", "YT"},
	{"hm", "ms"},
	{"end", "ey"},
	{"AL", "ey"},
	{"end", "hm"},
	{"rh", "hm"},
	{"dv", "ms"},
	{"AL", "dv"},
	{"ey", "SP"},
	{"hm", "lx"},
	{"dv", "start"},
	{"end", "lx"},
	{"zw", "AL"},
	{"hm", "AL"},
	{"lx", "zw"},
	{"ey", "zw"},
	{"zw", "dv"},
	{"YT", "ms"},
}

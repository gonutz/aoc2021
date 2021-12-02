package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	{
		lines := strings.Split(input, "\n")

		var horPos, depth int

		for _, line := range lines {
			var step *int

			if strings.HasPrefix(line, "forward") {
				line = strings.TrimPrefix(line, "forward ")
				step = &horPos
			} else if strings.HasPrefix(line, "up") {
				line = "-" + strings.TrimPrefix(line, "up ")
				step = &depth
			} else if strings.HasPrefix(line, "down") {
				line = strings.TrimPrefix(line, "down ")
				step = &depth
			} else {
				panic(line)
			}
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			*step += n
		}
		fmt.Println("part 1)", horPos*depth)
	}

	{
		lines := strings.Split(input, "\n")

		var horPos, depth, aim int

		for _, line := range lines {
			if strings.HasPrefix(line, "forward") {
				line = strings.TrimPrefix(line, "forward ")
				n, _ := strconv.Atoi(line)
				horPos += n
				depth += n * aim
			} else if strings.HasPrefix(line, "up") {
				line = strings.TrimPrefix(line, "up ")
				n, _ := strconv.Atoi(line)
				aim -= n
			} else if strings.HasPrefix(line, "down") {
				line = strings.TrimPrefix(line, "down ")
				n, _ := strconv.Atoi(line)
				aim += n
			}
		}
		fmt.Println("part 2)", horPos*depth)
	}
}

func ints(s string) []int {
	var i []int
	for _, line := range strings.Split(s, "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		i = append(i, n)
	}
	return i
}

const input = `forward 8
forward 3
forward 8
down 6
forward 3
up 6
down 3
down 8
down 5
down 1
down 4
up 4
forward 7
forward 7
down 8
forward 8
forward 2
forward 4
forward 6
down 2
forward 1
down 2
down 6
up 4
forward 4
forward 4
down 1
down 1
forward 8
down 9
forward 4
down 4
forward 9
down 9
forward 7
forward 4
forward 4
down 7
forward 9
up 5
forward 7
forward 4
up 5
up 6
down 9
forward 3
forward 3
forward 5
up 3
forward 8
down 9
forward 3
forward 5
down 2
forward 7
forward 8
forward 8
up 3
down 9
forward 2
forward 5
forward 4
up 1
down 3
forward 7
forward 4
up 3
down 8
up 8
forward 7
forward 7
down 2
down 7
forward 3
forward 1
down 7
down 4
down 9
forward 5
down 9
up 8
down 9
down 7
forward 2
down 5
down 6
forward 3
forward 2
forward 3
forward 2
down 4
up 1
forward 6
down 2
forward 2
down 5
up 1
up 8
forward 5
forward 6
up 8
down 4
forward 4
down 3
forward 5
forward 5
forward 8
forward 3
up 1
down 9
up 4
up 5
down 3
down 9
up 6
down 1
down 2
down 3
down 6
up 5
forward 9
forward 6
forward 6
forward 4
forward 7
up 6
down 2
down 9
down 9
forward 4
down 9
forward 2
down 4
down 8
down 3
up 1
down 2
down 7
down 1
down 5
up 5
forward 3
forward 2
down 2
up 4
up 9
forward 5
up 9
down 9
forward 1
down 4
forward 6
down 2
forward 7
forward 2
down 6
forward 9
down 9
down 8
forward 7
down 3
down 3
down 5
forward 5
up 9
up 5
down 1
forward 9
down 3
down 6
forward 3
up 4
up 1
down 3
down 1
up 6
forward 4
up 4
forward 5
down 3
down 5
forward 4
down 1
up 6
down 5
forward 1
forward 6
down 5
forward 1
forward 4
forward 2
down 1
down 3
down 9
down 9
down 8
forward 4
down 5
forward 6
up 6
down 3
forward 1
down 9
forward 4
down 2
up 5
down 4
forward 1
forward 2
down 6
forward 3
down 9
forward 6
forward 1
forward 7
up 1
up 7
up 2
forward 4
up 3
down 4
up 1
down 3
forward 1
down 5
up 6
forward 6
forward 6
up 3
up 2
down 8
down 2
down 2
down 7
forward 8
up 2
up 2
up 5
down 6
up 3
down 9
forward 7
down 7
forward 8
forward 2
forward 5
up 5
up 1
forward 5
up 2
up 3
up 2
up 3
down 3
down 2
forward 1
up 1
down 4
up 8
forward 5
down 9
forward 8
up 1
forward 7
forward 7
up 7
up 6
up 6
down 5
down 3
up 4
up 1
down 9
down 9
forward 4
down 7
forward 2
forward 8
forward 1
down 9
forward 2
forward 5
up 5
down 5
down 6
forward 1
down 4
forward 9
up 3
down 4
forward 7
forward 1
forward 3
forward 6
down 1
forward 3
up 5
up 6
down 3
forward 5
forward 8
forward 4
down 9
forward 2
forward 4
down 8
forward 6
down 7
up 7
forward 1
down 8
forward 5
forward 7
up 3
forward 7
forward 2
up 5
forward 3
down 4
up 5
down 6
up 6
forward 6
up 1
up 7
up 8
forward 4
down 6
down 8
up 8
down 2
forward 3
forward 9
down 9
down 6
down 2
up 4
forward 5
forward 6
forward 5
down 5
forward 5
down 2
down 9
down 3
down 4
forward 2
forward 7
down 7
down 2
down 4
down 8
up 6
down 7
forward 2
up 6
forward 6
down 4
up 8
forward 1
up 2
forward 6
forward 7
down 2
down 4
down 3
down 2
forward 5
down 4
forward 6
forward 6
down 8
forward 7
forward 1
forward 1
forward 4
forward 7
forward 8
down 2
forward 9
up 7
forward 1
down 1
forward 2
forward 2
up 4
forward 3
down 1
down 4
down 3
forward 8
forward 8
forward 9
forward 4
down 6
up 6
up 7
up 7
down 1
up 1
up 9
up 7
down 9
forward 4
up 8
down 5
up 3
down 2
forward 4
forward 6
up 7
up 2
down 5
down 3
down 9
forward 1
down 1
down 6
down 3
forward 9
down 3
up 6
up 4
up 6
down 6
up 7
down 5
up 4
up 3
up 3
up 2
forward 1
forward 5
forward 8
down 2
up 9
forward 6
forward 8
up 3
down 3
forward 7
forward 4
down 7
up 2
down 1
down 5
forward 8
down 6
down 2
down 9
down 5
up 7
up 5
forward 7
down 8
down 5
forward 7
up 5
up 1
down 4
forward 1
up 3
forward 7
forward 9
down 7
forward 4
down 9
down 2
up 2
forward 2
down 6
forward 3
down 5
up 8
forward 5
up 5
forward 1
down 4
forward 7
down 8
forward 7
up 2
forward 1
forward 2
down 7
forward 7
forward 5
up 3
down 8
forward 8
up 6
up 2
forward 9
down 7
up 6
down 1
up 5
forward 4
forward 9
forward 6
forward 3
forward 3
forward 1
forward 8
down 6
forward 3
up 2
up 6
down 3
down 7
forward 5
down 2
up 9
up 3
down 1
down 3
up 7
up 5
down 6
down 5
up 2
down 1
down 2
forward 1
down 1
forward 1
down 1
up 5
up 4
forward 1
down 7
forward 4
down 6
forward 2
forward 1
forward 4
forward 9
down 7
forward 7
down 7
down 5
forward 7
forward 3
forward 8
up 4
forward 9
down 1
down 9
forward 3
down 7
forward 1
forward 8
up 7
forward 5
down 8
forward 3
forward 6
forward 6
up 5
forward 7
up 3
down 9
forward 4
forward 4
forward 1
down 2
down 9
forward 8
forward 8
down 9
forward 5
up 4
down 6
forward 3
up 4
down 5
down 2
down 3
down 2
up 1
up 9
up 3
forward 5
forward 7
down 1
down 5
up 1
forward 8
down 5
forward 8
forward 8
down 2
forward 2
forward 7
forward 3
forward 6
up 9
down 3
forward 7
down 5
forward 3
up 1
down 3
down 9
forward 5
forward 5
up 5
down 3
down 3
up 1
forward 6
up 1
up 6
forward 5
down 3
down 9
forward 1
down 5
up 5
down 7
down 7
down 6
down 2
up 7
down 3
forward 2
up 6
down 2
forward 6
forward 9
down 6
down 4
down 5
down 7
forward 9
up 2
down 2
down 4
forward 4
down 9
forward 3
forward 8
forward 6
up 5
down 2
down 7
forward 7
up 6
down 4
up 8
forward 2
down 8
forward 7
up 8
up 5
up 8
down 1
forward 4
forward 7
down 5
forward 1
forward 5
down 9
down 6
up 8
up 5
down 7
down 4
forward 4
forward 2
forward 4
down 1
up 4
down 5
down 4
up 5
forward 2
up 3
down 9
down 1
down 4
up 1
up 7
down 5
forward 9
down 5
down 7
down 2
down 8
forward 7
down 5
down 6
forward 3
down 8
down 5
down 9
up 6
up 2
down 9
down 6
down 7
forward 8
down 8
forward 4
up 8
forward 3
down 2
up 2
forward 4
down 9
down 2
up 6
down 4
forward 6
down 7
forward 5
forward 4
down 3
up 6
forward 7
forward 1
up 6
down 6
forward 5
forward 3
down 8
up 8
up 8
forward 5
forward 1
up 1
forward 3
up 6
forward 2
down 8
forward 8
up 8
forward 1
forward 6
forward 8
up 7
up 3
forward 8
forward 5
down 4
down 4
forward 8
up 1
forward 5
down 9
forward 2
down 2
down 3
forward 1
down 2
up 3
down 1
up 1
up 1
forward 8
forward 5
forward 2
up 1
down 9
up 7
down 1
forward 9
up 4
forward 5
forward 5
forward 8
down 1
forward 3
forward 4
forward 8
up 2
forward 5
down 8
forward 5
up 7
forward 3
forward 2
forward 3
up 7
up 4
up 9
forward 7
forward 1
up 6
up 5
down 8
forward 3
down 5
forward 7
forward 3
up 9
forward 9
forward 7
up 2
down 7
forward 4
down 9
up 8
up 3
down 4
down 1
forward 4
up 3
down 6
down 2
forward 8
up 1
forward 5
up 7
down 5
forward 9
forward 6
forward 9
down 3
up 8
forward 6
forward 4
forward 2
forward 2
down 8
up 5
down 4
down 6
forward 3
forward 6
forward 1
forward 3
down 8
down 5
up 3
down 1
down 7
forward 8
forward 1
down 6
down 9
forward 2
up 5
down 6
up 6
down 5
down 8
forward 1
down 3
forward 1
forward 8
forward 2
down 4
forward 1
down 6
down 6
forward 3
up 7
forward 5
up 1
up 4
forward 7
forward 1
down 5
forward 2
down 1
forward 4
forward 2
forward 5
up 2
up 1
forward 2
down 4
down 5
forward 6
forward 1
down 7
down 6
down 2
forward 5
forward 6
up 8
up 6
forward 2
forward 6
down 9
down 4
forward 1
down 3
down 1
up 4
down 8
forward 3
down 5
up 4
down 5
up 3
down 1
down 1
down 9
forward 8
up 8
down 2
forward 9
down 8
down 4
up 4
up 9
up 4
forward 8
forward 9
forward 3
forward 2
down 2
forward 5
down 6
down 2
down 9
forward 3
up 4
forward 8
up 9
forward 2
forward 1
down 3
up 1
up 7
down 3
up 2
down 2
up 2
forward 4
down 7
forward 2
forward 4
forward 3
down 6
forward 4
down 3
forward 2
down 1
up 4
down 8
up 3
forward 4
up 6
forward 5
forward 3
forward 1
up 2
forward 4
forward 7
down 3
forward 9
up 7
down 9
forward 2
forward 4
down 9
down 1
up 3
forward 2
forward 7
down 7
forward 9
forward 5
forward 7
forward 2
forward 3
forward 4
forward 5
forward 6
down 1
forward 8
down 4
down 6
up 8
up 5
forward 5
down 9
down 2
down 2
forward 7
forward 8
up 8
down 4
up 5
forward 1
down 5
forward 3
up 3
down 6
forward 5
up 3
up 5
forward 4
forward 2
up 6
down 9
forward 7`

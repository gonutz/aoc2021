package main

import "fmt"

func main() {
	minx, maxx := 248, 285
	miny, maxy := -85, -56
	in := func(x, y int) bool {
		return minx <= x && x <= maxx && miny <= y && y <= maxy
	}
	over := func(x, y int) bool {
		return x > maxx || y < miny
	}

	totalmaxy := 0
	count := 0
	for yvel := -1000; yvel < 1000; yvel++ {
		for xvel := 0; xvel < 1000; xvel++ {
			vx := xvel
			vy := yvel
			x := 0
			y := 0
			thismaxy := 0

			for {
				x += vx
				y += vy
				if vx > 0 {
					vx--
				}
				if vx < 0 {
					vx++
				}
				vy--
				if y > thismaxy {
					thismaxy = y
				}
				if in(x, y) {
					count++
					if thismaxy > totalmaxy {
						totalmaxy = thismaxy
					}
					break
				}
				if over(x, y) {
					break
				}
			}
		}
	}

	fmt.Println("part 1)", totalmaxy)
	fmt.Println("part 2)", count)
}

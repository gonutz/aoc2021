package main

import "fmt"

func main() {
	flashes := 0

	for i := 0; i < 1000; i++ {
		flashesThisStep := 0

		for y := 1; y < 11; y++ {
			for x := 1; x < 11; x++ {
				input[x][y]++
			}
		}

		var flashed [12][12]bool
		changed := true
		for changed {
			changed = false
			for y := 1; y < 11; y++ {
				for x := 1; x < 11; x++ {
					if input[x][y] > 9 && !flashed[x][y] {
						flashed[x][y] = true
						flashesThisStep++
						changed = true
						for x2 := x - 1; x2 <= x+1; x2++ {
							for y2 := y - 1; y2 <= y+1; y2++ {
								if x != x2 || y != y2 {
									input[x2][y2]++
								}
							}
						}
					}
				}
			}
		}

		for y := 1; y < 11; y++ {
			for x := 1; x < 11; x++ {
				if input[x][y] > 9 {
					input[x][y] = 0
				}
			}
		}

		if flashesThisStep == 100 {
			defer func() {
				fmt.Println("part 2)", i+1)
			}()
			break
		}

		if i < 100 {
			flashes += flashesThisStep
		}
	}

	fmt.Println("part 1)", flashes)
}

const o = -999999999

var input = [][]int{
	{o, o, o, o, o, o, o, o, o, o, o, o},
	{o, 6, 6, 1, 7, 1, 1, 3, 5, 8, 4, o},
	{o, 6, 5, 4, 4, 2, 1, 8, 6, 3, 8, o},
	{o, 5, 4, 5, 7, 3, 3, 1, 4, 8, 8, o},
	{o, 1, 1, 3, 5, 6, 7, 5, 5, 8, 7, o},
	{o, 1, 2, 2, 1, 3, 5, 3, 2, 1, 6, o},
	{o, 1, 8, 1, 1, 1, 2, 4, 3, 7, 8, o},
	{o, 1, 3, 8, 7, 8, 6, 4, 3, 6, 8, o},
	{o, 4, 4, 2, 7, 6, 3, 7, 2, 6, 2, o},
	{o, 6, 7, 7, 8, 6, 4, 5, 4, 8, 6, o},
	{o, 3, 6, 8, 2, 1, 4, 6, 7, 4, 5, o},
	{o, o, o, o, o, o, o, o, o, o, o, o},
}

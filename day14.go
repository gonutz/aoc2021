package main

import "fmt"

func main() {
	fmt.Println("part 1)", doSteps(10))
	fmt.Println("part 2)", doSteps(40))
}

func doSteps(stepCount int) int {
	pairCounts := make(map[string]int)
	for i := 1; i < len(input); i++ {
		pairCounts[input[i-1:i+1]]++
	}
	letterCounts := make(map[string]int)
	for _, r := range input {
		letterCounts[string(r)]++
	}

	for i := 0; i < stepCount; i++ {
		newPairCounts := make(map[string]int)
		for pair, count := range pairCounts {
			insert := rule(pair)
			newPairCounts[pair[:1]+insert] += count
			newPairCounts[insert+pair[1:]] += count
			letterCounts[insert] += count
		}
		pairCounts = newPairCounts
	}

	min := letterCounts[input[:1]]
	max := min
	for _, c := range letterCounts {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	return max - min
}

func rule(s string) string {
	for _, r := range rules {
		if s == r[0] {
			return r[1]
		}
	}
	panic("no such rule: " + s)
}

const input = "SFBBNKKOHHHPFOFFSPFV"

var rules = [][2]string{
	{"HB", "C"},
	{"KO", "S"},
	{"KK", "N"},
	{"PF", "F"},
	{"VB", "F"},
	{"KC", "S"},
	{"BP", "H"},
	{"SS", "H"},
	{"BS", "B"},
	{"PB", "O"},
	{"VH", "C"},
	{"BK", "S"},
	{"BO", "F"},
	{"HN", "V"},
	{"NN", "K"},
	{"PV", "C"},
	{"NH", "P"},
	{"KP", "N"},
	{"NB", "V"},
	{"NF", "V"},
	{"PP", "O"},
	{"PN", "B"},
	{"VN", "K"},
	{"SC", "O"},
	{"NS", "N"},
	{"SV", "B"},
	{"BV", "P"},
	{"FV", "F"},
	{"OK", "H"},
	{"HF", "F"},
	{"CV", "K"},
	{"KB", "C"},
	{"OB", "B"},
	{"NO", "V"},
	{"OF", "B"},
	{"HP", "C"},
	{"BB", "F"},
	{"FB", "H"},
	{"OC", "K"},
	{"NV", "H"},
	{"OV", "S"},
	{"OP", "N"},
	{"SP", "N"},
	{"FK", "F"},
	{"VV", "B"},
	{"VK", "H"},
	{"OS", "F"},
	{"CO", "F"},
	{"CH", "V"},
	{"HV", "V"},
	{"FN", "B"},
	{"CS", "F"},
	{"PS", "F"},
	{"HS", "F"},
	{"VO", "K"},
	{"NP", "F"},
	{"FP", "B"},
	{"KF", "P"},
	{"CC", "N"},
	{"BF", "S"},
	{"VP", "F"},
	{"HO", "H"},
	{"FC", "F"},
	{"BH", "K"},
	{"NK", "S"},
	{"BN", "V"},
	{"SH", "K"},
	{"CP", "B"},
	{"VS", "K"},
	{"ON", "S"},
	{"FS", "P"},
	{"HK", "F"},
	{"PC", "O"},
	{"KN", "H"},
	{"CK", "N"},
	{"HH", "N"},
	{"CN", "S"},
	{"BC", "K"},
	{"PH", "N"},
	{"OO", "B"},
	{"FO", "O"},
	{"SK", "B"},
	{"FF", "V"},
	{"VC", "N"},
	{"SF", "N"},
	{"KH", "V"},
	{"SO", "F"},
	{"KS", "H"},
	{"SB", "K"},
	{"VF", "V"},
	{"PK", "O"},
	{"OH", "N"},
	{"HC", "F"},
	{"PO", "O"},
	{"NC", "F"},
	{"FH", "V"},
	{"KV", "V"},
	{"CB", "C"},
	{"CF", "O"},
	{"SN", "H"},
}

package days

import (
	"fmt"
	"strings"

	"github.com/colinkwiecinski/AdventOfCode/utils"
)

func DaySix() {
	data := utils.ReadLines("input/day6.txt")
	data = strings.Split(data[0], "")
	fmt.Println(day6A(data))
	fmt.Println(day6B(data))
}

func day6A(d []string) int {
	for i := 0; i < len(d) - 3; i++ {
		s := make(map[string]bool)
		s[d[i]] = true
		s[d[i + 1]] = true
		s[d[i + 2]] = true
		s[d[i + 3]] = true
		if len(s) == 4 {
			return i + 4
		}
	}
	return 0
}

func day6B(d []string) int {
	for i := 0; i < len(d) - 14; i++ {
		s := make(map[string]bool)
		s[d[i]] = true
		s[d[i + 1]] = true
		s[d[i + 2]] = true
		s[d[i + 3]] = true
		s[d[i + 4]] = true
		s[d[i + 5]] = true
		s[d[i + 6]] = true
		s[d[i + 7]] = true
		s[d[i + 8]] = true
		s[d[i + 9]] = true
		s[d[i + 10]] = true
		s[d[i + 11]] = true
		s[d[i + 12]] = true
		s[d[i + 13]] = true
		if len(s) == 14 {
			return i + 14
		}
	}
	return 0
}
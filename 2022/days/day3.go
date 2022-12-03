package days

import (
	"fmt"
	"strings"

	"github.com/colinkwiecinski/AdventOfCode/utils"
)

func DayThree() {
	data := utils.ReadLines("input/day3.txt")
	fmt.Println(day3A(data))
	fmt.Println(day3B(data))
}

func day3A(data []string) int {
	p := priority()
	sum := 0
	for _, l := range data {
		set := make(map[string]bool)
		line := strings.Split(l, "")
		front := line[0:len(line)/2]
		back := line[len(line)/2:]
		for _, c := range front {
			set[c] = true
		}
		for _, c := range back {
			if set[c] {
				sum += p[c]
				break;
			}
		}
	}
	return sum
}

func day3B(data []string) int {
	sum := 0
	p := priority()
	for i := 0; i < len(data) - 2; i+=3 {
		set1 := make(map[string]bool)
		set2 := make(map[string]bool)
		for _, c := range strings.Split(data[i], "") {
			set1[c] = true
		}
		for _, c := range strings.Split(data[i+1], "") {
			set2[c] = true
		}
		for _, c := range strings.Split(data[i+2], "") {
			if set1[c] && set2[c] {
				sum += p[c]
				break;
			}
		}
	}
	return sum
}

func priority() map[string]int {
	priority := make(map[string]int)
	j := 1
	for i := 97; i < (97 + 26); i++ {
		priority[string(rune(i))] = j
		j++
	}
	j = 27
	for i := 65; i < (65 + 26); i++ {
		priority[string(rune(i))] = j
		j++
	}
	return priority
}
package days

import (
	"fmt"
	"strconv"
	"github.com/colinkwiecinski/AdventOfCode/utils"
)

func DayOne() {
	fmt.Println("Day One")
	data := utils.ReadLines("input/day1.txt")
	fmt.Println(dayOneA(data))
	fmt.Println(dayOneB(data))
}

func dayOneA(input []string) int {
	sum := 0
	res := 0
	for _, str := range input {
		if str == "" {
			res = utils.Max(res, sum)
			sum = 0
		}
		temp, _ := strconv.Atoi(str)
		sum += temp
	}
	return res
}

func dayOneB(input []string) int {
	heap := utils.NewMaxHeap(len(input))
	sum := 0
	for _, str := range input {
		if str == "" {
			heap.Insert(sum)
			sum = 0
		}
		temp, _ := strconv.Atoi(str)
		sum += temp
	}

	TOP_ELVES := 3
	res := 0
	for i := 0; i < TOP_ELVES; i++ {
		res += heap.Remove()
	}
	return res
}

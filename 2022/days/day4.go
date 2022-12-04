package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/colinkwiecinski/AdventOfCode/utils"
)

func DayFour() {
	data := utils.ReadLines("input/day4.txt")
	fmt.Println(day4A(data))
	fmt.Println(day4B(data))
}

func day4A(data []string) int {
	sum := 0
	for _, l := range data {
		str := strings.FieldsFunc(l, func(r rune) bool { return strings.ContainsRune("-,", r) })
		nums := make([]int, 4)
		for i, n := range str {
			nums[i], _ = strconv.Atoi(n)
		}
		if (nums[0] <= nums[2] && nums[1] >= nums[3]) || (nums[2] <= nums[0] && nums[3] >= nums[1]) {
			sum++
		}
	}
	return sum
}

func day4B(data []string) int {
	sum := 0
	for _, l := range data {
		str := strings.FieldsFunc(l, func(r rune) bool { return strings.ContainsRune("-,", r) })
		nums := make([]int, 4)
		for i, n := range str {
			nums[i], _ = strconv.Atoi(n)
		}
		if (utils.Max(nums[0], nums[2]) <= utils.Min(nums[1], nums[3])){
			sum++
		}
	}
	return sum
}

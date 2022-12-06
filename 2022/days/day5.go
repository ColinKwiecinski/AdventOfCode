package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/colinkwiecinski/AdventOfCode/utils"
)


func DayFive() {
	data := utils.ReadLines("input/day5.txt")
	day5A(data)
	fmt.Println()
	day5B(data)
}

func day5A(data []string) {
	stacks := make(map[int]*utils.Stack)
	for j := 7; j >= 0; j-- {
		l := data[j]
		line := strings.Split(l, "")
			// fmt.Println("READING ", line, len(line))
			if (len(line) == 34) { //lazy hack to get rid of stack numbers
				continue
			}
			for i := 0; i < len(line) - 1; i++ {
				if line[i] != " " && line[i] != "]" && line[i] != "["{
					stacknum := ((i-1) / 4) + 1
					// fmt.Println("STACKNUM", stacknum, line[i])
					if val, ok := stacks[stacknum]; ok {
						val.Push(line[i])
					} else {
						v := utils.NewStack()
						v.Push(line[i])
						stacks[stacknum] = v
						// fmt.Println("STACK CREATION: ", stacknum, stacks[stacknum])
					}
				}
			}
	}
	for j := 10; j < len(data); j++ {
		l := data[j]
			line := strings.Split(l, " ")
			// fmt.Println("Movement Line: ", line)
			a, _ := strconv.Atoi(line[1])
			b, _ := strconv.Atoi(line[3])
			c, _ := strconv.Atoi(line[5])
			for i := 0; i < a; i++ {
				stacks[c].Push(stacks[b].Pop())
				// fmt.Println("TEMP ", stacks[b], stacks[c])
			}
	}
	for j := 1; j < len(stacks) + 1; j++ {
		v := stacks[j]
		fmt.Print(v.Pop())
	}
}

func day5B(data []string)  {
	stacks := make(map[int]*utils.Stack)
	for j := 7; j >= 0; j-- {
		l := data[j]
		line := strings.Split(l, "")
			// fmt.Println("READING ", line, len(line))
			if (len(line) == 34) { //lazy hack to get rid of stack numbers
				continue
			}
			for i := 0; i < len(line) - 1; i++ {
				if line[i] != " " && line[i] != "]" && line[i] != "["{
					stacknum := ((i-1) / 4) + 1
					// fmt.Println("STACKNUM", stacknum, line[i])
					if val, ok := stacks[stacknum]; ok {
						val.Push(line[i])
					} else {
						v := utils.NewStack()
						v.Push(line[i])
						stacks[stacknum] = v
						// fmt.Println("STACK CREATION: ", stacknum, stacks[stacknum])
					}
				}
			}
	}
	for j := 10; j < len(data); j++ {
		l := data[j]
			line := strings.Split(l, " ")
			// fmt.Println("Movement Line: ", line)
			a, _ := strconv.Atoi(line[1])
			b, _ := strconv.Atoi(line[3])
			c, _ := strconv.Atoi(line[5])
			anotherstack := utils.NewStack()
			for i := 0; i < a; i++ {
				anotherstack.Push(stacks[b].Pop())

				// fmt.Println("TEMP ", stacks[b], stacks[c])
			}
			for i := 0; i < a; i++ {
				stacks[c].Push(anotherstack.Pop())
			}
	}
	for j := 1; j < len(stacks) + 1; j++ {
		v := stacks[j]
		fmt.Print(v.Pop())
	}
}
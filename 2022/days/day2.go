package days

import (
	"fmt"

	"github.com/colinkwiecinski/AdventOfCode/utils"
)

func DayTwo() {
	fmt.Println("Day Two")
	data := utils.ReadWords("input/day2.txt")
	fmt.Println(day2A(data))
	fmt.Println(day2B(data))
}

// goofy structure because i misread and got me and them backwards at first
// also im sure there is a shorter way to do it but i want to play factorio now
func day2A(data []string) int {
	sum := 0
	for i := 0; i < len(data); i+=2 {
		them := data[i]
		me := data[i + 1]
		if me == Rock {
			sum += RockValue
			if them == RockA {
				sum += Tie
			} else if them == PaperB {
				sum += Lose
			} else {
				sum += Win
			}
		} else if me == Paper {
			sum += PaperValue
			if them == RockA {
				sum += Win
			} else if them == PaperB {
				sum += Tie
			} else {
				sum += Lose
			}
		} else { // Scissors
			sum += ScissorValue
			if them == RockA {
				sum += Lose
			} else if them == PaperB {
				sum += Win
			} else {
				sum += Tie
			}
		}
	}
	return sum
}

func day2B(data []string) int {
	sum := 0
	for i := 0; i < len(data); i+=2 {
		them := data[i]
		me := data[i + 1]
		if them == RockA {
			if me == X { // lose
				sum += Lose + ScissorValue
			} else if me == Y { // tie
				sum += Tie + RockValue
			} else { // win
				sum += Win + PaperValue
			}
		} else if them == PaperB {
			if me == X {
				sum += Lose + RockValue
			} else if me == Y {
				sum += Tie + PaperValue
			} else {
				sum += Win + ScissorValue
			}
		} else if them == ScisssorsC {
			if me == X {
				sum += Lose + PaperValue
			} else if me == Y {
				sum += Tie + ScissorValue
			} else {
				sum += Win + RockValue
			}
		}
	}
	return sum
}

// 6 for win
// 3 for tie
// 0 for fat L
const (
	Rock = "X" // 1
	Paper = "Y" // 2
	Scissors = "Z" // 3
	RockA = "A"
	PaperB = "B"
	ScisssorsC = "C"
	Win = 6
	Tie = 3
	Lose = 0
	X = "X"
	Y = "Y"
	Z = "Z"
	RockValue = 1
	PaperValue = 2
	ScissorValue = 3
)

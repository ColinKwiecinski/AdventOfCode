package utils

import (
	"bufio"
	"log"
	"os"
)

// Split on every line
func ReadLines(input string) []string {
	scanner := GetScanner(input)
	res := make([]string, 0)
	for scanner.Scan(){
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func GetScanner(input string) *bufio.Scanner {
	f, err := os.Open(input)
	if err != nil {
        log.Fatal(err)
    }
	return bufio.NewScanner(f)
}

func ReadWords(input string) []string {
	scanner := GetScanner(input)
	scanner.Split(bufio.ScanWords)
	res := make([]string, 0)
	for scanner.Scan(){
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}
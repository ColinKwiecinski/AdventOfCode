package utils

import (
	"bufio"
	"log"
	"os"
)

// Split on every line
func ReadLines(input string) []string {
	f, err := os.Open(input)
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	scanner := bufio.NewScanner(f)
	res := make([]string, 0)
	for scanner.Scan(){
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}
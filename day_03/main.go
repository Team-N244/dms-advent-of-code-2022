package main

import (
	"bufio"
	"fmt"
	"os"
)

const FILENAME = "test_input.txt"

// getCharPriority will return the Priority of the "item" identified by ch.
func getCharPriority(ch byte) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch) - int('a') + 1
	} else if ch >= 'A' && ch <= 'Z' {
		return int(ch) - int('A') + 27
	} else {
		panic("Invalid character!")
	}
}

func main() {
	fp, err := os.Open(FILENAME)
	if err != nil {
		panic("Could not read file!")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	// commonItem := '!'

	for scanner.Scan() {
		input := scanner.Text()
		firstHalf := input[:len(input)/2]
		secondHalf := input[len(input)/2:]
		fmt.Printf("First : %s\nSecond: %s\n", firstHalf, secondHalf)
	}
}

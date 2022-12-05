package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILENAME = "input.txt"

// getCharPriority will return the Priority of the "item" identified by ch.
func getCharPriority(ch int) int {
	if ch >= 'a' && ch <= 'z' {
		return ch - int('a') + 1
	} else if ch >= 'A' && ch <= 'Z' {
		return ch - int('A') + 27
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

	var commonItem int = '!'
	totalPriority := 0

	for scanner.Scan() {
		input := scanner.Text()
		firstHalf := input[:len(input)/2]
		secondHalf := input[len(input)/2:]
		fmt.Printf("First : %s\nSecond: %s\n", firstHalf, secondHalf)

		for _, ch := range firstHalf {
			if strings.Contains(secondHalf, string(ch)) {
				commonItem = int(ch)
				break
			}
		}
		totalPriority += getCharPriority(commonItem)
	}

	fmt.Printf("The total priority for the items is: %d\n", totalPriority)
}

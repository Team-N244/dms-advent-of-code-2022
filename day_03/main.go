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
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i += 3 {
		for _, ch := range lines[i] {
			if strings.Contains(lines[i+1], string(ch)) && strings.Contains(lines[i+2],
				string(ch)) {
				fmt.Printf("The common item was: %s\n", string(ch))
				commonItem = int(ch)
				break
			}
		}
		charPriority := getCharPriority(commonItem)
		fmt.Printf("Adding %s's priority of %d to total priority.\n", string(commonItem),
			charPriority)
		totalPriority += getCharPriority(commonItem)
	}

	fmt.Printf("The total priority for the items is: %d\n", totalPriority)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic("Could not read file!")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	mostCalsElf := 0
	currentElfTotal := 0

	for scanner.Scan() {
		input := scanner.Text()
		if input != "" {
			cals, err := strconv.Atoi(input)
			if err != nil {
				panic("Could not convert string to integer!")
			}
			fmt.Printf("Adding %d to current total.\n", cals)
			currentElfTotal += cals
			fmt.Printf("  Total is now: %d\n", currentElfTotal)
		} else {
			if currentElfTotal > mostCalsElf {
				fmt.Printf("New High Score: %d beats %d\n", currentElfTotal, mostCalsElf)
				mostCalsElf = currentElfTotal
			}
			currentElfTotal = 0
		}
	}

	fmt.Printf("\n\nThe Highest Total is: %d\n", mostCalsElf)

	if err := scanner.Err(); err != nil {
		panic("Error when reading file!")
	}
}

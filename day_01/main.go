package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic("Could not read file!")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var elves []int
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
			elves = append(elves, currentElfTotal)
			currentElfTotal = 0
		}
	}

	sort.Sort(sort.IntSlice(elves))

	highestElf := elves[len(elves)-1]

	fmt.Printf("\n\nThe Highest Total is: %d\n", highestElf)

	topThree := elves[len(elves)-3:]
	topTotal := 0
	for _, element := range topThree {
		topTotal += element
	}

	fmt.Printf("The Top Three are: %d\n", topThree)
	fmt.Printf("The Top Three's Total is: %d\n", topTotal)

	if err := scanner.Err(); err != nil {
		panic("Error when reading file!")
	}
}

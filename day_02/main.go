package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic("Could not read file!")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var moves = map[string]int{
		"A X": 4, // 1pt for playing Rock, 3pt for draw
		"A Y": 8, // 2pt for playing Paper, 6pt for win
		"A Z": 3, // 3pt for playing Scissors, 0pt for loss

		"B X": 1, // 1pt for playing Rock, 1pt for loss
		"B Y": 5, // 2pt for playing Paper, 3pt for draw
		"B Z": 9, // 3pt for playing Scissors, 6pt for win

		"C X": 7, // 1pt for playing Rock, 6pt for win
		"C Y": 2, // 2pt for playing Paper, 0pt for loss
		"C Z": 6, // 3pt for playing Scissors, 3pt for draw
	}

	pointTotal := 0

	for scanner.Scan() {
		input := scanner.Text()
		pointTotal += moves[input]
	}

	if err := scanner.Err(); err != nil {
		panic("Error when reading file!")
	}

	fmt.Printf("The total score is: %d\n", pointTotal)
}

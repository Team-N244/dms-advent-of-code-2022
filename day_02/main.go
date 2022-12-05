package main

import (
	"bufio"
	"fmt"
	"os"
)

const FILENAME = "input.txt"

func main() {
	fp, err := os.Open(FILENAME)
	if err != nil {
		panic("Could not read file!")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var moves = map[string]int{
		"A X": 3, // Opponent playing Rock, Need to lose (0pt), play Scissors (3pt)
		"A Y": 4, // Opponent playing Rock, Need to draw (3pt), play Rock (1pt)
		"A Z": 8, // Opponent playing Rock, Need to win  (6pt), play Paper (2pt)

		"B X": 1, // Opponent playing Paper, Need to lose (0pt), play Rock (1pt)
		"B Y": 5, // Opponent playing Paper, Need to draw (3pt), play Paper (2pt)
		"B Z": 9, // Opponent playing Paper, Need to win  (6pt), play Scissors (3pt)

		"C X": 2, // Opponent playing Scissors, Need to lose (0pt), play Paper (2pt)
		"C Y": 6, // Opponent playing Scissors, Need to draw (3pt), play Scissors (3pt)
		"C Z": 7, // Opponent playing Scissors, Need to win  (6pt), play Rock (1pt)
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

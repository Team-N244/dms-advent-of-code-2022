package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "test_input.txt"

type Sector struct {
	start int
	end   int
}

func readInput(filename string) []string {
	fp, err := os.Open(filename)
	if err != nil {
		panic("Could not open file!")
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic("Could not read file!")
	}
	return lines
}

func createSector(sectorDef string) Sector {
	start, err := strconv.Atoi(strings.Split(sectorDef, "-")[0])
	if err != nil {
		panic("Could not convert starting sector!")
	}
	end, err := strconv.Atoi(strings.Split(sectorDef, "-")[1])
	if err != nil {
		panic("Could not convert ending sector!")
	}
	return Sector{start: start, end: end}
}

func (s Sector) checkIntersection(other Sector) bool {
	return s.start >= other.start && s.end <= other.end
}

func main() {
	lines := readInput(FILENAME)

	numOverlaps := 0
	for _, line := range lines {
		splitLine := strings.Split(line, ",")
		firstSector := createSector(splitLine[0])
		secondSector := createSector(splitLine[1])
		fmt.Printf("First Sector: Start is %d, End is %d\n", firstSector.start, firstSector.end)
		fmt.Printf("Second Sector: Start is %d, End is %d\n", secondSector.start, secondSector.end)
		firstSecondOverlap := firstSector.checkIntersection(secondSector)
		secondFirstOverlap := secondSector.checkIntersection(firstSector)
		fmt.Printf("Does the first sector overlap the second? %v\n", firstSecondOverlap)
		fmt.Printf("Does the second sector overlap the first? %v\n", secondFirstOverlap)
		if firstSecondOverlap || secondFirstOverlap {
			numOverlaps++
		}
	}

	fmt.Printf("\nThere were %v overlaps in total!\n", numOverlaps)
}

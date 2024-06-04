package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var matchScores = map[string]int{
		"A X": 4,
		"B X": 1,
		"C X": 7,
		"A Y": 8,
		"B Y": 5,
		"C Y": 2,
		"A Z": 3,
		"B Z": 9,
		"C Z": 6,
	}

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		total += matchScores[scanner.Text()]
	}

	fmt.Println(total)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

//Each half of string is one of the compartments

func main() {
	var letterPriority = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}

	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		first := scanner.Text()[:len(scanner.Text())/2]
		second := scanner.Text()[len(scanner.Text())/2:]

		for _, c := range first {
			if strings.Contains(second, string(c)) {
				fmt.Printf("%c found in second compartment\n", c)
				if unicode.IsUpper(c) {
					total += 26 + letterPriority[strings.ToLower(string(c))]
					break
				}
				total += letterPriority[string(c)]
				break
			}
		}
	}

	fmt.Println(total)
}

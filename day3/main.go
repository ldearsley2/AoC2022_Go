package main

import (
	"bufio"
	"log"
	"os"
)

//Each half of string is one of the compartments

func main() {
	//var letterPriority = map[string]int{
	//	"a": 1,
	//	"b": 2,
	//	"c": 3,
	//	"d": 4,
	//	"e": 5,
	//	"f": 6,
	//	"g": 7,
	//	"h": 8,
	//	"i": 9,
	//	"j": 10,
	//	"k": 11,
	//	"l": 12,
	//	"m": 13,
	//	"n": 14,
	//	"o": 15,
	//	"p": 16,
	//	"q": 17,
	//	"r": 18,
	//	"s": 19,
	//	"t": 20,
	//	"w": 21,
	//	"x": 22,
	//	"y": 23,
	//	"z": 24,
	//}

	file, err := os.Open("day3/test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		first := scanner.Text()[:len(scanner.Text())/2]
		second := scanner.Text()[len(scanner.Text())/2:]

		for i := range len(first) {
		}
	}

}

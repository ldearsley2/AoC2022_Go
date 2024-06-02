package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tempCalories, mostCalories = 0, 0

	for scanner.Scan() {
		var calories, err = strconv.Atoi(scanner.Text())

		if err != nil {
			if tempCalories > mostCalories {
				mostCalories = tempCalories
			}
			tempCalories = 0
		}

		tempCalories += calories
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(mostCalories)
}

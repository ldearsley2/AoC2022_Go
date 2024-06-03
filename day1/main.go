package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	fmt.Println(mostCalories)
	fmt.Println(part2("day1/input.txt"))
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var tempCalories int = 0
	caloriesList := make([]int, 0)

	for scanner.Scan() {
		var calories, err = strconv.Atoi(scanner.Text())
		if err != nil {
			caloriesList = append(caloriesList, tempCalories)
			tempCalories = 0
		}
		tempCalories += calories
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total int = 0
	sort.Slice(caloriesList, func(i, j int) bool {
		return caloriesList[i] < caloriesList[j]
	})

	for i := range 3 {
		total += caloriesList[len(caloriesList)-1-i]
	}

	return total
}

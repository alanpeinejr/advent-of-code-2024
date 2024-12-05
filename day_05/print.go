package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//part 1
	fmt.Println(checkForValidUpdates(parseInput(readInput())))

	//part 2
	fmt.Println(fixInvalidUpdates(parseInput(readInput())))

}

func fixInvalidUpdates(updates [][]int, rules map[int]*Page) int {
	sum := 0
	for _, update := range updates {
		if !checkForValidUpdate(update, rules) {
			fixInvalidUpdate(update, rules)
			sum += update[len(update)/2]
		}
	}
	return sum
}

func fixInvalidUpdate(update []int, rules map[int]*Page) {
	slices.SortStableFunc(update, func(i, j int) int {
		if _, exists := rules[i].Precedes[rules[j]]; exists {
			return -1
		}
		if _, exists := rules[j].Precedes[rules[i]]; exists {
			return 1
		}
		return 0
	})
}

func checkForValidUpdates(updates [][]int, rules map[int]*Page) int {
	//score is middle number of each valid update
	sum := 0
	for _, update := range updates {
		if checkForValidUpdate(update, rules) {
			sum += update[len(update)/2]
		}
	}
	return sum

}

func checkForValidUpdate(update []int, rules map[int]*Page) bool {

	for i, number := range update {
		//if number preceeds any numbers before it return false
		for j := 0; j < i; j++ {
			if _, exists := rules[number].Precedes[rules[update[j]]]; exists {
				return false
			}
		}
	}
	return true
}

func addPage(number int) *Page {
	return &Page{number, make(map[*Page]bool), make(map[*Page]bool)}
}

func parseInput(input string) ([][]int, map[int]*Page) {
	rulesUpdates := strings.Split(input, "\n\n")
	rulesStrings := strings.Split(rulesUpdates[0], "\n")
	rules := make(map[int]*Page)
	for _, ruleString := range rulesStrings {
		precedesProceeds := strings.Split(ruleString, "|")
		precedes, proceeds := stringToInt(precedesProceeds[0]), stringToInt(precedesProceeds[1])
		//if either doesn't exist, add them to our map...book
		if _, exists := rules[precedes]; !exists {
			rules[precedes] = addPage(precedes)
		}
		if _, exists := rules[proceeds]; !exists {
			rules[proceeds] = addPage(proceeds)
		}
		//the second number comes after the first number
		rules[proceeds].Proceeds[rules[precedes]] = true
		//the first number comes before the second number
		rules[precedes].Precedes[rules[proceeds]] = true
	}

	updatesStrings := strings.Split(rulesUpdates[1], "\n")
	updates := make([][]int, len(updatesStrings))
	for i, updateString := range updatesStrings {
		update := make([]int, 0)
		for _, number := range strings.Split(updateString, ",") {
			update = append(update, stringToInt(number))

		}
		updates[i] = update
	}

	return updates, rules
}

type (
	Page struct {
		Number   int
		Precedes map[*Page]bool
		Proceeds map[*Page]bool
	}
)

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}

func readInput() string {
	var filename string
	if len(os.Args) < 2 {
		fmt.Println("Assuming local file input.txt")
		filename = "./input.txt"
	} else {
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		panic(err)
	}

	//return and account for windows
	return strings.ReplaceAll(string(data), "\r\n", "\n")
}

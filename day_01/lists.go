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
	fmt.Println(sum(findDifferences(parseInput(readInput()))))
	//part 2
	fmt.Println(sum(findHitsAndScore(parseInput(readInput()))))

}

func findHitsAndScore(left []int, right []int) []int {
	hits := make([]int, len(left))
	for i, _ := range left {
		//score = number * its occurences in the other list
		hits[i] = findOccurences(left[i], right) * left[i]
	}
	return hits
}

func findOccurences(value int, slice []int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func sum(differences []int) int {
	sum := 0
	for _, v := range differences {
		sum += v
	}
	return sum
}

func findDifferences(left []int, right []int) []int {
	differences := make([]int, len(left))
	slices.Sort(left)
	slices.Sort(right)

	for i, _ := range left {
		differences[i] = absoluteValue(left[i] - right[i])
	}
	return differences

}

func parseInput(input string) ([]int, []int) {
	rows := strings.Split(input, "\n")
	left, right := make([]int, len(rows)), make([]int, len(rows))
	for i, value := range rows {
		leftRight := strings.Split(value, "   ")
		left[i] = stringToInt(leftRight[0])
		right[i] = stringToInt(leftRight[1])

	}
	return left, right
}

func absoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return int(value)
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

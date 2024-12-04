package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	//part 1
	fmt.Println(search(parseInput(readInput()), "XMAS"))

	//part 2
	fmt.Println(crossMas(parseInput(readInput())))

}

func crossMas(lines []string) int {
	count := 0
	for i, line := range lines {
		for j, char := range line {
			if char == 'A' {
				if i == 0 || j == 0 || i == len(lines)-1 || j == len(line)-1 {
					continue
				}
				sublines := make([]string, 0)
				sublines = append(sublines, lines[i-1][j-1:j+2], lines[i][j-1:j+2], lines[i+1][j-1:j+2])
				diagonals := getDiagonals(sublines)
				//was going to use the search Function but it constructs allStrings, not just diagonals
				if (findWord("MAS", diagonals) + findWord("SAM", diagonals)) == 2 {
					count++
				}
			}
		}
	}
	return count
}

func search(lines []string, word string) int {
	count := 0
	allStrings := getStrings(lines)
	reversedWord := reverseString(word)
	count += findWord(word, allStrings)
	count += findWord(reversedWord, allStrings)
	return count
}

func findWord(word string, allStrings []string) int {
	count := 0
	for _, str := range allStrings {
		count += strings.Count(str, word)
	}
	return count

}

func getDiagonals(lines []string) []string {
	linesInverse := make([]string, len(lines))
	for i, line := range lines {
		linesInverse[i] = reverseString(line)
	}
	diagonals := make([]string, 0)
	diagonals = append(diagonals, constructHalfDiagonals(lines)...)
	diagonals = append(diagonals, constructHalfDiagonals(linesInverse)...)

	//reverse the slices to get other halves
	slices.Reverse(lines)
	slices.Reverse(linesInverse)

	//get other halves
	diagonals = append(diagonals, constructHalfDiagonals(lines)...)
	//but we have to cut the longest because it repeats
	diagonals = diagonals[:len(diagonals)-1]
	diagonals = append(diagonals, constructHalfDiagonals(linesInverse)...)
	diagonals = diagonals[:len(diagonals)-1]
	slices.Reverse(lines) //just for the sake of it
	return diagonals
}

func getStrings(lines []string) []string {
	rows := make([]string, len(lines))
	for i, _ := range rows {
		for _, line := range lines {
			rows[i] += string(line[i])
		}
	}

	diagonals := getDiagonals(lines)

	//combine the lists
	allStrings := make([]string, 0)
	allStrings = append(allStrings, lines...)
	allStrings = append(allStrings, rows...)
	allStrings = append(allStrings, diagonals...)

	return allStrings
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func constructHalfDiagonals(axis []string) []string {
	diagonalDown := make([]string, len(axis))
	for i, axii := range axis {
		diagonalDown[i] = string(axii[0])
		for j := i - 1; j >= 0; j-- {
			diagonalDown[i] += string(axis[j][i-j])
		}
	}
	return diagonalDown
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

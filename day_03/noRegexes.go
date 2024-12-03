package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//part 1
	fmt.Println(sum(noRegexes(parseInput(readInput()), false)))
	//part 2
	fmt.Println(sum(noRegexes(parseInput(readInput()), true)))

}

func sum(values []Multiplicatives) int {
	total := 0
	for _, value := range values {
		total += value.A * value.B
	}
	return total
}

func noRegexes(input string, part2 bool) []Multiplicatives {
	multiplicatives := make([]Multiplicatives, 0)
	pending := new(Multiplicatives)
	index := 0
	stage := FUNCTION
	for index < len(input) {
		switch stage {
		case DONT:
			index = findDo(index, input)
			stage = FUNCTION
			break
		case FUNCTION:
			oldIndex := index
			dont := false
			index = findFunction(index, input)
			//at first "don't, change index to that, update stage, break
			if part2 {
				dont, index = didWeDont(oldIndex, index, input)
				if dont {
					stage = DONT
					break
				}
			}
			if index < len(input) {
				stage = FIRSTPAREN
				//kinda dumb but eh
				pending.FunctionIndex = index - 3
			}
			break
		case FIRSTPAREN:
			if isOpeningParen(index, input) {
				stage = NUMBERS
			} else {
				stage = FUNCTION
			}
			index++
			break
		case NUMBERS:
			var first, second int
			first, second, index = findNumbers(index, input)
			if first != -1 && second != -1 {
				pending.A = first
				pending.B = second
				stage = SECONDPAREN
			} else {
				stage = FUNCTION
			}
			break
		case SECONDPAREN:
			if isClosingParen(index, input) {
				multiplicatives = append(multiplicatives, Multiplicatives{pending.FunctionIndex, pending.A, pending.B})
			}
			stage = FUNCTION
			index++
			break

		}
	}
	return multiplicatives
}

func didWeDont(oldIndex int, newIndex int, input string) (bool, int) {
	hit := strings.Index(input[oldIndex:newIndex], "don't()")
	if hit == -1 {
		return false, newIndex
	}
	//oldIndex + index within slice of start of d, plux length of dont
	return true, oldIndex + hit + 7
}

func findDo(index int, input string) int {
	//locate position of the string "do"
	hit := strings.Index(input[index:], "do()")

	if hit == -1 {
		return len(input)
	} else {
		return index + hit + 4
	}

}

func findFunction(index int, input string) int {
	//locate position of the string "mul"
	hit := strings.Index(input[index:], "mul")
	if hit == -1 {
		return len(input)
	} else {
		//where we were, the hit, and length of mul
		return index + hit + 3
	}
}
func isOpeningParen(index int, input string) bool {
	return input[index] == '('
}

func findNumbers(index int, input string) (int, int, int) {
	firstNumber := findNumberString(input[index:])
	if len(firstNumber) <= 3 && len(firstNumber) > 0 {
		//we dont add 1 becuase numer is at least 1?
		if input[index+len(firstNumber)] == ',' {
			secondNumber := findNumberString(input[index+len(firstNumber)+1:])
			if len(secondNumber) <= 3 && len(secondNumber) > 0 {
				return stringToInt(firstNumber), stringToInt(secondNumber), index + len(firstNumber) + 1 + len(secondNumber)
			}
		}
	}
	return -1, -1, index + 1
}
func findNumberString(input string) string {
	//iterate until first non digit
	number := ""
	for _, char := range input {
		if unicode.IsDigit(char) {
			number += string(char)
		} else {
			break
		}
	}
	return number
}

func isClosingParen(index int, input string) bool {
	return input[index] == ')'
}

type Multiplicatives struct {
	FunctionIndex int
	A             int
	B             int
}

const (
	DONT        = "don't"
	FUNCTION    = "mul"
	FIRSTPAREN  = "("
	NUMBERS     = ","
	SECONDPAREN = ")"
)

func parseInput(input string) string {
	//cut the lines
	runon := strings.ReplaceAll(input, "\n", "")
	return runon
}

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

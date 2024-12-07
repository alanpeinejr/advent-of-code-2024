package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part 1
	fmt.Println(solveEquations(parseInput(readInput()), false))

	//part 2
	fmt.Println(solveEquations(parseInput(readInput()), true))

}

func encodeBase3(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		result = string("012"[num%3]) + result
		num /= 3
	}
	return result
}

func solveEquations(equations []*Equation, part2 bool) int {
	total := 0
	for _, equation := range equations {
		if part2 {
			if solveEquation3(equation) {
				total += equation.Result
			}
		} else {
			if solveEquation(equation) {
				total += equation.Result
			}
		}
	}
	return total
}

func solveEquation3(equation *Equation) bool {
	operatorLength := len(equation.Operands) - 1
	//combinations = 3^operatorLength
	for i := 0; i < powInt(3, operatorLength); i++ {
		s := fmt.Sprintf("%0*s", operatorLength, encodeBase3(i))
		operators := make([]rune, operatorLength)
		for index, value := range s {
			switch value {
			case '0':
				operators[index] = '+'
				break
			case '1':
				operators[index] = '*'
				break
			case '2':
				operators[index] = '|'
			default:
				panic("whoops")
			}
		}
		if equation.isValidWith(operators) {
			equation.Operators = operators
			return true
		}
	}
	return false
}

func solveEquation(equation *Equation) bool {
	operatorLength := len(equation.Operands) - 1
	//combinations = 2^operatorLength
	for i := 0; i < powInt(2, operatorLength); i++ {
		//thers a better way to do this, but string of bits
		s := fmt.Sprintf("%0*b", operatorLength, i)
		operators := make([]rune, operatorLength)
		for index, value := range s {
			switch value {
			case '0':
				operators[index] = '+'
				break
			case '1':
				operators[index] = '*'
				break
			default:
				panic("whoops")
			}
		}
		if equation.isValidWith(operators) {
			equation.Operators = operators
			return true
		}
	}
	return false
}

type (
	Equation struct {
		Result    int
		Operators []rune
		Operands  []int
	}
)

func parseInput(input string) []*Equation {
	lines := strings.Split(input, "\n")
	equations := make([]*Equation, len(lines))
	for i, line := range lines {
		equation := new(Equation)
		resultOperands := strings.Split(line, ": ")
		equation.Result = stringToInt(resultOperands[0])
		operators := strings.Split(resultOperands[1], " ")
		equation.Operands = make([]int, len(operators))
		for j, operator := range operators {
			equation.Operands[j] = stringToInt(operator)
		}
		equation.Operators = make([]rune, 0)
		equations[i] = equation
	}

	return equations
}

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}

func intToString(this int) string {
	return strconv.Itoa(this)
}
func (this Equation) isValidWith(operators []rune) bool {
	total := this.Operands[0]
	for i, operator := range operators {
		switch operator {
		case '+':
			total += this.Operands[i+1]
		case '*':
			total *= this.Operands[i+1]
		case '|':
			total = stringToInt(intToString(total) + intToString(this.Operands[i+1]))
		}
	}
	return total == this.Result
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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

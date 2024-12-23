package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//test := parseInput(readInput())
	//fmt.Println(removeIndex(test[0], 0))
	//fmt.Println(test[0])

	//part 1
	fmt.Println(countSafeReports(parseInput(readInput())))
	//part 2
	fmt.Println(countSafishReports(parseInput(readInput())))

}

func countSafeReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func countSafishReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if isSafish(report) {
			safeReports++
		}
	}
	return safeReports
}

func isSafish(report []int) bool {
	if isSafe(report) {
		return true
	}
	for i, _ := range report {
		if isSafe(removeIndex(report, i)) {
			return true
		}
	}

	return false

}

func isSafe(report []int) bool {
	return (allIncreasing(report) || allDecreasing(report)) && withinTolerance(report)
}
func allIncreasing(report []int) bool {
	for i, value := range report {
		if i > 0 && value < report[i-1] {
			return false
		}
	}
	return true
}

func allDecreasing(report []int) bool {
	for i, value := range report {
		if i > 0 && value > report[i-1] {
			return false
		}
	}
	return true

}

func withinTolerance(report []int) bool {
	for i, value := range report {
		if i > 0 {
			difference := absoluteValue(value - report[i-1])
			if difference < 1 || difference > 3 {
				return false
			}
		}
	}
	return true

}

func removeIndex(s []int, index int) []int {
	// go's append can modify the original slice, what?. ?...? Why does it return a slice then?
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func parseInput(input string) [][]int {
	reportLines := strings.Split(input, "\n")
	reports := make([][]int, len(reportLines))
	for i, value := range reportLines {
		levelStrings := strings.Split(value, " ")
		levels := make([]int, len(levelStrings))
		for j, levelString := range levelStrings {
			levels[j] = stringToInt(levelString)
		}
		reports[i] = levels
	}
	return reports
}

func absoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

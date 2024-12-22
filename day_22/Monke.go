package main

import (
  "fmt"
  "os"
  "slices"
  "strconv"
  "strings"
)

func main() {
  values := parseInput(readInput())
  sum, results := generateForAllNumbers(values, 2000)
  //part1
  fmt.Println(sum)
  //part 2
  fmt.Println(findChangeOccurrences(results)[0].Data.Sum)
}

type (
  ChangeSequence struct {
    One, Two, Three, Four int
  }
  InterestingData struct {
    Occurrences  int
    HitResultMap map[int]int
    Sum          int
  }
  SortableChangeSequuence struct {
    Sequence ChangeSequence
    Data     *InterestingData
  }
)

func findChangeOccurrences(results [][]int) []SortableChangeSequuence {
  changes := map[ChangeSequence]*InterestingData{}
  changeResults := make([][]int, len(results))
  //for every change sequence of 4 across all values,
  for i := 0; i < len(results); i++ {
    changeResult := make([]int, len(results[i])-1)
    for j := 1; j < len(results[i]); j++ {
      changeResult[j-1] = results[i][j] - results[i][j-1]
      if j > 3 {
        sequence := ChangeSequence{changeResult[j-4], changeResult[j-3], changeResult[j-2], changeResult[j-1]}
        if _, exists := changes[sequence]; !exists {
          changes[sequence] = &InterestingData{1, map[int]int{i: results[i][j]}, results[i][j]}
        } else {
          
          changes[sequence].Occurrences += 1
          if _, exists := changes[sequence].HitResultMap[i]; !exists {
            //can only use the earliest hit
            changes[sequence].HitResultMap[i] = results[i][j]
            changes[sequence].Sum += results[i][j]
          }
        }
      }
    }
    changeResults[i] = changeResult
  }
  changeArray := make([]SortableChangeSequuence, len(changes))
  i := 0
  for k, v := range changes {
    changeArray[i] = SortableChangeSequuence{k, v}
    i++
  }
  slices.SortFunc(changeArray, func(a, b SortableChangeSequuence) int {
    //so the highest occurrence is first
    return b.Data.Sum - a.Data.Sum
  })
  return changeArray
}

func generateForAllNumbers(values []int, generations int) (int, [][]int) {
  sum := 0
  results := make([][]int, len(values))
  for i, value := range values {
    prices := make([]int, generations+1)
    prices[0] = value % 10
    sum += generateNumberX(value, generations, prices)
    results[i] = prices
  }
  return sum, results
}

func generateNumberX(number int, x int, prices []int) int {
  if x == 0 {
    return number
  }
  mult := ((64 * number) ^ number) % 16777216
  divide := ((mult / 32) ^ mult) % 16777216
  mult2 := ((divide * 2048) ^ divide) % 16777216
  prices[len(prices)-x] = mult2 % 10
  return generateNumberX(mult2, x-1, prices)
}

func parseInput(input string) []int {
  lines := strings.Split(input, "\n")
  values := make([]int, len(lines))
  for i, line := range lines {
    values[i] = stringToInt(line)
  }
  return values
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

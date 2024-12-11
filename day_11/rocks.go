package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(countStonesAfterBlinking(parseInput(readInput()), 25))
  
  //part 2
  fmt.Println(countStonesAfterBlinking(parseInput(readInput()), 75))
  
}

func countStonesAfterBlinking(rocks []int, x int) int {
  count := 0
  cache := make(map[Cache]int)
  for _, rock := range rocks {
    count += blinkX(rock, x, cache)
  }
  
  return count
  
}

type (
  Cache struct {
    X     int
    Value int
  }
)

func blinkX(rock int, x int, cache map[Cache]int) int {
  if x == 0 {
    return 1
  }
  //if cache hit return value
  if cachedValue, exists := cache[Cache{x, rock}]; exists {
    return cachedValue
  }
  
  var descendents int
  valueString := intToString(rock)
  switch {
  case rock == 0:
    descendents = blinkX(1, x-1, cache)
  case len(valueString)%2 == 0:
    leftString, rightString := valueString[:len(valueString)/2], valueString[len(valueString)/2:]
    leftValue, rightValue := stringToInt(leftString), stringToInt(rightString)
    descendents = blinkX(leftValue, x-1, cache) + blinkX(rightValue, x-1, cache)
  default:
    descendents = blinkX(rock*2024, x-1, cache)
  }
  //cache value and return
  cache[Cache{x, rock}] = descendents
  return descendents
}

func parseInput(input string) []int {
  rockStrings := strings.Split(input, " ")
  rocks := make([]int, len(rockStrings))
  for i, rockString := range rockStrings {
    rocks[i] = stringToInt(rockString)
  }
  return rocks
}

func stringToInt(this string) int {
  value, _ := strconv.Atoi(this)
  return value
}

func intToString(this int) string {
  return strconv.Itoa(this)
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

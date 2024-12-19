package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  //part 1
  //part 2
  fmt.Println(countPossiblePatterns(parseInput(readInput())))
  
}

func countPossiblePatterns(towels []string, patterns []string) (int, int) {
  countPossible := 0
  totalPossibleWays := 0
  cache := make(map[string]int)
  for _, pattern := range patterns {
    isPossible := isPatternPossible(towels, pattern, cache)
    if isPossible > 0 {
      countPossible++
      totalPossibleWays += isPossible
    }
  }
  return countPossible, totalPossibleWays
}

func isPatternPossible(towels []string, pattern string, cache map[string]int) int {
  if len(pattern) == 0 {
    return 1
  }
  if value, exists := cache[pattern]; exists {
    return value
  }
  possiblePatterns := 0
  for _, towel := range towels {
    //find the towels that could start the pattern
    if len(towel) <= len(pattern) && towel == pattern[:len(towel)] {
      possiblePatterns += isPatternPossible(towels, pattern[len(towel):], cache)
    }
  }
  
  cache[pattern] = possiblePatterns
  return possiblePatterns
}

func parseInput(input string) ([]string, []string) {
  towelsPatterns := strings.Split(input, "\n\n")
  towels := strings.Split(towelsPatterns[0], ", ")
  patterns := strings.Split(towelsPatterns[1], "\n")
  return towels, patterns
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

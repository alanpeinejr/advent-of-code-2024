package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  sequences := parseInput(readInput())
  //part 1
  fmt.Println(getSequencesScores(sequences, 2))
  
  //part 2
  fmt.Println(getSequencesScores(sequences, 25))
}

type (
  Cache struct {
    Sequence string
    Depth    int
  }
)

var cache = make(map[Cache]int)

func getSequencesScores(sequences []string, arrowPads int) int {
  score := 0
  for _, sequence := range sequences {
    numeric := getSequenceNumeric(sequence)
    length := getSequenceLength(sequence, arrowPads+1)
    fmt.Println(numeric, length)
    score += numeric * length
  }
  return score
}

func getSequenceNumeric(sequence string) int {
  return stringToInt(sequence[:len(sequence)-1])
}

func getSequenceLength(sequence string, depth int) int {
  if depth == 0 {
    return len(sequence)
  }
  
  cacheKey := Cache{sequence, depth}
  if value, exists := cache[cacheKey]; exists {
    return value
  }
  
  //sequences always start over A because it was the last button pressed
  current := 'A'
  length := 0
  for _, char := range sequence {
    length += getSequenceLength(ShortestPaths[current][char]+"A", depth-1)
    current = char
  }
  cache[cacheKey] = length
  return length
}

var ShortestPaths = map[rune]map[rune]string{
  'A': {'A': "", '0': "<", '1': "^<<", '2': "<^", '3': "^", '4': "^^<<", '5': "<^^", '6': "^^", '7': "^^^<<", '8': "<^^^", '9': "^^^", '<': "v<<", '^': "<", 'v': "<v", '>': "v"},
  '0': {'A': ">", '0': "", '1': "^<", '2': "^", '3': "^>", '4': "<^^", '5': "^^", '6': "^^>", '7': "^^^<", '8': "^^^", '9': "^^^>"},
  '1': {'A': "v>>", '0': ">v", '1': "", '2': ">", '3': ">>", '4': "^", '5': "^>", '6': "^>>", '7': "^^", '8': "^^>", '9': "^^>>"},
  '2': {'A': "v>", '0': "v", '1': "<", '2': "", '3': ">", '4': "<^", '5': "^", '6': "^>", '7': "<^^", '8': "^^", '9': "^^>"},
  '3': {'A': "v", '0': "<v", '1': "<<", '2': "<", '3': "", '4': "<<^", '5': "<^", '6': "^", '7': "<<^^", '8': "<^^", '9': "^^"},
  '4': {'A': ">>vv", '0': ">vv", '1': "v", '2': "v>", '3': "v>>", '4': "", '5': ">", '6': ">>", '7': "^", '8': "^>", '9': "^>>"},
  '5': {'A': "vv>", '0': "vv", '1': "<v", '2': "v", '3': "v>", '4': "<", '5': "", '6': ">", '7': "<^", '8': "^", '9': "^>"},
  '6': {'A': "vv", '0': "<vv", '1': "<<v", '2': "<v", '3': "v", '4': "<<", '5': "<", '6': "", '7': "<<^", '8': "<^", '9': "^"},
  '7': {'A': ">>vvv", '0': ">vvv", '1': "vv", '2': "vv>", '3': "vv>>", '4': "v", '5': "v>", '6': "v>>", '7': "", '8': ">", '9': ">>"},
  '8': {'A': "vvv>", '0': "vvv", '1': "<vv", '2': "vv", '3': "vv>", '4': "<v", '5': "v", '6': "v>", '7': "<", '8': "", '9': ">"},
  '9': {'A': "vvv", '0': "<vvv", '1': "<<vv", '2': "<vv", '3': "vv", '4': "<<v", '5': "<v", '6': "v", '7': "<<", '8': "<", '9': ""},
  '<': {'<': "", 'v': ">", '^': ">^", '>': ">>", 'A': ">>^"},
  'v': {'<': "<", 'v': "", '^': "^", '>': ">", 'A': "^>"},
  '^': {'<': "v<", 'v': "v", '^': "", '>': "v>", 'A': ">"},
  '>': {'<': "<<", '^': "<^", 'v': "<", '>': "", 'A': "^"},
}

func parseInput(input string) []string {
  return strings.Split(input, "\n")
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

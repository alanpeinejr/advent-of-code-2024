package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  maze, start, end, bottomRight := parseInput(readInput())
  
  //part 1
  fmt.Println(len(findPathsUnder(maze, start, end, bottomRight, 100, 2)))
  
  //part 2
  fmt.Println(len(findPathsUnder(maze, start, end, bottomRight, 100, 20)))
  
}

func findPathsUnder(maze map[Position]rune, start, end Position, bottomRight Position, requiredShortCut int, cheatDistance int) map[Cheat]bool {
  startingRacer := &Racer{start, 0}
  lengthsFromTheEnd := findLengthsFromTheEndWithoutCheats(maze, start, end)
  length := lengthsFromTheEnd[start] - requiredShortCut
  visited := map[Position]bool{start: true}
  queue := NewPriorityQueue[*Racer](func(a, b *Racer) int {
    return a.Path - b.Path
  })
  queue.Push(startingRacer)
  cheats := make(map[Cheat]bool)
  
  for !queue.isEmpty() {
    current := queue.Pop()
    if current.Path > length {
      break
    }
    
    //for all noncheaters
    for _, direction := range ALL {
      newPosition := current.Position.add(direction)
      if !visited[newPosition] {
        if maze[newPosition] == OPEN {
          newRacer := &Racer{newPosition, current.Path + 1}
          visited[newPosition] = true
          queue.Push(newRacer)
        }
      }
    }
    //for all possible cheats from here
    for _, cheat := range getAllPossibleCheatsFromHere(maze, current.Position, cheatDistance) {
      if lengthsFromTheEnd[current.Position] > lengthsFromTheEnd[cheat] {
        if (current.Path + current.Position.getDistance(cheat) + lengthsFromTheEnd[cheat]) <= length {
          cheats[Cheat{current.Position, cheat}] = true
        }
      }
    }
    
  }
  return cheats
}

func getAllPossibleCheatsFromHere(maze map[Position]rune, current Position, cheatDistance int) []Position {
  cheats := make([]Position, 0)
  for y := current.Y - cheatDistance; y < current.Y+cheatDistance+1; y++ {
    for x := current.X - cheatDistance; x < current.X+cheatDistance+1; x++ {
      // it exists
      cheatPosition := Position{x, y}
      if char, exists := maze[cheatPosition]; exists {
        //its not a wall
        if char != WALL {
          //its a valid cheat
          if current.getDistance(cheatPosition) <= cheatDistance {
            cheats = append(cheats, cheatPosition)
          }
        }
      }
    }
  }
  //panic("break")
  return cheats
}

func findLengthsFromTheEndWithoutCheats(maze map[Position]rune, start, end Position) map[Position]int {
  startingRacer := &Racer{end, 0}
  queue := NewPriorityQueue[*Racer](func(a, b *Racer) int {
    return a.Path - b.Path
  })
  queue.Push(startingRacer)
  lengthFromEnd := map[Position]int{end: 0}
  for !queue.isEmpty() {
    current := queue.Pop()
    for _, direction := range ALL {
      newPosition := current.Position.add(direction)
      if maze[newPosition] != WALL {
        //we can't go into walls
        if _, exists := lengthFromEnd[newPosition]; !exists {
          newRacer := &Racer{newPosition, current.Path + 1}
          lengthFromEnd[newPosition] = newRacer.Path
          queue.Push(newRacer)
        }
      }
    }
  }
  return lengthFromEnd
}

func (this Position) add(that Position) Position {
  return Position{this.X + that.X, this.Y + that.Y}
}

func (this Position) getDistance(that Position) int {
  return absoluteValue(this.X-that.X) + absoluteValue(this.Y-that.Y)
}

func absoluteValue(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

var ALL = []Position{RIGHT, LEFT, UP, DOWN}
var LEFT = Position{-1, 0}
var RIGHT = Position{1, 0}
var UP = Position{0, -1}
var DOWN = Position{0, 1}

const (
  WALL = '#'
  OPEN = '.'
)

type (
  Position struct {
    X, Y int
  }
  Racer struct {
    Position Position
    Path     int
  }
  Cheat struct {
    Start, End Position
  }
)

func parseInput(input string) (map[Position]rune, Position, Position, Position) {
  lines := strings.Split(input, "\n")
  maze := make(map[Position]rune)
  var start, end, bottomRight Position
  for y, line := range lines {
    for x, char := range line {
      position := Position{x, y}
      if char == 'S' {
        start = position
      }
      if char == 'E' {
        end = position
      }
      maze[position] = char
    }
  }
  bottomRight = Position{len(lines[0]) - 1, len(lines) - 1}
  return maze, start, end, bottomRight
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

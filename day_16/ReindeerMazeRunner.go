package main

import (
  "fmt"
  "math"
  "os"
  "strings"
)

func main() {
  score, uniqOnPath := findPath(parseInput(readInput()))
  //part 1
  fmt.Println(score)
  
  //part 2
  fmt.Println(uniqOnPath)
  
}

func findPath(maze map[Position]rune, start, end Position) (int, int) {
  startingReindeer := &Reindeer{start, RIGHT, []Position{start}, 0}
  queue := NewPriorityQueue[*Reindeer](func(a, b *Reindeer) int {
    return a.Points - b.Points
  })
  queue.Push(startingReindeer)
  
  //shouldn't happen given priority, but if we visit a place later for cheaper we'll revisit
  //ended up using visited to allow for multiple routes to the end
  visited := make(map[Visited]int)
  finalScore := math.MaxInt
  routes := make([]*Reindeer, 0)
  
  for !queue.isEmpty() {
    current := queue.Pop()
    if current.Points > finalScore {
      //we've seen the last of the shortest paths
      break
    }
    
    if current.Position == end {
      finalScore = current.Points
      routes = append(routes, current)
    }
    for _, direction := range current.Direction.getOptions() {
      newPosition := current.Position.add(direction)
      if maze[newPosition] == WALL {
        continue
      }
      newCost := current.Direction.getCost(direction)
      visitedKey := Visited{newPosition, direction}
      //if we've not been there for the same amount
      if visited[visitedKey] == 0 || visited[visitedKey] == current.Points+newCost {
        newReindeer := &Reindeer{newPosition, direction, addToPath(current.Path, newPosition), current.Points + newCost}
        visited[visitedKey] = newReindeer.Points
        queue.Push(newReindeer)
      }
    }
  }
  return finalScore, getUniquePointsInPath(routes)
}

func getUniquePointsInPath(reindeers []*Reindeer) int {
  uniq := map[Position]bool{}
  for _, reindeer := range reindeers {
    for _, position := range reindeer.Path {
      uniq[position] = true
    }
  }
  return len(uniq)
}

func addToPath(path []Position, position Position) []Position {
  newPath := make([]Position, len(path)+1)
  copy(newPath, path)
  newPath[len(path)] = position
  return newPath
}

func (this Position) add(that Position) Position {
  return Position{this.X + that.X, this.Y + that.Y}
}

func (this Position) getCost(that Position) int {
  if this == that {
    return 1
  }
  //its a turn
  return 1001
}

func (this Position) getOptions() []Position {
  switch this {
  case DOWN:
    return []Position{this, RIGHT, LEFT}
  case UP:
    return []Position{this, RIGHT, LEFT}
  case LEFT:
    return []Position{this, UP, DOWN}
  case RIGHT:
    return []Position{this, UP, DOWN}
  }
  fmt.Println(this)
  panic("whoops")
}

var LEFT = Position{-1, 0}
var RIGHT = Position{1, 0}
var UP = Position{0, -1}
var DOWN = Position{0, 1}

const (
  WALL     = '#'
  OPEN     = '.'
  BOX      = 'O'
  ROBBIT   = '@'
  BOXLEFT  = '['
  BOXRIGHT = ']'
)

type (
  Position struct {
    X, Y int
  }
  Reindeer struct {
    Position  Position
    Direction Position
    Path      []Position
    Points    int
  }
  Visited struct {
    Position  Position
    Direction Position
  }
)

func parseInput(input string) (map[Position]rune, Position, Position) {
  lines := strings.Split(input, "\n")
  maze := make(map[Position]rune)
  var start, end Position
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
  return maze, start, end
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

package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  bricks, grid := parseInput(readInput())
  fmt.Println(part1(bricks, grid, 1024))
  
  //part 2
  bricks, grid = parseInput(readInput())
  fmt.Println(part2(bricks, grid))
}

var gridSize = 71

func part1(bricks []Position, grid map[Position]rune, letFall int) int {
  for i := 0; i < letFall; i++ {
    placeBrick(grid, bricks[i])
  }
  return findPath(grid)
}

func part2(bricks []Position, grid map[Position]rune) Position {
  blocker := 0
  for i, brick := range bricks {
    placeBrick(grid, brick)
    if findPath(grid) == 0 {
      blocker = i
      break
    }
  }
  return bricks[blocker]
}

func findPath(grid map[Position]rune) int {
  start := Position{0, 0}
  end := Position{gridSize - 1, gridSize - 1}
  queue := []Elf{{start, []Position{start}}}
  visited := map[Position]bool{start: true}
  shortestPath := 0
  for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]
    if current.Position == end {
      shortestPath = len(current.Path) - 1 //first spot doesn't count
      break
    }
    
    for _, direction := range []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
      newPosition := current.Position.add(direction)
      if _, exists := grid[newPosition]; !exists {
        continue
      }
      if grid[newPosition] == '#' {
        continue
      }
      if visited[newPosition] {
        continue
      }
      visited[newPosition] = true
      path := copyPath(current.Path)
      newElf := Elf{newPosition, append(path, newPosition)}
      queue = append(queue, newElf)
    }
  }
  return shortestPath
}

func copyPath(path []Position) []Position {
  newPath := make([]Position, len(path))
  copy(newPath, path)
  return newPath
}

func placeBrick(grid map[Position]rune, brick Position) {
  grid[brick] = '#'
}

func (this Position) add(that Position) Position {
  return Position{this.X + that.X, this.Y + that.Y}
}

type (
  Position struct {
    X, Y int
  }
  Elf struct {
    Position Position
    Path     []Position
  }
)

func stringToInt(this string) int {
  value, _ := strconv.Atoi(this)
  return value
}

func parseInput(input string) ([]Position, map[Position]rune) {
  lines := strings.Split(input, "\n")
  //parse the grid
  bricks := make([]Position, len(lines))
  for i, line := range lines {
    xy := strings.Split(line, ",")
    bricks[i] = Position{stringToInt(xy[0]), stringToInt(xy[1])}
  }
  grid := make(map[Position]rune)
  for y := 0; y < gridSize; y++ {
    for x := 0; x < gridSize; x++ {
      grid[Position{x, y}] = '.'
    }
  }
  
  return bricks, grid
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

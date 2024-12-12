package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  //part 1
  fmt.Println(calculateFencingCosts(discoverRegions(parseInput(readInput()))))
  
  //part 2
  plots := parseInput(readInput())
  fmt.Println(calculateDiscountedFencingCosts(discoverRegions(plots), plots))
  
}

var LEFT = Position{-1, 0}
var RIGHT = Position{1, 0}
var UP = Position{0, -1}
var DOWN = Position{0, 1}
var DOWNRIGHT = Position{1, 1}
var DOWNLEFT = Position{-1, 1}
var UPLEFT = Position{-1, -1}
var UPRIGHT = Position{1, -1}

func calculateDiscountedFencingCosts(regions []*Region, plots map[Position]*Plot) int {
  totalCost := 0
  for _, region := range regions {
    findAreaAndPerimeter(region)
    findSides(region, plots)
    totalCost += region.Area * region.Sides
  }
  return totalCost
}

func calculateFencingCosts(regions []*Region) int {
  totalCost := 0
  for _, region := range regions {
    findAreaAndPerimeter(region)
    totalCost += region.Area * region.Perimeter
  }
  return totalCost
}

func findAreaAndPerimeter(region *Region) {
  region.Area = len(region.Members)
  for _, plot := range region.Members {
    //a plot starts with perimeter of 4 and removes 1 for each neighbor
    region.Perimeter += 4 - plot.Neighbors
  }
}

func findSides(region *Region, plots map[Position]*Plot) {
  sides := 0
  for _, plot := range region.Members {
    //exterior corners
    for _, direction := range [][]Position{{LEFT, UP}, {LEFT, DOWN}, {RIGHT, UP}, {RIGHT, DOWN}} {
      if !plot.isSameCharacter(plots[Position{plot.Position.X + direction[0].X, plot.Position.Y + direction[0].Y}]) &&
        !plot.isSameCharacter(plots[Position{plot.Position.X + direction[1].X, plot.Position.Y + direction[1].Y}]) {
        sides++
      }
    }
    //interior corners
    for _, direction := range [][]Position{{DOWNLEFT, DOWN, LEFT}, {DOWNRIGHT, DOWN, RIGHT}, {UPLEFT, UP, LEFT}, {UPRIGHT, UP, RIGHT}} {
      if !plot.isSameCharacter(plots[Position{plot.Position.X + direction[0].X, plot.Position.Y + direction[0].Y}]) &&
        plot.isSameCharacter(plots[Position{plot.Position.X + direction[1].X, plot.Position.Y + direction[1].Y}]) &&
        plot.isSameCharacter(plots[Position{plot.Position.X + direction[2].X, plot.Position.Y + direction[2].Y}]) {
        sides++
      }
      
    }
  }
  region.Sides = sides
}
func (this *Plot) isSameCharacter(that *Plot) bool {
  return that != nil && this.Char == that.Char
}

func discoverRegions(plots map[Position]*Plot) []*Region {
  regions := make([]*Region, 0)
  visited := make(map[Position]bool)
  for position, plot := range plots {
    //if position hasn't been visited, queue it up
    //bfs for neighbors that are same char, recording them as a region
    if _, exists := visited[position]; !exists {
      region := &Region{plot.Char, make([]*Plot, 0), 0, 0, 0}
      queue := []Position{position}
      visited[position] = true
      for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        plots[current].Region = region
        region.Members = append(region.Members, plots[current])
        for _, direction := range []Position{LEFT, RIGHT, UP, DOWN} {
          neighbor := Position{current.X + direction.X, current.Y + direction.Y}
          if _, neighborExists := plots[neighbor]; neighborExists {
            if plots[neighbor].Char == plot.Char {
              plot.Neighbors++
              if _, neighborVisited := visited[neighbor]; !neighborVisited {
                visited[neighbor] = true
                queue = append(queue, neighbor)
              }
            }
            
          }
        }
      }
      regions = append(regions, region)
    }
  }
  return regions
}

type (
  Position struct {
    X int
    Y int
  }
  Plot struct {
    Char      rune
    Neighbors int
    Position  Position
    Region    *Region
  }
  Region struct {
    Char      rune
    Members   []*Plot
    Area      int
    Perimeter int
    Sides     int
  }
)

func parseInput(input string) map[Position]*Plot {
  plots := make(map[Position]*Plot)
  plotLines := strings.Split(input, "\n")
  for y, line := range plotLines {
    for x, char := range line {
      position := Position{x, y}
      plots[position] = &Plot{char, 0, position, nil}
    }
  }
  return plots
  
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

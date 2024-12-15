package main

import (
  "bufio"
  "fmt"
  "os"
  "slices"
  "strings"
)

func main() {
  //part 1
  fmt.Println(sumGpsOfBoxes(followInstructionsToMoveBoxes(parseInput(readInput()))))
  
  //part 2
  fmt.Println(sumGpsOfBoxes(followInstructionsToMoveBoxes(parseInput(makeItBigger(readInput())))))
  
}

func sumGpsOfBoxes(boxes map[Position]rune) int {
  sum := 0
  for key, value := range boxes {
    //only count left edges of the box
    if value == BOXLEFT || value == BOX {
      sum += 100*key.Y + key.X
      
    }
  }
  return sum
}

func followInstructionsToMoveBoxes(walls map[Position]bool, robbit Position, boxes map[Position]rune, instructions []Position) map[Position]rune {
  for _, instruction := range instructions {
    //output := fmt.Sprintf("\n\n%v\n", instruction)
    robbit = moveRobbit(walls, robbit, boxes, instruction)
    //output += toString(walls, robbit, boxes)
    //writeToFile(output)
  }
  return boxes
}

func moveRobbit(walls map[Position]bool, robbit Position, boxes map[Position]rune, instruction Position) Position {
  //returns new position of robbit
  itsTryingToMove := moveIfItCan(walls, boxes, robbit, instruction)
  if itsTryingToMove == robbit {
    //robbit can't move
    return robbit
  } else {
    //robbit can move
    return itsTryingToMove
  }
}

func moveIfItCan(walls map[Position]bool, boxes map[Position]rune, position Position, instruction Position) Position {
  attemptedMovement := position.add(instruction)
  if _, exists := walls[attemptedMovement]; exists {
    //wall, can't move
    return position
  }
  if box, _ := boxes[attemptedMovement]; box != rune(0) {
    //logic is only complicated if its part 2 and moving up or down
    //if moving up/down part 2, we need to collect all the movements, but not make them until the recursion ends
    if (instruction == UP || instruction == DOWN) && boxes[attemptedMovement] != BOX {
      side1Attempts := recordMovesIfItCan(walls, boxes, attemptedMovement, instruction)
      side2Attempts := recordMovesIfItCan(walls, boxes, getRestOfBox(attemptedMovement, box), instruction)
      if len(side1Attempts) == 0 || len(side2Attempts) == 0 {
        //can't move
        return position
      }
      applyMovements(boxes, combineMaps(side1Attempts, side2Attempts), instruction)
    } else {
      //recur to move box, then move self
      tryingToPushBox := moveIfItCan(walls, boxes, attemptedMovement, instruction)
      if tryingToPushBox == attemptedMovement {
        //box can't move
        return position
      } else {
        boxes[tryingToPushBox] = boxes[attemptedMovement]
        delete(boxes, attemptedMovement)
      }
    }
    
  }
  
  return attemptedMovement
}

func recordMovesIfItCan(walls map[Position]bool, boxes map[Position]rune, position Position, instruction Position) map[Position]Position {
  attemptedMovement := position.add(instruction)
  if _, exists := walls[attemptedMovement]; exists {
    //wall, can't move
    return map[Position]Position{}
  }
  if box, _ := boxes[attemptedMovement]; box != rune(0) {
    //if there's a box, try to move it and its other side
    //if either cant move, neither can, return empty map else combine map and return that
    attempedUpdates := recordMovesIfItCan(walls, boxes, attemptedMovement, instruction)
    attemptedRestOfBoxUpdates := recordMovesIfItCan(walls, boxes, getRestOfBox(attemptedMovement, box), instruction)
    if len(attempedUpdates) == 0 || len(attemptedRestOfBoxUpdates) == 0 {
      return map[Position]Position{}
    } else {
      for key, value := range attemptedRestOfBoxUpdates {
        attempedUpdates[key] = value
      }
      combined := combineMaps(attempedUpdates, attemptedRestOfBoxUpdates)
      //this spot can also be moved
      combined[position] = attemptedMovement
      return combined
    }
  }
  return map[Position]Position{position: attemptedMovement}
}

func applyMovements(boxes map[Position]rune, updates map[Position]Position, direction Position) {
  keys := make([]Position, 0)
  for key, _ := range updates {
    keys = append(keys, key)
  }
  slices.SortFunc(keys, func(a, b Position) int {
    return b.Y - a.Y
  })
  if direction == UP {
    //reverse order
    slices.Reverse(keys)
  }
  for _, key := range keys {
    boxes[updates[key]] = boxes[key]
    delete(boxes, key)
  }
}

func combineMaps(a, b map[Position]Position) map[Position]Position {
  for key, value := range b {
    a[key] = value
  }
  return a
}

func getRestOfBox(position Position, side rune) Position {
  switch side {
  case BOXLEFT:
    return position.add(RIGHT)
  case BOXRIGHT:
    return position.add(LEFT)
  case BOX:
    return position
  }
  panic(string(side))
}

func (this Position) add(that Position) Position {
  return Position{this.X + that.X, this.Y + that.Y}
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
)

func makeItBigger(input string) string {
  input = strings.ReplaceAll(input, "#", "##")
  input = strings.ReplaceAll(input, "O", "[]")
  input = strings.ReplaceAll(input, ".", "..")
  input = strings.ReplaceAll(input, "@", "@.")
  return input
}

func parseInput(input string) (walls map[Position]bool, robbit Position, boxes map[Position]rune, instructions []Position) {
  wareHouseInstructions := strings.Split(input, "\n\n")
  //parse warehouse map
  wareHouseLines := strings.Split(wareHouseInstructions[0], "\n")
  walls = make(map[Position]bool)
  boxes = make(map[Position]rune)
  for y, line := range wareHouseLines {
    for x, char := range line {
      position := Position{x, y}
      switch char {
      case WALL:
        walls[position] = true
      case BOX:
        boxes[position] = BOX
      case BOXLEFT:
        boxes[position] = BOXLEFT
      case BOXRIGHT:
        boxes[position] = BOXRIGHT
      case ROBBIT:
        robbit = position
      }
    }
  }
  // parse instructions
  noBreaks := strings.ReplaceAll(wareHouseInstructions[1], "\n", "")
  instructions = make([]Position, len(noBreaks))
  for i, char := range noBreaks {
    var instruction Position
    switch char {
    case '^':
      instruction = UP
    case 'v':
      instruction = DOWN
    case '<':
      instruction = LEFT
    case '>':
      instruction = RIGHT
    }
    instructions[i] = instruction
  }
  
  return walls, robbit, boxes, instructions
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

func toString(walls map[Position]bool, robbit Position, boxes map[Position]rune) string {
  output := ""
  //print warehouse
  for y := 0; y < 10; y++ {
    for x := 0; x < 20; x++ {
      position := Position{x, y}
      if _, exists := walls[position]; exists {
        output += string(WALL)
      } else if char, _ := boxes[position]; char != rune(0) {
        output += string(char)
      } else if position == robbit {
        output += string(ROBBIT)
      } else {
        output += string(OPEN)
      }
    }
    output += "\n"
  }
  return output
}

func writeToFile(cycles string) {
  //f, _ := os.Create()
  f, _ := os.OpenFile("./output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  
  defer f.Close()
  
  w := bufio.NewWriter(f)
  
  fmt.Fprintln(w, cycles)
  
  w.Flush()
}

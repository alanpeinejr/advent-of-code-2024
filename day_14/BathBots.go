package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(getSafetyScore(wait(parseInput(readInput()), 100)))
  //part 2
  writeToFile(80) //grouped horizontally at 12, again at 115, vertically at 69, again at 170
  //vertical got 2 closer, so 69 - 12 = 57, 57/2 28.5. Tried it at 27, and we're 5 off now
  //at 29, we're 1 off, same with 30... add 101?
  //nope, exact same, so its in the middle? 29 + 101/2 = 79.5?
  //at 79 we're off by 2, one more aligns them?
  // its 80. 103 * 80 + 12 = 8149
  //8149 is correct
  fmt.Println()
  
}

const WIDTH = 101
const HEIGHT = 103

func getSafetyScore(robbits []*Bot) int {
  //first quadrant
  firstQuadrant := getSafetyScoreInQuadrant(robbits, 0, HEIGHT/2, 0, WIDTH/2)
  secondQuadrant := getSafetyScoreInQuadrant(robbits, 0, HEIGHT/2, WIDTH/2+1, WIDTH)
  thirdQuadrant := getSafetyScoreInQuadrant(robbits, HEIGHT/2+1, HEIGHT, 0, WIDTH/2)
  fourthQuadrant := getSafetyScoreInQuadrant(robbits, HEIGHT/2+1, HEIGHT, WIDTH/2+1, WIDTH)
  debug(robbits)
  fmt.Println(firstQuadrant, secondQuadrant, thirdQuadrant, fourthQuadrant)
  return firstQuadrant * secondQuadrant * thirdQuadrant * fourthQuadrant
}

func getSafetyScoreInQuadrant(robbits []*Bot, y1, y2, x1, x2 int) int {
  fmt.Println(y1, y2, x1, x2)
  safetyScore := 0
  for y := y1; y < y2; y++ {
    for x := x1; x < x2; x++ {
      for _, robbit := range robbits {
        if robbit.Position.X == x && robbit.Position.Y == y {
          safetyScore++
        }
      }
    }
  }
  return safetyScore
}

func print(robbits []*Bot) string {
  out := ""
  for y := 0; y < HEIGHT; y++ {
    for x := 0; x < WIDTH; x++ {
      found := false
      for _, robbit := range robbits {
        if robbit.Position.X == x && robbit.Position.Y == y {
          out += "X"
          found = true
          break
        }
      }
      if !found {
        out += "."
      }
    }
    out += "\n"
  }
  return out
}

func debug(robbits []*Bot) {
  for _, robbit := range robbits {
    fmt.Println(robbit)
  }
}

func wait(robbits []*Bot, time int) []*Bot {
  for _, robbit := range robbits {
    robbit.moveRobbit(time)
  }
  return robbits
}

func (this *Bot) moveRobbit(time int) {
  newX := (this.Position.X + this.xIncrease*time) % WIDTH
  newY := (this.Position.Y + this.yIncrease*time) % HEIGHT
  if newX < 0 {
    newX = WIDTH + newX
  }
  if newY < 0 {
    newY = HEIGHT + newY
  }
  this.Position.X = newX
  this.Position.Y = newY
}

type (
  Position struct {
    X, Y int
  }
  Bot struct {
    Position             Position
    xIncrease, yIncrease int
  }
)

func stringToInt(this string) int {
  value, _ := strconv.Atoi(this)
  return value
}

func absoluteValue(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func parseInput(input string) []*Bot {
  input = strings.ReplaceAll(input, "p=", "")
  input = strings.ReplaceAll(input, "v=", "")
  input = strings.ReplaceAll(input, " ", ",")
  lines := strings.Split(input, "\n")
  robbits := make([]*Bot, len(lines))
  for i, line := range lines {
    xyValues := strings.Split(line, ",")
    robbits[i] = &Bot{Position{stringToInt(xyValues[0]), stringToInt(xyValues[1])}, stringToInt(xyValues[2]), stringToInt(xyValues[3])}
  }
  
  return robbits
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

func writeToFile(cycle int) {
  f, _ := os.Create("./output.txt")
  defer f.Close()
  
  w := bufio.NewWriter(f)
  // thing will cycle every 103, so lets see cycles
  for i := 103*cycle - 103; i < 103*cycle; i++ {
    fmt.Fprintln(w, i)
    robits := parseInput(readInput())
    robits = wait(robits, i)
    _, _ = fmt.Fprintf(w, "%s\n\n\n\n", print(robits))
  }
  w.Flush()
}

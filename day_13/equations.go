package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  fmt.Println(getCostForPrizes(parseInput(readInput())))
  
  //part 2
  fmt.Println(oopsyDoopsy(parseInput(readInput())))
  
}

func getCostForPrizes(machines []*Machine) int {
  totalCost := 0
  for _, machine := range machines {
    findPressesNeeded(machine.A, machine.B, machine.Prize)
    totalCost += machine.A.PressesRequired * machine.A.Cost
    totalCost += machine.B.PressesRequired * machine.B.Cost
  }
  return totalCost
}

func oopsyDoopsy(machines []*Machine) int {
  for _, machine := range machines {
    machine.A.PressesRequired = 0
    machine.B.PressesRequired = 0
    machine.Prize.X += 10000000000000
    machine.Prize.Y += 10000000000000
  }
  return getCostForPrizes(machines)
}

func findPressesNeeded(a *Button, b *Button, prize Position) {
  //Ax + By = prize, x,y are ticket costs
  //A = (prize.X - (b.xIncrease)(B)) / a.xIncrease
  //B = a.yIncrease * A + b.yIncrease(B) = prize.Y
  //B = a.yIncrease * ((prize.X - (b.xIncrease)(B)) / a.xIncrease) + b.yIncrease(B) = prize.Y
  //B = ((a.yIncrease * prize.X - a.yIncrease *(b.xIncrease)(B) + (a.yIncrease * b.yIncrease(B))) / a.xIncrease) = prize.Y
  //B = prize.Y * a.xIncrease = ((a.yIncrease * prize.X - a.yIncrease *(b.xIncrease)(B) + (a.xIncrease * b.yIncrease(B)))
  //B = prize.Y * a.xIncrease - (a.yIncrease * prize.X) = (-a.yIncrease * b.xIncrease + a..xIncrease  * b.yIncrease)(B)
  //B = prize.Y * a.xIncrease - (a.yIncrease * prize.X) / (-a.yIncrease * b.xIncrease + a..xIncrease  * b.yIncrease)
  
  B := (prize.Y*a.xIncrease - (a.yIncrease * prize.X)) / (-a.yIncrease*b.xIncrease + a.xIncrease*b.yIncrease)
  A := (prize.X - (b.xIncrease * B)) / a.xIncrease
  
  if A < 0 || B < 0 || A != float64(int64(A)) || B != float64(int64(B)) {
    //no solution, negative or parital button press
    return
  }
  
  a.PressesRequired = int(A)
  b.PressesRequired = int(B)
}

type (
  Position struct {
    X, Y float64
  }
  Button struct {
    xIncrease       float64
    yIncrease       float64
    Cost            int
    PressesRequired int
  }
  Machine struct {
    A     *Button
    B     *Button
    Prize Position
  }
)

func parseInput(input string) []*Machine {
  input = strings.ReplaceAll(input, "Button A: X+", "")
  input = strings.ReplaceAll(input, "Button B: X+", "")
  input = strings.ReplaceAll(input, "Prize: X=", "")
  input = strings.ReplaceAll(input, "Y=", "")
  input = strings.ReplaceAll(input, "Y+", "")
  //now all machines are groups of csv sperated by double new lines
  machineStrings := strings.Split(input, "\n\n")
  machines := make([]*Machine, len(machineStrings))
  
  for i, machineString := range machineStrings {
    lines := strings.Split(machineString, "\n")
    aStrings := strings.Split(lines[0], ", ")
    bStrings := strings.Split(lines[1], ", ")
    prizeStrings := strings.Split(lines[2], ", ")
    a := &Button{stringToFloat(aStrings[0]), stringToFloat(aStrings[1]), 3, 0}
    b := &Button{stringToFloat(bStrings[0]), stringToFloat(bStrings[1]), 1, 0}
    prize := Position{stringToFloat(prizeStrings[0]), stringToFloat(prizeStrings[1])}
    machines[i] = &Machine{a, b, prize}
  }
  return machines
}

func stringToFloat(this string) float64 {
  value, _ := strconv.ParseFloat(this, 64)
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

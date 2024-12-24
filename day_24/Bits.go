package main

import (
  "fmt"
  "math"
  "os"
  "sort"
  "strings"
)

func main() {
  //part1
  values, gates := parseInput(readInput())
  fmt.Println(getZValues(digitalLogic(values, gates)))
  //part 2
  findALUBug(gates)
  //wpd<-->z11, jqf<-->skh, mdd <-->z19, wts<-->z37
  swaps := []string{"wpd", "z11", "jqf", "skh", "mdd", "z19", "wts", "z37"}
  sort.Strings(swaps)
  fmt.Println(strings.Join(swaps, ","))
}

func getZValues(values map[string]string) int {
  return getValueAndCalculateNumber('z', values)
}

func getValueAndCalculateNumber(prefix rune, values map[string]string) int {
  prefixValues := make([]string, 0)
  for k, _ := range values {
    if rune(k[0]) == prefix {
      prefixValues = append(prefixValues, k)
    }
  }
  sort.Strings(prefixValues)
  sum := 0
  for i, value := range prefixValues {
    if values[value] == "1" {
      sum += int(math.Pow(2, float64(i)))
    }
  }
  return sum
}

func findALUBug(gates map[string]*Gate) {
  //we've got a full adder
  // x xor y = xy
  // xy xor cIN = s --a z value
  // xy and cIN = cA
  // x and y = cB
  // cA or cB = cOUT
  cIn := ""
  cOut := ""
  cA := ""
  
  //checked, there's 46 z's
  for i := 0; i < 45; i++ {
    xI := "x" + indexToGateID(i)
    yI := "y" + indexToGateID(i)
    xy := findGateForOperation(xI, yI, "XOR", gates, i)
    if cIn != "" {
      sI := findGateForOperation(xy, cIn, "XOR", gates, i)
      if sI != "z"+indexToGateID(i) {
        fmt.Println(i, sI)
        panic("this is wrong, it should output to zI")
      }
      cA = findGateForOperation(xy, cIn, "AND", gates, i)
    }
    cB := findGateForOperation(xI, yI, "AND", gates, i)
    
    if cA != "" {
      cOut = findGateForOperation(cA, cB, "OR", gates, i)
    } else {
      //will only happen once
      cOut = cB
    }
    cIn = cOut
  }
  return
}
func indexToGateID(index int) string {
  return fmt.Sprintf("%0*d", 2, index)
}

func findGateForOperation(input1, input2, operation string, gates map[string]*Gate, index int) string {
  for k, v := range gates {
    //order in inputs doesnt matter
    if ((v.Input1 == input1 && v.Input2 == input2) || (v.Input1 == input2 && v.Input2 == input1)) && v.Operation == operation {
      return k
    }
  }
  //suspect i'll get errors here too
  fmt.Println(operation, input1, input2, index)
  panic("whoops")
}

func digitalLogic(values map[string]string, gates map[string]*Gate) map[string]string {
  queue := make([]string, 0)
  for k, _ := range gates {
    queue = append(queue, k)
  }
  for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]
    gate := gates[current]
    value1, input1Available := values[gate.Input1]
    value2, input2Available := values[gate.Input2]
    if !(input1Available && input2Available) {
      queue = append(queue, current)
      continue
    }
    values[current] = doOperation(gate.Operation, value1, value2)
  }
  return values
}

func doOperation(operation string, input1 string, input2 string) string {
  switch operation {
  case "AND":
    if input1 == "1" && input2 == "1" {
      return "1"
    }
  case "OR":
    if input1 == "1" || input2 == "1" {
      return "1"
    }
  case "XOR":
    if input1 != input2 {
      return "1"
    }
  }
  return "0"
}

type (
  Gate struct {
    Operation string
    Input1    string
    Input2    string
    Output    string
  }
)

func parseInput(input string) (map[string]string, map[string]*Gate) {
  values := map[string]string{}
  input = strings.ReplaceAll(input, "-> ", "")
  initialsOperations := strings.Split(input, "\n\n")
  initals := strings.Split(initialsOperations[0], "\n")
  for _, initial := range initals {
    nameValue := strings.Split(initial, ": ")
    values[nameValue[0]] = nameValue[1]
  }
  gates := map[string]*Gate{}
  for _, operation := range strings.Split(initialsOperations[1], "\n") {
    oneOpTwoWhere := strings.Split(operation, " ")
    gate := &Gate{oneOpTwoWhere[1], oneOpTwoWhere[0], oneOpTwoWhere[2], oneOpTwoWhere[3]}
    gates[gate.Output] = gate
  }
  
  return values, gates
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

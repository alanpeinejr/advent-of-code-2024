package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

func main() {
  //part 1
  runProgram(parseInput(readInput()))
  fmt.Println(Output)
  
  //part 2
  //outputInput()
  //fmt.Println(math.Pow(2, 17*8))
}

//manual brute force/depth. Saw quickly that number of outputs scaled with size so incremented by 1 trillion. Checked the ends because they're determined by
//most significant digits. When they start matching we can break, start checking for earlier instructions, and consider changing the sizeFactor
//this could be programitcally down but eh.
func outputInput() {
  instructions := parseInput(readInput())
  input := ""
  for i := 0; i < len(instructions); i++ {
    input += string(instructions[i]) + ","
  }
  sizeFactor := 1
  foundChange := 190384609485312
  for i := foundChange; i < 1000000000000000; i += sizeFactor {
    A.Value = i
    B.Value = 0
    C.Value = 0
    Output = ""
    runProgram(instructions)
    fmt.Println(i, Output)
    if Output == input {
      break
    }
    
    lookingBack := 32
    if len(Output) == len(input) && Output[len(Output)-lookingBack:] == input[len(input)-lookingBack:] {
      fmt.Println(Output[len(Output)-lookingBack:])
      break
    }
    
  }
}

var (
  A      = &Register{'a', 0}
  B      = &Register{'b', 0}
  C      = &Register{'c', 0}
  Output = ""
)

func operandToValue(operand rune) int {
  switch operand {
  case '0':
    return 0
  case '1':
    return 1
  case '2':
    return 2
  case '3':
    return 3
  case '4':
    return A.Value
  case '5':
    return B.Value
  case '6':
    return C.Value
  case '7':
    fallthrough
  default:
    panic("whoops")
  }
}

func runProgram(instructions []rune) {
  index := 0
  for index < len(instructions)-1 {
    operator := instructions[index]
    operand := instructions[index+1]
    index = operation(operator, operand, index)
  }
}

func operation(operator rune, operand rune, index int) (nextIndex int) {
  switch operator {
  case '0':
    //adv
    A.Value = A.Value / int(math.Pow(2, float64(operandToValue(operand))))
    index += 2
  case '1':
    B.Value = B.Value ^ runeToInt(operand)
    index += 2
  case '2':
    B.Value = operandToValue(operand) % 8
    index += 2
  case '3':
    if A.Value == 0 {
      index += 2
    } else {
      index = runeToInt(operand)
    }
  case '4':
    B.Value = B.Value ^ C.Value
    index += 2
  case '5':
    Output += intToString(operandToValue(operand)%8) + ","
    index += 2
  case '6':
    B.Value = A.Value / int(math.Pow(2, float64(operandToValue(operand))))
    index += 2
  case '7':
    C.Value = A.Value / int(math.Pow(2, float64(operandToValue(operand))))
    index += 2
  default:
    panic(string(operator) + " is not a valid operator")
  }
  return index
}

type (
  Register struct {
    Name  rune
    Value int
  }
)

func stringToInt(this string) int {
  value, _ := strconv.Atoi(this)
  return value
}

func intToString(this int) string {
  return strconv.Itoa(this)
}

func runeToInt(this rune) int {
  return stringToInt(string(this))
}

func parseInput(input string) []rune {
  input = strings.ReplaceAll(input, "Register A: ", "")
  input = strings.ReplaceAll(input, "Register B: ", "")
  input = strings.ReplaceAll(input, "Register C: ", "")
  input = strings.ReplaceAll(input, "Program: ", "")
  input = strings.ReplaceAll(input, ",", "")
  registerInstructions := strings.Split(input, "\n\n")
  abc := strings.Split(registerInstructions[0], "\n")
  A.Value = stringToInt(abc[0])
  B.Value = stringToInt(abc[1])
  C.Value = stringToInt(abc[2])
  
  instructions := make([]rune, len(registerInstructions[1]))
  for i, instruction := range registerInstructions[1] {
    instructions[i] = instruction
  }
  
  return instructions
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

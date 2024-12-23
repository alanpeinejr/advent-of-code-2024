package main

import (
  "fmt"
  "os"
  "sort"
  "strings"
)

func main() {
  //part1
  fmt.Println(findHistorian(findConnections(parseInput(readInput()))))
  //part 2
  largest := findLargestNetwork(parseInput(readInput()))
  fmt.Println(joinStrings(largest))
}

func joinStrings(array []string) string {
  return strings.Join(array, ",")
}

func findLargestNetwork(computers map[string]*Computer) []string {
  largest := map[string]bool{}
  for k, _ := range computers {
    networked := findNetworks(k, computers, map[string]bool{})
    if len(networked) > len(largest) {
      largest = networked
    }
  }
  array := make([]string, 0)
  for k, _ := range largest {
    array = append(array, k)
  }
  sort.Strings(array)
  return array
}

func findNetworks(id string, computers map[string]*Computer, networked map[string]bool) map[string]bool {
  networked[id] = true
  for connectionId, _ := range computers[id].Connections {
    if connectionId != id {
      add := true
      //when we see a computer that has every networked computer so far, add it
      //the last computer added has all the previous ones that had all networks
      for networkedId, _ := range networked {
        if _, exists := computers[networkedId].Connections[connectionId]; !exists {
          add = false
          break
        }
      }
      
      if add {
        networked[connectionId] = true
      }
    }
  }
  return networked
}

func findHistorian(triples map[Triple]int) int {
  count := 0
  for k, v := range triples {
    if v >= 6 && k.hasValueStartsWithT() {
      count++
    }
  }
  return count
}
func (this Triple) hasValueStartsWithT() bool {
  return strings.HasPrefix(this.One, "t") || strings.HasPrefix(this.Two, "t") || strings.HasPrefix(this.Three, "t")
}

func findConnections(computers map[string]*Computer) map[Triple]int {
  triples := map[Triple]int{}
  //find every triple
  //loop through the triple, and created a triple in map with keys ordered alphabetically
  for k, v := range computers {
    for k2, v2 := range v.Connections {
      for k3, _ := range v2.Connections {
        unsortedTripleKeys := []string{k, k2, k3}
        sort.Strings(unsortedTripleKeys)
        triple := Triple{unsortedTripleKeys[0], unsortedTripleKeys[1], unsortedTripleKeys[2]}
        if _, exists := triples[triple]; !exists {
          triples[triple] = 0
        }
        triples[triple]++
      }
    }
  }
  return triples
}

type (
  Computer struct {
    ID          string
    Connections map[string]*Computer
  }
  Triple struct {
    One, Two, Three string
  }
)

func parseInput(input string) map[string]*Computer {
  lines := strings.Split(input, "\n")
  computers := make(map[string]*Computer)
  for _, line := range lines {
    oneTwo := strings.Split(line, "-")
    if _, exists := computers[oneTwo[0]]; !exists {
      computers[oneTwo[0]] = &Computer{oneTwo[0], make(map[string]*Computer)}
    }
    if _, exists := computers[oneTwo[1]]; !exists {
      computers[oneTwo[1]] = &Computer{oneTwo[1], make(map[string]*Computer)}
    }
    computers[oneTwo[0]].Connections[oneTwo[1]] = computers[oneTwo[1]]
    computers[oneTwo[1]].Connections[oneTwo[0]] = computers[oneTwo[0]]
  }
  return computers
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

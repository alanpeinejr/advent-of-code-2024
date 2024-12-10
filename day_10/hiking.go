package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part 1
	fmt.Println(sumTrailHeadScores(findTrailHeads(parseInput(readInput()))))

	//part 2
	fmt.Println(sumTrailHeadRatings(findTrailHeads(parseInput(readInput()))))

}

func sumTrailHeadScores(trailHeads map[*Location]int) int {
	sum := 0
	for trailHead, _ := range trailHeads {
		sum += len(trailHead.EndsReachable)
	}
	return sum
}

func sumTrailHeadRatings(trailHeads map[*Location]int) int {
	sum := 0
	for _, rating := range trailHeads {
		sum += rating
	}
	return sum
}

func findTrailHeads(locations map[Position]*Location, nines []*Location) map[*Location]int {
	trailHeads := map[*Location]int{}
	queue := make([]*Location, len(nines))
	copy(queue, nines)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		//check all directions
		if current.Height == 0 {
			//trail head
			if _, exists := trailHeads[current]; !exists {
				trailHeads[current] = 0

			}
			//every time we visited a trail head, it means we got there a different way, thus, its rating
			trailHeads[current]++
			continue
		}

		for _, direction := range []Position{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			newX, newY := current.Position.X+direction.X, current.Position.Y+direction.Y
			newPosition := Position{newX, newY}
			if location, exists := locations[newPosition]; exists {
				if current.Height-location.Height == 1 {
					//we can reach this location
					queue = append(queue, locations[newPosition])
					for position, _ := range current.EndsReachable {
						locations[newPosition].EndsReachable[position] = true
					}
				}
			}
		}
	}

	return trailHeads

}

type (
	Location struct {
		Height        int
		Position      Position
		EndsReachable map[Position]bool
	}
	Position struct {
		X int
		Y int
	}
)

func parseInput(input string) (map[Position]*Location, []*Location) {
	locations := map[Position]*Location{}
	nines := make([]*Location, 0)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, char := range line {
			position := Position{x, y}
			locations[position] = &Location{stringToInt(string(char)), position, make(map[Position]bool)}
			if char == '9' {
				nines = append(nines, locations[position])
				//self is reachable
				locations[position].EndsReachable[position] = true
			}
		}
	}
	return locations, nines
}

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}

func intToString(this int) string {
	return strconv.Itoa(this)
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

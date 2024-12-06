package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part 1
	station, guard := parseInput(readInput())
	positionsVisited := patrol(station, guard)

	fmt.Println(countPositionsVisited(positionsVisited))

	//part 2
	fmt.Println(findLoops(positionsVisited, station, guard))

}

func countPositionsVisited(positions map[Position]bool) int {
	return len(positions)
}

func findLoops(positionsVisited map[Position]bool, station [][]rune, guard *Guard) int {
	//set the guard to start

	loops := map[Position]bool{}
	for key, _ := range positionsVisited {
		//put an obstacle there, set the guard off, record new positions, if we hit a new position seen, its a loop
		station[key.Y][key.X] = '#'
		guard.Position = guard.Start
		guard.Direction = UP
		newPositions := map[PositionVisited]bool{}
		for true {
			visitedPosition, leftStation := moveGuard(station, guard)
			if leftStation {
				break
			}
			newVisitedPosition := PositionVisited{visitedPosition, guard.Direction}
			if newPositions[newVisitedPosition] {
				loops[Position{key.X, key.Y}] = true
				break
			}
			newPositions[newVisitedPosition] = true
		}

		station[key.Y][key.X] = '.' //reset
	}
	//check the start position and adjust manually if it made a loop
	fmt.Println(loops[guard.Start])
	return len(loops)
}

func patrol(station [][]rune, guard *Guard) map[Position]bool {
	positionsVisited := map[Position]bool{guard.Position: true}
	for true {
		visitedPosition, leftStation := moveGuard(station, guard)
		if leftStation {
			break
		}
		positionsVisited[visitedPosition] = true

	}
	return positionsVisited
}

func moveGuard(positions [][]rune, guard *Guard) (position Position, leftStation bool) {
	leftStation = false
	//find guards next position
	nextPosition := guard.Position.move(guard.Direction)
	if nextPosition.isInBounds(positions) {
		//if its #, turn right, return
		if nextPosition.isObstacle(positions) {
			guard.Direction = guard.Direction.turnRight()
			position = guard.Position
		} else {
			position = nextPosition
		}
	} else {
		leftStation = true
	}
	guard.Position = position

	return position, leftStation
}

func (this Position) move(direction Position) Position {
	return Position{this.X + direction.X, this.Y + direction.Y}
}
func (this Position) isInBounds(positions [][]rune) bool {
	return this.Y >= 0 && this.Y < len(positions) && this.X >= 0 && this.X < len(positions[this.Y])
}
func (this Position) isObstacle(positions [][]rune) bool {
	return positions[this.Y][this.X] == '#'
}
func (this Position) isEqual(position Position) bool {
	return this.X == position.X && this.Y == position.Y
}

func (this Position) turnRight() Position {
	switch this {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	panic("whoops")
}

type (
	Position struct {
		X int
		Y int
	}
	PositionVisited struct {
		Position  Position
		Direction Position
	}
	Guard struct {
		Direction Position
		Position  Position
		Start     Position
	}
)

var LEFT = Position{-1, 0}
var RIGHT = Position{1, 0}
var UP = Position{0, -1}
var DOWN = Position{0, 1}

func parseInput(input string) ([][]rune, *Guard) {
	lines := strings.Split(input, "\n")
	positions := make([][]rune, len(lines))
	guard := new(Guard)
	for y, line := range lines {
		position := make([]rune, len(line))
		for x, char := range line {
			position[x] = char
			if char == '^' {
				guard.Direction = UP
				guard.Position = Position{x, y}
				guard.Start = guard.Position
				position[x] = '.'
			}
		}
		positions[y] = position
	}
	return positions, guard
}

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
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

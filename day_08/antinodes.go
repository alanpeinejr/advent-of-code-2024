package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part 1
	fmt.Println(len(findAntinodes(parseInput(readInput()))))

	//part 2
	fmt.Println(len(findHarmonicAntinodes(parseInput(readInput()))))

}

func findHarmonicAntinodes(locations [][]*Location, antennaLocationMap map[rune][]*Location) map[Position]bool {
	uniqueAntinodes := map[Position]bool{}
	//for each unique antenna, find all the antenna locations,
	for char, antennas := range antennaLocationMap {
		for i, location := range antennas {
			if len(antennas) != 1 {
				//add self
				uniqueAntinodes[Position{location.X, location.Y}] = true
			}
			for j, other := range antennas {
				if i == j {
					continue
				}
				antinodeXDiff, antinodeYDiff := location.X-other.X, location.Y-other.Y
				//we'll loop until we break bounds, doing part 1 at each range
				harmonicRange := 1
				for true {
					if (location.Y+(antinodeYDiff*harmonicRange)) < 0 || (location.Y+(antinodeYDiff*harmonicRange)) >= len(locations) || (location.X+(antinodeXDiff*harmonicRange)) < 0 || (location.X+(antinodeXDiff*harmonicRange)) >= len(locations[0]) {
						break //out of bounds antinode
					}
					antinodeLocation := locations[location.Y+(antinodeYDiff*harmonicRange)][location.X+(antinodeXDiff*harmonicRange)]
					antinodeLocation.Antinodes = append(antinodeLocation.Antinodes, char)
					uniqueAntinodes[Position{antinodeLocation.X, antinodeLocation.Y}] = true
					harmonicRange += 1
				}

			}
		}
	}
	return uniqueAntinodes

}

func findAntinodes(locations [][]*Location, antennaLocationMap map[rune][]*Location) map[Position]bool {
	uniqueAntinodes := map[Position]bool{}
	//for each unique antenna, find all the antenna locations,
	for char, antennas := range antennaLocationMap {
		for i, location := range antennas {
			for j, other := range antennas {
				if i == j {
					continue
				}
				antinodeXDiff, antinodeYDiff := location.X-other.X, location.Y-other.Y
				if (location.Y+antinodeYDiff) < 0 || (location.Y+antinodeYDiff) >= len(locations) || (location.X+antinodeXDiff) < 0 || (location.X+antinodeXDiff) >= len(locations[0]) {
					continue //out of bounds antinode
				}
				antinodeLocation := locations[location.Y+antinodeYDiff][location.X+antinodeXDiff]
				antinodeLocation.Antinodes = append(antinodeLocation.Antinodes, char)
				uniqueAntinodes[Position{antinodeLocation.X, antinodeLocation.Y}] = true
			}
		}
	}
	return uniqueAntinodes

}

type (
	Location struct {
		Char      rune
		X         int
		Y         int
		Antinodes []rune
	}
	Position struct {
		X int
		Y int
	}
)

func parseInput(input string) ([][]*Location, map[rune][]*Location) {
	lines := strings.Split(input, "\n")
	locations := make([][]*Location, len(lines))
	uniqueAntenna := map[rune][]*Location{}
	for y, line := range lines {
		location := make([]*Location, len(line))
		for x, char := range line {
			location[x] = &Location{char, x, y, make([]rune, 0)}
			if char != '.' {
				if _, exists := uniqueAntenna[char]; !exists {
					uniqueAntenna[char] = make([]*Location, 0)
				}
				uniqueAntenna[char] = append(uniqueAntenna[char], location[x])
			}
		}
		locations[y] = location
	}
	return locations, uniqueAntenna
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

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	disk1, disk2 := parseInput(readInput())
	//part 1
	fmt.Println(checkSum(reOrderBlocks(disk1)))

	//part 2
	fmt.Println(checkSum(reOrderBlocks2(disk2)))

}

func printString(this []Block) {
	result := ""
	for _, block := range this {
		for i := 0; i < block.Length; i++ {
			if block.Id != -1 {
				result += intToString(block.Id)
			} else {
				result += "."
			}
		}
		result += " "

	}
	fmt.Println(result)
}

func checkSum(disk []Block) int {
	sum := 0
	blockLengths := 0
	for _, block := range disk {
		for j := 0; j < block.Length; j++ {
			if block.Id != -1 {
				sum += (blockLengths + j) * block.Id
			}
		}
		blockLengths += block.Length

	}
	return sum
}

func reOrderBlocks2(disk []Block) []Block {
	processed := map[int]bool{}
	lastProcessedIndex := len(disk) - 1
	for lastProcessedIndex > 0 {
		if processed[disk[lastProcessedIndex].Id] && disk[lastProcessedIndex].Id != -1 {
			//dont reprocess Id's
			lastProcessedIndex--
			continue
		}
		//found a new block, print and mark as processed
		processed[disk[lastProcessedIndex].Id] = true
		//always start from the beginning
		lastOpenEmpty := -1
		for lastProcessedIndex > lastOpenEmpty {
			nextOpenEmpty := findNextEmpty(disk, lastOpenEmpty)
			if nextOpenEmpty > lastProcessedIndex {
				break
			}

			if disk[lastProcessedIndex].Id != -1 && disk[lastProcessedIndex].Length <= disk[nextOpenEmpty].Length {
				//swap and fill in blanks leftover
				leftover := disk[nextOpenEmpty].Length - disk[lastProcessedIndex].Length
				disk[lastProcessedIndex], disk[nextOpenEmpty] = disk[nextOpenEmpty], disk[lastProcessedIndex]
				//the empty Block needs shortened
				//insert leftover space
				if leftover > 0 {
					disk[lastProcessedIndex].Length -= leftover
					disk = slices.Insert(disk, nextOpenEmpty+1, Block{-1, leftover})
					//we just increased the length of the array
					lastProcessedIndex++
				}
				//go to next block
				break
			}
			lastOpenEmpty = nextOpenEmpty
		}
		lastProcessedIndex--
	}
	return mergeBlanks(disk)
}
func mergeBlanks(disk []Block) []Block {
	newDisk := make([]Block, 0)
	i := 0
	for i < len(disk)-1 {
		newDisk = append(newDisk, disk[i])
		for i+1 < len(disk) && newDisk[len(newDisk)-1].Id == -1 && disk[i+1].Id == -1 && i < len(disk) {
			newDisk[len(newDisk)-1].Length += disk[i+1].Length
			i += 1
		}
		i += 1
	}
	return newDisk
}

func reOrderBlocks(disk []Block) []Block {
	lastOpenEmpty := -1
	lastProcessedIndex := len(disk) - 1
	for lastProcessedIndex > lastOpenEmpty {
		if disk[lastProcessedIndex].Id == -1 {
			//empty, skip
			lastProcessedIndex--
			continue
		}
		// find Next Empty
		nextOpenEmpty := findNextEmpty(disk, lastOpenEmpty)
		if nextOpenEmpty > lastProcessedIndex {
			break
		}
		//swap
		disk[nextOpenEmpty], disk[lastProcessedIndex] = disk[lastProcessedIndex], disk[nextOpenEmpty]
		//updated, so go to next
		lastOpenEmpty = nextOpenEmpty
		//decrement
		lastProcessedIndex--

	}
	return disk
}

func findNextEmpty(disk []Block, lastEmpty int) int {
	for i := lastEmpty + 1; i < len(disk); i++ {
		if disk[i].Id == -1 {
			return i
		}
	}
	return len(disk)
}

type (
	Block struct {
		Id     int
		Length int
	}
)

func parseInput(input string) ([]Block, []Block) {
	disk := make([]Block, 0)
	disk2 := make([]Block, 0)
	for i, char := range input {
		if i%2 == 0 {
			//even, fileSize
			for j := 0; j < stringToInt(string(char)); j++ {
				disk = append(disk, Block{i / 2, 1})
				if j == 0 {
					disk2 = append(disk2, Block{i / 2, stringToInt(string(char))})

				}
			}
		} else {
			//free space
			for j := 0; j < stringToInt(string(char)); j++ {
				disk = append(disk, Block{-1, 1})
				if j == 0 {
					disk2 = append(disk2, Block{-1, stringToInt(string(char))})

				}
			}

		}
	}
	return disk, disk2
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

package advent20201224

import (
	"bufio"
	"fmt"
	"os"
)

type Offset struct {
	North, East int
}

func (o Offset) String() string {
	return fmt.Sprintf("<N: %v E:%v>", o.North, o.East)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewOffset(North, East int) Offset {
	return Offset{North: North, East: East}
}

func RecordsFromFile(filename string) (results []string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}
	return
}

func Canonicalize(input string) Offset {
	runes := []rune(input)
	offset := Offset{}
	for len(runes) > 0 {
		switch runes[0] {
		case 'e':
			offset.East += 2
			runes = runes[1:]
		case 'w':
			offset.East -= 2
			runes = runes[1:]
		case 'n':
			offset.North++
			runes = runes[1:]
			switch runes[0] {
			case 'e':
				offset.East++
				runes = runes[1:]
			case 'w':
				offset.East--
				runes = runes[1:]
			}
		case 's':
			offset.North--
			runes = runes[1:]
			switch runes[0] {
			case 'e':
				offset.East++
				runes = runes[1:]
			case 'w':
				offset.East--
				runes = runes[1:]
			}
		}
	}
	return offset
}

func UniqueCount(input []string) map[Offset]int {
	countMap := make(map[Offset]int)
	for _, value := range input {
		canon := Canonicalize(value)
		countMap[canon]++
	}
	return countMap
}

// Part2 solves part2
func Part1(filename string) (blackTiles int) {
	data := RecordsFromFile(filename)
	countMap := UniqueCount(data)
	for _, value := range countMap {
		if value%2 == 1 {
			blackTiles++
		}
	}
	return
}

// Part2 solves part2
func Part2(filename string) int {
	return 0
}

package advent20201224

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Offset struct {
	North, East int
}

type TileMap map[Offset]int

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

func UniqueCount(input []string) TileMap {
	countMap := make(TileMap)
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
	return len(BlackOnlyMap(countMap))
}

func (o Offset) Neighbors() (results []Offset) {
	return []Offset{
		Offset{East: o.East + 1, North: o.North + 1},
		Offset{East: o.East + 2, North: o.North},
		Offset{East: o.East + 1, North: o.North - 1},
		Offset{East: o.East - 1, North: o.North + 1},
		Offset{East: o.East - 2, North: o.North},
		Offset{East: o.East - 1, North: o.North - 1},
	}
}

func (tm TileMap) ctBlackNeighbors(o Offset) (ct int) {
	for _, offset := range o.Neighbors() {
		ct += tm[offset]
	}
	return
}

func (tm TileMap) BlackTiles() (tiles []Offset) {
	for k := range tm {
		tiles = append(tiles, k)
	}
	return
}

func (tm TileMap) WhiteTiles() (tiles []Offset) {
	for o := range tm {
		for _, neighbor := range o.Neighbors() {
			if tm[neighbor] == 0 {
				log.Print(o, neighbor, tm[neighbor])
				tiles = append(tiles, neighbor)
			}
		}
	}
	return
}

func Iterate(tm TileMap) TileMap {
	returnMap := make(TileMap)
	for _, offset := range tm.BlackTiles() {
		blackNeighbors := tm.ctBlackNeighbors(offset)
		if blackNeighbors == 1 || blackNeighbors == 2 {
			returnMap[offset] = 1
		}
	}
	// fmt.Println(tm.WhiteTiles())
	log.Print("White tiles: ", tm.WhiteTiles())
	for _, offset := range tm.WhiteTiles() {
		blackNeighbors := tm.ctBlackNeighbors(offset)
		log.Print(offset, blackNeighbors)
		if blackNeighbors == 2 {
			returnMap[offset] = 1
		}
	}
	log.Print(len(returnMap), returnMap)
	return returnMap
}

// Part2 solves part2
func Part2(filename string, times int) int {
	data := RecordsFromFile(filename)
	countMap := UniqueCount(data)
	answerMap := BlackOnlyMap(countMap)
	log.Print(len(answerMap), answerMap)
	for i := 1; i <= times; i++ {
		answerMap = Iterate(answerMap)
	}
	return len(answerMap)
}

func BlackOnlyMap(tm TileMap) TileMap {
	blackMap := make(TileMap)
	for k, v := range tm {
		if v%2 == 1 {
			blackMap[k] = 1
		}
	}
	return blackMap
}

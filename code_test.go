package advent20201224_test

import (
	advent "advent20201224"
	"fmt"
	"testing"
)

func TestRecordsFromFile(t *testing.T) {
	advent.RecordsFromFile("sample.txt")
	advent.RecordsFromFile("input.txt")
	t.Fail()
}

func TestPart1(t *testing.T) {
	if advent.Part1("sample.txt") != 10 {
		t.Fail()
	}
	fmt.Println(advent.Part1("input.txt"))
	t.Fail()
}

func TestCanonical(t *testing.T) {
	if advent.Canonicalize("nwwswee") != advent.NewOffset(0, 0) {
		t.Fail()
	}
}

// func TestPart2(t *testing.T) {
// 	fmt.Println(advent.Part2("sample.txt"))
// 	fmt.Println(advent.Part2("input.txt"))
// 	t.Fail()
// }

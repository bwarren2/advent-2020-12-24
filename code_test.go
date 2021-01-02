
import (
	advent "XXX"
	"fmt"
	"testing"
)

func TestRecodsFromFile(t *testing.T) {
	advent.RecordsFromFile("sample.txt")
	advent.RecordsFromFile("input.txt")
	// t.Fail()
}

func TestPart1(t *testing.T) {
	fmt.Println(advent.Part1("sample.txt"))
	fmt.Println(advent.Part1("input.txt"))
	// t.Fail()
}

func TestPart2(t *testing.T) {
	fmt.Println(advent.Part2("sample.txt"))
	fmt.Println(advent.Part2("input.txt"))
	t.Fail()
}

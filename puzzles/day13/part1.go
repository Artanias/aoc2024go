package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	exampleFilePath string = "puzzles/day13/example.txt"
	dataFilePath    string = "puzzles/day13/data.txt"
)
const (
	maxPress  uint64 = 100
	costPushA uint64 = 3
	costPushB uint64 = 1
)

const (
	buttonRe string = `X\+(\d+), Y\+(\d+)`
	prizeRe  string = `Prize:\s+X=(\d+),\s+Y=(\d+)`
)

type Vertex struct {
	x int
	y int
}

func getFileContent(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(contents), "\n")
}

func stringToNumber(number string) int {
	int_number, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return int_number
}

func extractNumbers(s string, rex string) (Vertex, error) {
	re := regexp.MustCompile(rex)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		return Vertex{0, 0}, fmt.Errorf("failed to find numbers from: ", matches)
	}

	return Vertex{stringToNumber(matches[1]), stringToNumber(matches[2])}, nil
}

func calcResult(filePath string) (result uint64, err error) {
	lines := strings.Split(getFileContent(filePath), "\n\n")
	for _, clawMap := range lines {
		var buttonA, buttonB, prize Vertex
		for i, line := range strings.Split(clawMap, "\n") {
			switch i {
			case 0:
				buttonA, err = extractNumbers(line, buttonRe)
				if err != nil {
					return 0, err
				}
			case 1:
				buttonB, err = extractNumbers(line, buttonRe)
				if err != nil {
					return 0, err
				}
			case 2:
				prize, err = extractNumbers(line, prizeRe)
				if err != nil {
					return 0, err
				}
			default:
				return 0, fmt.Errorf("Invalid index, expected from 0 to 2, got: ", i)
			}
		}
		var interRes = Vertex{}
		for cntA := 1; cntA <= int(maxPress); cntA++ {
			for cntB := 1; cntB <= int(maxPress); cntB++ {
				interRes = Vertex{
					buttonB.x*cntB + buttonA.x*cntA,
					buttonB.y*cntB + buttonA.y*cntA,
				}
				if interRes.x == prize.x && interRes.y == prize.y {
					result += uint64(cntA)*costPushA + uint64(cntB)*costPushB
					break
				}
			}
		}
	}
	return result, nil
}

func main() {
	exampleResult, err := calcResult(exampleFilePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Example result: ", exampleResult)
	}
	dataResult, err := calcResult(dataFilePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data result: ", dataResult)
	}
}

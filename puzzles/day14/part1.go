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
	exampleFilePath string = "puzzles/day14/example.txt"
	exampleWidth    int    = 11
	exampleTall     int    = 7
)

const (
	dataFilePath string = "puzzles/day14/data.txt"
	dataWidth    int    = 101
	dataTall     int    = 103
)

const (
	secondsElapsed int = 100
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

func extractNumbers(s string) (Vertex, Vertex, error) {
	re := regexp.MustCompile(`p=(\d+),(\d+)\s+v=(-?\d+),(-?\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 5 {
		return Vertex{0, 0}, Vertex{0, 0}, fmt.Errorf("failed to find numbers from: ", matches)
	}

	return Vertex{stringToNumber(matches[1]), stringToNumber(matches[2])},
		Vertex{stringToNumber(matches[3]), stringToNumber(matches[4])}, nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func calcResult(filePath string, width int, tall int) (result int, err error) {
	excludedX := width / 2
	excludedY := tall / 2
	field := make(map[Vertex]int)
	for _, line := range strings.Split(getFileContent(filePath), "\n") {
		p, v, err := extractNumbers(line)
		if err != nil {
			return 0, err
		}
		lastPos := Vertex{mod(p.x+v.x*secondsElapsed, width), mod(p.y+v.y*secondsElapsed, tall)}
		if lastPos.x == excludedX || lastPos.y == excludedY {
			continue
		}
		field[lastPos] += 1
	}
	result = 1
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for k, v := range field {
		if k.x < excludedX && k.y < excludedY {
			q1 += v
		} else if k.x < excludedX && k.y > excludedY {
			q2 += v
		} else if k.x > excludedX && k.y < excludedY {
			q3 += v
		} else {
			q4 += v
		}
	}
	return q1 * q2 * q3 * q4, nil
}

func main() {
	exampleResult, err := calcResult(exampleFilePath, exampleWidth, exampleTall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Example result: ", exampleResult)
	}
	dataResult, err := calcResult(dataFilePath, dataWidth, dataTall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data result: ", dataResult)
	}
}

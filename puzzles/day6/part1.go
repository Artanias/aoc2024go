package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)


const exampleFilePath string = "puzzles/day6/example.txt"
const dataFilePath string = "puzzles/day6/data.txt"
type Direction int64
const (
	Up Direction = iota
	Down
	Left
	Right
)


func getFileContent(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(contents), "\n")
}

func findStartPosition(lines []string, rows int, columns int) (int, int) {
	searched_sym := "^"
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if lines[i][j:j+1] == searched_sym {
				return i, j
			}
		}
	}
	panic("Can't find start position.")
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func calcResult(filePath string) int {
	result := 0
	lines := strings.Split(getFileContent(filePath), "\n")
	lines_cnt := len(lines)
	columns_cnt := len(lines[0])
	i, j := findStartPosition(lines, lines_cnt, columns_cnt)
	lines[i] = replaceAtIndex(lines[i], 'X', j)
	direction := Up
	for ;; {
		if direction == Up {
			if lines[i-1][j:j+1] == "#" {
				direction = Right
				continue
			}
			i -= 1
		} else if direction == Right {
			if lines[i][j+1:j+2] == "#" {
				direction = Down
				continue
			}
			j += 1
		} else if direction == Down {
			if lines[i+1][j:j+1] == "#" {
				direction = Left
				continue
			}
			i += 1
		} else {
			if lines[i][j-1:j] == "#" {
				direction = Up
				continue
			}
			j -= 1
		}
		lines[i] = replaceAtIndex(lines[i], 'X', j)
		if i == (lines_cnt - 1) || j == (columns_cnt - 1) || i == 0 || j == 0 {
			break
		}
	}
	for i := 0; i < lines_cnt; i++ {
		for j := 0; j < columns_cnt; j++ {
			if lines[i][j:j+1] == "X" {
				result += 1
			}
		}
	}
	return result
}

func main() {
	example_result := calcResult(exampleFilePath)
	expected := 41
	if example_result != expected {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected, example_result))
	}
	fmt.Printf("Example result: %d.\n", example_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

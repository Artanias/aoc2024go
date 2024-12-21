package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"slices"
)


const example1FilePath string = "puzzles/day10/example1.txt"
const example2FilePath string = "puzzles/day10/example2.txt"
const example3FilePath string = "puzzles/day10/example3.txt"
const dataFilePath string = "puzzles/day10/data.txt"


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

func stringToNumber(number string) int {
	if number == "." {
		return -1
	}
	int_number, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return int_number
}

func calcScore(lines []string, i int, j int) int {
	result := 0
	current_number := 0
	stack := [][2]int{{i, j}}
	stack_cnt := 1
	last_direction := -1
	rollback := false
	ends := [][2]int{}
	for ;len(stack) > 0; {
		if !rollback && (i-1) >= 0 && stringToNumber(lines[i-1][j:j+1]) == (current_number + 1) {
			stack = append(stack, [2]int{i-1, j})
			stack_cnt += 1
			i = i - 1
			current_number += 1
		} else if (rollback && (last_direction < 1) || !rollback) && (j-1) >= 0 && stringToNumber(lines[i][j-1:j]) == (current_number + 1) {
			stack = append(stack, [2]int{i, j-1})
			stack_cnt += 1
			j = j - 1
			current_number += 1
			last_direction = 1
		} else if (rollback && (last_direction < 2) || !rollback) && (i+1) < len(lines) && stringToNumber(lines[i+1][j:j+1]) == (current_number + 1) {
			stack = append(stack, [2]int{i + 1, j})
			stack_cnt += 1
			i = i + 1
			current_number += 1
		} else if (rollback && (last_direction < 3) || !rollback) && (j+1) < len(lines[i]) && stringToNumber(lines[i][j+1:j+2]) == (current_number + 1) {
			stack = append(stack, [2]int{i, j + 1})
			stack_cnt += 1
			j = j + 1
			current_number += 1
			last_direction = 3
		} else if stringToNumber(lines[i][j:j+1]) == 9 {
			if !slices.Contains(ends, [2]int{i, j}) {
				ends = append(ends, [2]int{i, j})
				result += 1
			}
			current_number -= 1
			stack = stack[:stack_cnt - 1]
			stack_cnt -= 1
			rollback = true
			if (i - stack[stack_cnt - 1][0]) == -1 {
				last_direction = 0
			} else if (i - stack[stack_cnt - 1][0]) == 1 {
				last_direction = 2
			} else if (j - stack[stack_cnt - 1][1]) == -1 {
				last_direction = 1
			} else if (j - stack[stack_cnt - 1][1]) == 1 {
				last_direction = 3
			}
			i = stack[stack_cnt - 1][0]
			j = stack[stack_cnt - 1][1]
			continue
		} else {
			current_number -= 1
			stack = stack[:stack_cnt - 1]
			stack_cnt -= 1
			if stack_cnt == 0 {
				break
			}
			rollback = true
			if (i - stack[stack_cnt - 1][0]) == -1 {
				last_direction = 0
			} else if (i - stack[stack_cnt - 1][0]) == 1 {
				last_direction = 2
			} else if (j - stack[stack_cnt - 1][1]) == -1 {
				last_direction = 1
			} else if (j - stack[stack_cnt - 1][1]) == 1 {
				last_direction = 3
			}
			i = stack[stack_cnt - 1][0]
			j = stack[stack_cnt - 1][1]
			continue
		}
		rollback = false
	}
	return result
}

func calcResult(filePath string) int {
	result := 0
	lines := strings.Split(getFileContent(filePath), "\n")
	lines_cnt := len(lines)
	columns_cnt := len(lines[0])
	for i := 0; i < lines_cnt; i++ {
		for j := 0; j < columns_cnt; j++ {
			if lines[i][j:j+1] == "0" {
				result += calcScore(lines, i, j)
			}
		}
	}
	return result
}

func main() {
	example1_result := calcResult(example1FilePath)
	expected1 := 4
	if example1_result != expected1 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected1, example1_result))
	}
	fmt.Printf("Example 1 result: %d.\n", example1_result)
	example2_result := calcResult(example2FilePath)
	expected2 := 3
	if example2_result != expected2 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected2, example2_result))
	}
	fmt.Printf("Example 2 result: %d.\n", example2_result)
	example3_result := calcResult(example3FilePath)
	expected3 := 36
	if example3_result != expected3 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected3, example3_result))
	}
	fmt.Printf("Example 3 result: %d.\n", example3_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"slices"
)


const example1FilePath string = "puzzles/day12/example1.txt"
const example2FilePath string = "puzzles/day12/example2.txt"
const example3FilePath string = "puzzles/day12/example3.txt"
const dataFilePath string = "puzzles/day12/data.txt"


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

func calcResult(filePath string) int {
	result := 0
	lines := strings.Split(getFileContent(filePath), "\n")
	lines_cnt := len(lines)
	columns_cnt := len(lines[0])
	fence := [][2]int{}
	fence_items := 0
	for start_i := 0; start_i < lines_cnt; start_i++ {
		for start_j := 0; start_j < lines_cnt; start_j++ {
			area, p := 1, 0
			letter := lines[start_i][start_j]
			if slices.Contains(fence, [2]int{start_i, start_j}) {
				continue
			}
			fence = append(fence, [2]int{start_i, start_j})
			start := fence_items
			fence_items += 1
			for k := start; k < fence_items; k++ {
				current_item := fence[k]
				i := current_item[0]
				j := current_item[1]
				if (i - 1) < 0 {
					p += 1
				} else if lines[i-1][j] == letter && !slices.Contains(fence, [2]int{i-1, j}) {
					area += 1
					fence = append(fence, [2]int{i-1, j})
					fence_items += 1
				} else if lines[i-1][j] != letter {
					p += 1
				}
				if (j - 1) < 0 {
					p += 1
				} else if lines[i][j-1] == letter && !slices.Contains(fence, [2]int{i, j-1}) {
					area += 1
					fence = append(fence, [2]int{i, j-1})
					fence_items += 1
				} else if lines[i][j-1] != letter {
					p += 1
				}
				if (i + 1) == lines_cnt {
					p += 1
				} else if lines[i+1][j] == letter && !slices.Contains(fence, [2]int{i+1, j}) {
					area += 1
					fence = append(fence, [2]int{i+1, j})
					fence_items += 1
				} else if lines[i+1][j] != letter {
					p += 1
				}
				if (j + 1) == columns_cnt {
					p += 1
				} else if lines[i][j+1] == letter && !slices.Contains(fence, [2]int{i, j+1}) {
					area += 1
					fence = append(fence, [2]int{i, j+1})
					fence_items += 1
				} else if lines[i][j+1] != letter {
					p += 1
				}
			}
			result += area * p
		}
	}
	return result
}

func main() {
	example1_result := calcResult(example1FilePath)
	expected1 := 140
	if example1_result != expected1 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected1, example1_result))
	}
	fmt.Printf("Example 1 result: %d.\n", example1_result)
	example2_result := calcResult(example2FilePath)
	expected2 := 772
	if example2_result != expected2 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected2, example2_result))
	}
	fmt.Printf("Example 2 result: %d.\n", example2_result)
	example3_result := calcResult(example3FilePath)
	expected3 := 1930
	if example3_result != expected3 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected3, example3_result))
	}
	fmt.Printf("Example 3 result: %d.\n", example3_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

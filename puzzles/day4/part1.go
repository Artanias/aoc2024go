package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)


const exampleFilePath string = "puzzles/day4/example.txt"
const example2FilePath string = "puzzles/day4/example2.txt"
const dataFilePath string = "puzzles/day4/data.txt"
const expected string = "XMAS"


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


func checkAndAdd(word string, result *int) {
	if word == expected {
		*result += 1
	}
}

func calcXMASCnt(filePath string) int {
	result := 0
	lines := strings.Split(getFileContent(filePath), "\n")
	rows := len(lines)
	columns := len(lines[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if lines[i][j:j+1] != "X" {
				continue
			}
			if ((j + 4) <= columns) {
				checkAndAdd(lines[i][j:j + 4], &result)
			}
			if ((j - 3) >= 0) {
				left := lines[i][j:j+1] + lines[i][j-1:j] + lines[i][j-2:j-1] + lines[i][j-3:j-2]
				checkAndAdd(left, &result)
			}
			if ((i + 4) <= rows) {
				down := lines[i][j:j+1] + lines[i + 1][j:j+1] + lines[i + 2][j:j+1] + lines[i + 3][j:j+1]
				checkAndAdd(down, &result)
			}
			if ((i - 3) >= 0) {
				up := lines[i][j:j+1] + lines[i - 1][j:j+1] + lines[i - 2][j:j+1] + lines[i - 3][j:j+1]
				checkAndAdd(up, &result)
			}
			if ((j + 4) <= columns) {
				if ((i + 4) <= rows) {
					diag_r_down := lines[i][j:j+1] + lines[i + 1][j+1:j+2] + lines[i + 2][j+2:j+3] + lines[i + 3][j+3:j+4]
					checkAndAdd(diag_r_down, &result)
				}
				if ((i - 3) >= 0) {
					diag_r_up := lines[i][j:j+1] + lines[i - 1][j+1:j+2] + lines[i - 2][j+2:j+3] + lines[i - 3][j+3:j+4]
					checkAndAdd(diag_r_up, &result)
				}
			}
			if ((j - 3) >= 0) {
				if ((i + 4) <= rows) {
					diag_l_down := lines[i][j:j+1] + lines[i + 1][j-1:j] + lines[i + 2][j-2:j-1] + lines[i + 3][j-3:j-2]
					checkAndAdd(diag_l_down, &result)
				}
				if ((i - 3) >= 0) {
					diag_l_up := lines[i][j:j+1] + lines[i - 1][j-1:j] + lines[i - 2][j-2:j-1] + lines[i - 3][j-3:j-2]
					checkAndAdd(diag_l_up, &result)
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Printf("Example result: %d\n", calcXMASCnt(exampleFilePath))
	fmt.Printf("Example 2 result: %d\n", calcXMASCnt(example2FilePath))
	fmt.Printf("Data result: %d\n", calcXMASCnt(dataFilePath))
}

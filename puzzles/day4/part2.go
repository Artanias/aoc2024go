package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)


const example3FilePath string = "puzzles/day4/example3.txt"
const example4FilePath string = "puzzles/day4/example4.txt"
const dataFilePath string = "puzzles/day4/data.txt"
const expected string = "MAS"


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

func calcXCnt(filePath string) int {
	result := 0
	lines := strings.Split(getFileContent(filePath), "\n")
	rows := len(lines)
	columns := len(lines[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if lines[i][j:j+1] != "A" {
				continue
			}
			mas_cnt := 0
			if ((j + 2) <= columns && (j - 1) >= 0) {
				if ((i + 2) <= rows && (i - 1) >= 0) {
					diag_r_down := lines[i-1][j-1:j] + lines[i][j:j+1] + lines[i + 1][j+1:j+2]
					checkAndAdd(diag_r_down, &mas_cnt)
				}
				if ((i - 1) >= 0 && (i + 1) < rows) {
					diag_r_up := lines[i+1][j-1:j] + lines[i][j:j+1] + lines[i - 1][j+1:j+2]
					checkAndAdd(diag_r_up, &mas_cnt)
				}
			}
			if ((j - 1) >= 0 && (j + 1) < columns) {
				if ((i + 2) <= rows && (i - 1) >= 0) {
					diag_l_down := lines[i - 1][j+1:j+2] + lines[i][j:j+1] + lines[i + 1][j-1:j]
					checkAndAdd(diag_l_down, &mas_cnt)
				}
				if ((i - 1) >= 0 && (i + 1) < rows) {
					diag_l_up := lines[i+1][j+1:j+2] + lines[i][j:j+1] + lines[i - 1][j-1:j]
					checkAndAdd(diag_l_up, &mas_cnt)
				}
			}
			if mas_cnt == 2 {
				result += 1
			}
		}
	}
	return result
}

func main() {
	fmt.Printf("Example 3 result: %d\n", calcXCnt(example3FilePath))
	fmt.Printf("Example 4 result: %d\n", calcXCnt(example4FilePath))
	fmt.Printf("Data result: %d\n", calcXCnt(dataFilePath))
}

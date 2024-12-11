package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"slices"
)


const exampleFilePath string = "puzzles/day7/example.txt"
const dataFilePath string = "puzzles/day7/data.txt"


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
	for i := 0; i < lines_cnt; i++ {
		before, after, found := strings.Cut(lines[i], ":")
		if !found {
			panic("Invalid data.")
		}
		expected := stringToNumber(before)
		numbers := strings.Split(after, " ")
		int_numbers := []int{}
		cnt_numbers := len(numbers)
		for j := 1; j < cnt_numbers; j++ {
			int_numbers = append(int_numbers, stringToNumber(numbers[j]))
		}
		results := []int{int_numbers[0]}
		results_cnt := 1
		cnt_numbers = len(int_numbers)
		for j := 1; j < cnt_numbers; j++ {
			number := int_numbers[j]
			for k := 0; k < results_cnt; k++ {
				results = append(results, results[k] * number)
				results[k] = results[k] + number
			}
			results_cnt *= 2
		}
		if slices.Contains(results, expected) {
			result += expected
		}
	}
	return result
}

func main() {
	example_result := calcResult(exampleFilePath)
	expected := 3749
	if example_result != expected {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected, example_result))
	}
	fmt.Printf("Example result: %d.\n", example_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

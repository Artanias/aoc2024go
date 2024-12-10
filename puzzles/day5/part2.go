package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"slices"
)


const exampleFilePath string = "puzzles/day5/example.txt"
const dataFilePath string = "puzzles/day5/data.txt"


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

func correctPages(pages []int, rules_map map[int][]int, cnt_pages int) {
	correct := false
	for ;!correct; {
		correct = true
		prev := -1
		for i := 0; i < cnt_pages; i++ {
			if slices.Contains(rules_map[pages[i]], prev) {
				tmp := pages[i]
				pages[i] = prev
				pages[i - 1] = tmp
				prev = tmp
				correct = false
			}
			prev = pages[i]
		}
	}
}

func calcResult(filePath string) int {
	result := 0
	rules, page_sequences, found := strings.Cut(getFileContent(filePath), "\n\n")
	if !found {
		panic("Invalid file.")
	}
	lines := strings.Split(rules, "\n")
	rows := len(lines)
	rules_map := make(map[int][]int)
	for i := 0; i < rows; i++ {
		left, right, found := strings.Cut(lines[i], "|")
		if !found {
			panic("Invalid sequence.")
		}
		left_int := stringToNumber(left)
		right_int := stringToNumber(right)
		rules_map[left_int] = append(rules_map[left_int], right_int)
	}
	lines = strings.Split(page_sequences, "\n")
	rows = len(lines)
	for i := 0; i < rows; i++ {
		pages := strings.Split(lines[i], ",")
		pages_int := []int{}
		cnt_pages := len(pages)
		for j := 0; j < cnt_pages; j++ {
			pages_int = append(pages_int, stringToNumber(pages[j]))
		}
		prev := -1
		correct := true
		for j := 0; j < cnt_pages; j++ {
			if slices.Contains(rules_map[pages_int[j]], prev) {
				correct = false
				correctPages(pages_int, rules_map, cnt_pages)
				break
			}
			prev = pages_int[j]
		}
		if !correct {
			result += pages_int[cnt_pages / 2]
		}
	}
	return result
}

func main() {
	example_result := calcResult(exampleFilePath)
	expected := 123
	if example_result != expected {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected, example_result))
	}
	fmt.Printf("Example result: %d.\n", example_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

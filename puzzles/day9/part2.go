package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)


const exampleFilePath string = "puzzles/day9/example.txt"
const dataFilePath string = "puzzles/day9/data.txt"


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

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func stringToNumber(number string) int {
	int_number, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return int_number
}

func calcNewRanges(disk_map []string) [][2]int {
	dot_ranges := [][2]int{}
	start := -1
	end := 0
	for i := 0; i < len(disk_map); i++ {
		if start == -1 && disk_map[i] == "." {
			start = i
			end = start + 1
		} else if disk_map[i] == "." {
			end += 1
		} else if start != -1 {
			dot_ranges = append(dot_ranges, [2]int{start, end})
			start = -1
		}
	}
	return dot_ranges
}

func calcResult(filePath string) int {
	result := 0
	line := getFileContent(filePath)
	id := 0
	disk_map := []string{}
	index := 0
	num_ranges := [][2]int{}
	dot_ranges := [][2]int{}
	for i := 0; i < len(line); i++ {
		number := stringToNumber(line[i:i+1])
		string_id := strconv.Itoa(id)
		if i % 2 == 0 {
			num_ranges = append(num_ranges, [2]int{index, index + number})
			for j := 0; j < number; j++ {
				disk_map = append(disk_map, string_id)
				index += 1
			}
			id += 1
		} else {
			dot_ranges = append(dot_ranges, [2]int{index, index + number})
			for j := 0; j < number; j++ {
				disk_map = append(disk_map, ".")
				index += 1
			}
		}
	}
	for i := len(num_ranges) - 1; i >= 0; i-- {
		for j := 0; j < len(dot_ranges); j++ {
			free_space := dot_ranges[j][1] - dot_ranges[j][0]
			busy_space := num_ranges[i][1] - num_ranges[i][0]
			if busy_space > free_space {
				continue
			}
			if dot_ranges[j][0] > num_ranges[i][0] {
				continue
			}
			for k := 0; k < (num_ranges[i][1] - num_ranges[i][0]); k++ {
				disk_map[dot_ranges[j][0]+k] = disk_map[num_ranges[i][0]+k]
				disk_map[num_ranges[i][0]+k] = "."
			}
			dot_ranges = calcNewRanges(disk_map)
			break
		}
	}
	for i := 0; i < len(disk_map); i++ {
		if disk_map[i] == "." {
			continue
		}
		result += (i * stringToNumber(disk_map[i]))
	}
	return result
}

func main() {
	example_result := calcResult(exampleFilePath)
	expected := 2858
	if example_result != expected {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected, example_result))
	}
	fmt.Printf("Example result: %d.\n", example_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

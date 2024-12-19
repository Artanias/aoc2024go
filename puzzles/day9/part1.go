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

func calcResult(filePath string) int {
	result := 0
	line := getFileContent(filePath)
	id := 0
	disk_map := []string{}
	for i := 0; i < len(line); i++ {
		number := stringToNumber(line[i:i+1])
		string_id := strconv.Itoa(id)
		if i % 2 == 0 {
			for j := 0; j < number; j++ {
				disk_map = append(disk_map, string_id)
			}
			id += 1
		} else {
			for j := 0; j < number; j++ {
				disk_map = append(disk_map, ".")
			}
		}
	}
	i := 0
	j := len(disk_map) - 1
	for {
		if i == j {
			break
		}
		if disk_map[i] != "." {
			i++
			continue
		}
		if disk_map[j] == "." {
			j--
			continue
		}
		disk_map[i] = disk_map[j]
		disk_map[j] = "."
	}
	for i := 0; i < len(disk_map); i++ {
		if disk_map[i] == "." {
			break
		}
		result += (i * stringToNumber(disk_map[i]))
	}
	return result
}

func main() {
	example_result := calcResult(exampleFilePath)
	expected := 1928
	if example_result != expected {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected, example_result))
	}
	fmt.Printf("Example result: %d.\n", example_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)


const exampleFilePath string = "puzzles/day3/example.txt"
const dataFilePath string = "puzzles/day3/data.txt"


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

func calcMulResults(filePath string) int {
	totalMulResult := 0
	content := getFileContent(filePath)
	lines := strings.Split(content, "\n")
	pattern := regexp.MustCompile(`mul\((?P<first_number>\d+),(?P<second_number>\d+)\)`)
	//names := FindAllStringSubmatch
	for i := 0; i < len(lines); i++ {
		matches := pattern.FindAllStringSubmatch(lines[i], -1)
		for j := 0; j < len(matches); j++ {
			first_number := stringToNumber(matches[j][1])
			second_number := stringToNumber(matches[j][2])
			totalMulResult += first_number * second_number
		}
	}
	return totalMulResult
}

func main() {
	fmt.Printf("Example result: %d\n", calcMulResults(exampleFilePath))
	fmt.Printf("Data result: %d\n", calcMulResults(dataFilePath))
}

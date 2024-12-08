package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)


const exampleFilePath string = "puzzles/day2/example.txt"
const dataFilePath string = "puzzles/day2/data.txt"


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

func isReportSafe(string_numbers []string) bool {
	safe := true
	prev_number := stringToNumber(string_numbers[0])
	prev_diff := 0
	for j := 1; j < len(string_numbers); j++ {
		int_number := stringToNumber(string_numbers[j])
		diff := prev_number - int_number
		if diff == 0 {
			safe = false
			break
		} else if diff > 0 && prev_diff < 0 {
			safe = false
			break
		} else if diff < 0 && prev_diff > 0 {
			safe = false
			break
		} else if diff > 3 || diff < -3 {
			safe = false
			break
		}
		prev_number = int_number
		prev_diff = diff
	}
	return safe
}

func calcSafeReports(filePath string) int {
	totalSafeReports := 0
	content := getFileContent(filePath)
	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		string_numbers := strings.Split(lines[i], " ")
		if isReportSafe(string_numbers) {
			totalSafeReports += 1
			continue
		}
		if isReportSafe(string_numbers[1:]) {
			totalSafeReports += 1
			continue
		}
		if isReportSafe(string_numbers[:len(string_numbers) - 1]) {
			totalSafeReports += 1
			continue
		}
		for j := 1; j < len(string_numbers) - 1; j++ {
			dst_string_numbers := make([]string, len(string_numbers))
			copy(dst_string_numbers, string_numbers)
			if isReportSafe(append(dst_string_numbers[0:j], dst_string_numbers[j + 1:]...)) {
				totalSafeReports += 1
				break
			}
		}
	}
	return totalSafeReports
}

func main() {
	fmt.Printf("Example total safe reports: %d\n", calcSafeReports(exampleFilePath))
	fmt.Printf("Data total safe reports: %d\n", calcSafeReports(dataFilePath))
}

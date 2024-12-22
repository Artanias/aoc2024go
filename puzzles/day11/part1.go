package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)


const example1FilePath string = "puzzles/day11/example.txt"
const dataFilePath string = "puzzles/day11/data.txt"


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
	int_numbers := []int{}
	numbers := strings.Split(getFileContent(filePath), " ")
	numbers_cnt := len(numbers)
	for i := 0; i < numbers_cnt; i++ {
		int_numbers = append(int_numbers, stringToNumber(numbers[i]))
	}
	blinks_cnt := 25
	for i := 0; i < blinks_cnt; i++ {
		for j := 0; j < numbers_cnt; j++ {
			string_number := strconv.Itoa(int_numbers[j])
			if int_numbers[j] == 0 {
				int_numbers[j] += 1
			} else if len(string_number) % 2 == 0 {
				first_number := stringToNumber(string_number[:len(string_number) / 2])
				second_number := stringToNumber(string_number[len(string_number) / 2:])
				int_numbers[j] = first_number
				new_int_numbers := []int{}
				for k := 0; k <= j; k++ {
					new_int_numbers = append(new_int_numbers, int_numbers[k])
				}
				new_int_numbers = append(new_int_numbers, second_number)
				new_int_numbers = append(new_int_numbers, int_numbers[j+1:]...)
				int_numbers = new_int_numbers
				numbers_cnt += 1
				j += 1
			} else {
				int_numbers[j] *= 2024
			}
		}
	}
	return numbers_cnt
}

func main() {
	example1_result := calcResult(example1FilePath)
	expected1 := 55312
	if example1_result != expected1 {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected1, example1_result))
	}
	fmt.Printf("Example 1 result: %d.\n", example1_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

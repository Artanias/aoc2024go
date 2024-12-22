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

func calcSmallPart(initial_number int, iterations int) int {
	int_numbers := []int{initial_number}
	numbers_cnt := 1
	for i := 0; i < (iterations); i++ {
		local_numbers_cnt := numbers_cnt
		for j := 0; j < local_numbers_cnt; j++ {
			if int_numbers[j] == 0 {
				int_numbers[j] = 1
				continue
			}
			string_number := strconv.Itoa(int_numbers[j])
			len_string_number := len(string_number)
			if len_string_number % 2 == 0 {
				len_string_number_divided_by_2 := len_string_number / 2
				first_number := stringToNumber(string_number[:len_string_number_divided_by_2])
				second_number := stringToNumber(string_number[len_string_number_divided_by_2:])
				int_numbers[j] = first_number
				int_numbers = append(int_numbers, second_number)
				numbers_cnt += 1
			} else {
				int_numbers[j] *= 2024
			}
		}
	}
	return numbers_cnt - 1
}

func calcResult(filePath string) int {
	numbers := strings.Split(getFileContent(filePath), " ")
	initial_numbers_cnt := len(numbers)
	initial_int_numbers := make([]int, initial_numbers_cnt, initial_numbers_cnt)
	for i := 0; i < initial_numbers_cnt; i++ {
		initial_int_numbers[i] = stringToNumber(numbers[i])
	}
	total_numbers_cnt := initial_numbers_cnt
	initial_steps := 40
	blinks_cnt := 75
	additional_iterations := blinks_cnt - initial_steps
	cache := make(map[int]map[int]int)
	for k := 0; k < initial_steps; k++ {
		local_numbers_cnt := total_numbers_cnt
		for j := 0; j < local_numbers_cnt; j++ {
			if initial_int_numbers[j] == 0 {
				initial_int_numbers[j] = 1
				continue
			}
			string_number := strconv.Itoa(initial_int_numbers[j])
			len_string_number := len(string_number)
			if len_string_number % 2 == 0 {
				len_string_number_divided_by_2 := len_string_number / 2
				first_number := stringToNumber(string_number[:len_string_number_divided_by_2])
				second_number := stringToNumber(string_number[len_string_number_divided_by_2:])
				initial_int_numbers[j] = first_number
				initial_int_numbers = append(initial_int_numbers, second_number)
				total_numbers_cnt += 1
				initial_numbers_cnt += 1
			} else {
				initial_int_numbers[j] *= 2024
			}
		}
	}
	fmt.Println("First part done.")
	for k := 0; k < initial_numbers_cnt; k++ {
		_, ok := cache[initial_int_numbers[k]]
		if !ok {
			cache[initial_int_numbers[k]] = map[int]int{}
		}
		_, ok = cache[initial_int_numbers[k]][additional_iterations]
		if !ok {
			cache[initial_int_numbers[k]][additional_iterations] = calcSmallPart(initial_int_numbers[k], additional_iterations)
		}
		total_numbers_cnt += cache[initial_int_numbers[k]][additional_iterations]
	}

	if blinks_cnt == 60 && total_numbers_cnt != 470980814376 {
		panic(fmt.Sprintf("Current: %d, Expected: %d.", total_numbers_cnt, 470980814376))
	}
	return total_numbers_cnt
}

func main() {
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

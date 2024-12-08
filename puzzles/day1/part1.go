package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)


const exampleFilePath string = "puzzles/day1/example.txt"
const dataFilePath string = "puzzles/day1/data.txt"


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


func calcTotalDistance(filePath string) int {
	totalDistance := 0
	content := getFileContent(filePath)
	lines := strings.Split(content, "\n")
	left := []int{}
	right := []int{}
	for i := 0; i < len(lines); i++ {
		before, after, found := strings.Cut(lines[i], "   ")
		if found != true {
			panic("Incorrect data in the file.")
		}
		before_int, err := strconv.Atoi(before)
		if err != nil {
			panic(err)
		}
		after_int, err := strconv.Atoi(after)
		if err != nil {
			panic(err)
		}
		left = append(left, before_int)
		right = append(right, after_int)
	}
	sort.Ints(left)
	sort.Ints(right)
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			totalDistance += left[i] - right[i]
		} else {
			totalDistance += right[i] - left[i]
		}
	}
	return totalDistance
}

func main() {
	fmt.Printf("Example total distance: %d\n", calcTotalDistance(exampleFilePath))
	fmt.Printf("Data total distance: %d\n", calcTotalDistance(dataFilePath))
}

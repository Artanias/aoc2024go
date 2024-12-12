package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"math"
)


const exampleFilePath string = "puzzles/day8/example.txt"
const dataFilePath string = "puzzles/day8/data.txt"


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

func calcResult(filePath string) int {
	result := 0
	lines := strings.Split(getFileContent(filePath), "\n")
	lines_cnt := len(lines)
	columns_cnt := len(lines[0])
	antennas := make(map[string][][2]int)
	for i := 0; i < lines_cnt; i++ {
		for j := 0; j < columns_cnt; j++ {
			symbol := lines[i][j:j+1]
			_, ok := antennas[symbol]
			if !ok && symbol != "." {
				antennas[symbol] = [][2]int{}
			}
			if symbol != "." {
				antennas[symbol] = append(antennas[symbol], [2]int{i, j})
			}
		}
	}
	for _, value := range antennas {
		cnt_elements := len(value)
		for i := 0; i < cnt_elements; i++ {
			left := value[i]
			for j := i + 1; j < cnt_elements; j++ {
				right := value[j]
				xd:= left[0] - right[0]
				xy := left[1] - right[1]
				for ;math.Abs(float64(xd)) < float64(lines_cnt) && math.Abs(float64(xy)) < float64(columns_cnt); {
					antinode1 := [2]int{left[0] + xd, left[1] + xy}
					antinode2 := [2]int{right[0] - xd, right[1] - xy}
					if antinode1[0] < lines_cnt && antinode1[0] >= 0 && antinode1[1] < columns_cnt && antinode1[1] >= 0 {
						lines[antinode1[0]] = replaceAtIndex(lines[antinode1[0]], '#', antinode1[1])
					}
					if antinode2[0] < lines_cnt && antinode2[0] >= 0 && antinode2[1] < columns_cnt && antinode2[1] >= 0 {
						lines[antinode2[0]] = replaceAtIndex(lines[antinode2[0]], '#', antinode2[1])
					}
					xd += (left[0] - right[0])
					xy += (left[1] - right[1])
				}
				xd -= (left[0] - right[0])
				xy -= (left[1] - right[1])
				for ;math.Abs(float64(xd)) < float64(lines_cnt) && math.Abs(float64(xy)) < float64(columns_cnt); {
					antinode1 := [2]int{left[0] + xd, left[1] + xy}
					antinode2 := [2]int{right[0] - xd, right[1] - xy}
					if antinode1[0] < lines_cnt && antinode1[0] >= 0 && antinode1[1] < columns_cnt && antinode1[1] >= 0 {
						lines[antinode1[0]] = replaceAtIndex(lines[antinode1[0]], '#', antinode1[1])
					}
					if antinode2[0] < lines_cnt && antinode2[0] >= 0 && antinode2[1] < columns_cnt && antinode2[1] >= 0 {
						lines[antinode2[0]] = replaceAtIndex(lines[antinode2[0]], '#', antinode2[1])
					}
					xd -= (left[0] - right[0])
					xy -= (left[1] - right[1])
				}
			}
		}
	}
	for i := 0; i < lines_cnt; i++ {
		for j := 0; j < columns_cnt; j++ {
			if lines[i][j:j+1] == "#" {
				result += 1
			}
		}
	}
	return result
}

func main() {
	example_result := calcResult(exampleFilePath)
	expected := 34
	if example_result != expected {
		panic(fmt.Sprintf("Expected: %d. Gotten: %d.", expected, example_result))
	}
	fmt.Printf("Example result: %d.\n", example_result)
	fmt.Printf("Data result: %d.\n", calcResult(dataFilePath))
}

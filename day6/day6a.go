package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := convertToIntSlice(scanner.Text())

	scanner.Scan()
	distances := convertToIntSlice(scanner.Text())

	product := 1

	for idx, time := range times {
		count := 0
		for i := 0; i < time; i++ {
			if (time-i)*i > distances[idx] {
				count++
			}
		}
		if count == 0 {
			break
		}

		product *= count
	}

	fmt.Println(product)
}

func convertToIntSlice(str string) []int {

	colonIndex := strings.Index(str, ":")
	str = str[colonIndex+1:]
	strSlice := strings.Fields(str)
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		num, _ := strconv.Atoi(str)
		intSlice[i] = num
	}
	return intSlice
}

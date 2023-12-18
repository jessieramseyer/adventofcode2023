package day9b

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		diff_array := make([][]int, 0, 0)
		string_numbers := strings.Fields(line)
		numbers := make([]int, 0, 0)
		for _, numStr := range string_numbers {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		diff_array = append(diff_array, numbers)

		diff := 1
		temp := numbers
		done := false
		for !done {
			temp2 := make([]int, 0, 0)
			for i, val := range temp {
				if i == 0 {
					continue
				}
				diff = val - temp[i-1]
				temp2 = append(temp2, diff)
			}
			temp = temp2
			diff_array = append(diff_array, temp)
			done = true
			for _, item := range temp {
				if item != 0 {
					done = false
					break
				}
			}
		}

		// alternating between adding and subtracting the first element in each array will yield the number we want
		for i, group := range diff_array {

			if i%2 == 0 {
				//even = add
				sum += group[0]
			} else {
				//odd = subtract
				sum -= group[0]
			}
		}
	}

	fmt.Println(sum)

}

package day3b

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Run() {

	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	contents := make([]string, 0, 0)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		contents = append(contents, line)
	}
	numRows := len(contents)
	numCols := len(contents[0])

	sum := 0

	moves := [][]int{
		{-1, 0},  //up
		{1, 0},   //down
		{0, -1},  //left
		{0, 1},   //right
		{-1, -1}, //up left
		{-1, 1},  // up right
		{1, -1},  // down left
		{1, 1},   // down right
	}

	gearsMap := make(map[string][]int)

	for rowidx, row := range contents {

		gearId := ""
		current_num := 0
		for colidx, digit := range row {

			//if character is not a number move on
			if !unicode.IsNumber(digit) && gearId == "" {
				current_num = 0
				continue
				// if first non number character after part number, add the number is adjacent to a gear
			} else if !unicode.IsNumber(digit) && gearId != "" {
				gearsMap[gearId] = append(gearsMap[gearId], current_num)
				gearId = ""
				current_num = 0
				continue
			}

			//try to verify if it's a gear only if no previous numbers have determined that
			if gearId == "" {
				for _, move := range moves {

					row_check := rowidx + move[0]
					col_check := colidx + move[1]
					// if outside of bounds of array then move on
					if row_check < 0 || row_check > numRows-1 || col_check < 0 || col_check > numCols-1 {
						continue
					}
					//if character is * save the index as a string so that we can use it as a key in the gearsMap
					if string(contents[row_check][col_check]) == "*" {
						gearId = fmt.Sprintf("%v, %v", row_check, col_check)
						break
					}
				}
			}
			//converting rune to int
			current_num = current_num*10 + int(digit-48)
		}
		//case where number is at end of row
		if gearId != "" {
			gearsMap[gearId] = append(gearsMap[gearId], current_num)
		}
	}

	for _, item := range gearsMap {
		if len(item) == 2 {
			sum += item[0] * item[1]
		}
	}

	fmt.Println("Sum:", sum)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

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

	for rowidx, row := range contents {

		part_num := false
		current_num := 0
		for colidx, digit := range row {

			//if character is not a number move on
			if !unicode.IsNumber(digit) && !part_num {
				current_num = 0
				continue
				// if first non number character after part number, add it to the sum and reset part_num and current_num
			} else if !unicode.IsNumber(digit) && part_num {
				sum += current_num
				part_num = false
				current_num = 0
				continue
			}

			//try to verify if it's a part number only if no previous numbers have determined that
			if !part_num {
				for _, move := range moves {

					row_check := rowidx + move[0]
					col_check := colidx + move[1]
					// if outside of bounds of array then move on
					if row_check < 0 || row_check > numRows-1 || col_check < 0 || col_check > numCols-1 {
						continue
					}
					// if character is not a symbol move on
					if unicode.IsNumber(rune(contents[row_check][col_check])) || string(contents[row_check][col_check]) == "." {
						continue
					}
					//if symbol found
					part_num = true
					break
				}
			}
			//converting rune to int
			current_num = current_num*10 + int(digit-48)
		}
		//case where number is at end of row
		if part_num {
			sum += current_num
		}
	}

	fmt.Println("Sum:", sum)

}

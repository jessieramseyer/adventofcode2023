package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	sectionMap := make(map[int][][]int)
	section := 0

	// first line
	scanner.Scan()
	line := scanner.Text()
	colonIndex := strings.Index(line, ":")
	line = line[colonIndex+1:]
	string_seeds := strings.Fields(line)
	seeds := convertToIntSlice(string_seeds)

	for scanner.Scan() {
		line := scanner.Text()
		//marks that we are at new section
		if len(line) == 0 {
			continue
		}
		if unicode.IsLetter(rune(line[0])) {
			section++
		} else if unicode.IsNumber(rune(line[0])) {
			vals := strings.Fields(line)
			nums := convertToIntSlice(vals)
			sectionMap[section] = append(sectionMap[section], nums)
		}
	}

	location := 100000000000
	for _, seed := range seeds {
		current := seed
		for i := 1; i <= len(sectionMap); i++ {
			for _, vals := range sectionMap[i] {
				// destination, start, range
				if current >= vals[1] && current <= vals[1]+vals[2]-1 {
					current = current - vals[1] + vals[0]
					break
				}
			}
		}
		if current < location {
			location = current
		}
	}
	fmt.Println(location)
}

func convertToIntSlice(strSlice []string) []int {
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		num, _ := strconv.Atoi(str)
		intSlice[i] = num
	}
	return intSlice
}

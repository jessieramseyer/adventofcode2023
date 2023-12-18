package day8a

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Run() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	RightMap := make(map[string]string)
	LeftMap := make(map[string]string)
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define a regular expression to extract AAA, BBB, and CCC
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	//get first line
	scanner.Scan()
	instructions := scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		// Find submatches that match the regex pattern
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			value := matches[1]
			left := matches[2]
			right := matches[3]
			RightMap[value] = right
			LeftMap[value] = left
		}
	}

	next := "AAA"
	steps := 0

	for next != "ZZZ" {
		for _, char := range instructions {
			if char == 'R' {
				next = RightMap[next]
			} else if char == 'L' {
				next = LeftMap[next]
			}
			steps++
		}
	}
	fmt.Println(steps)

}

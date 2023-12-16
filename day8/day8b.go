package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
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

	endMap := make(map[string]string)

	for value, _ := range RightMap {
		next := value
		for _, char := range instructions {
			if char == 'R' {
				next = RightMap[next]
			} else if char == 'L' {
				next = LeftMap[next]
			}
		}
		endMap[value] = next
	}

	queue1 := make([]string, 0, 6)
	for item, _ := range RightMap {
		if item[2] == 'A' {
			queue1 = append(queue1, item)
		}
	}

	LCMs := make([]int, 0, 6)
	// determine minimum number of steps to go from each starting value to a value with a Z
	for _, next := range queue1 {
		steps := 0
		for next[2] != 'Z' {
			next = endMap[next]
			steps++
		}
		LCMs = append(LCMs, steps*len(instructions))
	}

	// find lowest common multiple
	lcm := LCMs[0]
	for i := 1; i < len(LCMs); i++ {
		lcm = lcm * LCMs[i] / gcd(lcm, LCMs[i])
	}

	fmt.Println(lcm)

}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

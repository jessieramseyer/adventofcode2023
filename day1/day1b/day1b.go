package day1b

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Run() {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	re1 := regexp.MustCompile(`^.*?(one|two|three|four|five|six|seven|eight|nine|\d)`)
	re2 := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|\d).*$`)
	sum := 0

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Find all matches in the text
		leftgroups := re1.FindAllStringSubmatch(line, -1)
		rightgroups := re2.FindAllStringSubmatch(line, -1)
		num1, _ := numMap[leftgroups[0][1]]
		var num2 int
		if rightgroups[0][1] == "" {
			num2 = num1
		} else {
			num2, _ = numMap[rightgroups[0][1]]
		}

		sum += num1*10 + num2
	}
	fmt.Println("Sum:", sum)
}

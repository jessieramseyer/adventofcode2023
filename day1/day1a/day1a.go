package day1a

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Run() {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`^\D*(\d).*?(\d)?\D*$`)
	sum := 0

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Find all matches in the text
		groups := re.FindAllStringSubmatch(line, -1)
		num1, _ := strconv.Atoi(groups[0][1])
		var num2 int
		if groups[0][2] == "" {
			num2 = num1
		} else {
			num2, _ = strconv.Atoi(groups[0][2])
		}
		sum += num1*10 + num2
	}
	fmt.Println("Sum:", sum)
}

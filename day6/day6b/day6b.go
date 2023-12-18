package day6b

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := convertToInt(scanner.Text())
	scanner.Scan()
	distance := convertToInt(scanner.Text())

	count := 0
	for i := 0; i < time; i++ {
		if (time-i)*i > distance {
			count++
		}
	}

	fmt.Println(count)
}

func convertToInt(str string) int {
	colonIndex := strings.Index(str, ":")
	str = str[colonIndex+1:]
	str = strings.ReplaceAll(str, " ", "")
	integer, _ := strconv.Atoi(str)

	return integer
}

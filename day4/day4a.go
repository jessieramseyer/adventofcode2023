package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		count := -1.0

		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]

		cardMap := make(map[string]string)

		cardData := strings.Split(line, "|")
		// store winner values to compare to
		winners := strings.Split(cardData[0], " ")
		for _, winner := range winners {
			winner = strings.TrimSpace(winner)
			if winner != "" {
				cardMap[winner] = winner
			}
		}
		// count matches
		numbers := strings.Split(cardData[1], " ")
		for _, number := range numbers {
			number = strings.TrimSpace(number)
			if _, exists := cardMap[number]; exists {
				count += 1
			}
		}
		sum += int(math.Pow(2, count))
	}

	fmt.Println("Sum:", sum)
}

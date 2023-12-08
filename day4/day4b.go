package main

import (
	"bufio"
	"fmt"
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

	fullCardMap := make(map[int]int)
	card_num := 1

	for scanner.Scan() {
		line := scanner.Text()
		count := 0

		fullCardMap[card_num] += 1

		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]

		matchMap := make(map[string]string)

		cardData := strings.Split(line, "|")
		winners := strings.Split(cardData[0], " ")
		// store winner values to compare to
		for _, winner := range winners {
			winner = strings.TrimSpace(winner)
			if winner != "" {
				matchMap[winner] = winner
			}
		}

		numbers := strings.Split(cardData[1], " ")
		// count matches
		for _, number := range numbers {
			number = strings.TrimSpace(number)
			if _, exists := matchMap[number]; exists {
				count += 1
			}
		}
		// increment the number of following cards based on match count
		for i := 1; i <= count; i++ {
			fullCardMap[card_num+i] += fullCardMap[card_num]
		}

		card_num++
	}

	for _, count := range fullCardMap {
		sum += count
	}

	fmt.Println("Sum:", sum)
}

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

	sum := 0

	// Read file line by line
	scanner := bufio.NewScanner(file)
	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()

		seen := make(map[string]int)

		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]

		gameData := strings.Split(line, ";")
		for _, data := range gameData {
			// Splitting each set by commas to get individual colors
			colors := strings.Split(data, ",")
			for _, color := range colors {
				// Trimming spaces from the color string
				color = strings.TrimSpace(color)
				// Extracting the color name and count
				var count int
				var name string
				// Checking if there's a number and color
				fmt.Sscanf(color, "%d %s", &count, &name)
				if value, exists := seen[name]; exists {
					if count > value {
						seen[name] = count
					}
				} else {
					seen[name] = count
				}
			}

		}

		sum += seen["red"] * seen["green"] * seen["blue"]
	}

	fmt.Println("Sum:", sum)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

		re := regexp.MustCompile(`Game (\d+)`)
		groups := re.FindAllStringSubmatch(line, -1)
		gameNum, _ := strconv.Atoi(groups[0][1])

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

		if seen["red"] <= 12 && seen["green"] <= 13 && seen["blue"] <= 14 {
			sum += gameNum
		}
	}

	fmt.Println("Sum:", sum)
}

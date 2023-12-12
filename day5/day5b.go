package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// first line
	scanner.Scan()
	line := scanner.Text()
	colonIndex := strings.Index(line, ":")
	line = line[colonIndex+1:]
	string_seeds := strings.Fields(line)
	seed_ranges := convertToIntSlice(string_seeds)
	seeds := make([][]int, 0)

	for j := 0; j < len(seed_ranges); j += 2 {
		seed := []int{seed_ranges[j], seed_ranges[j+1] + seed_ranges[j] - 1}
		seeds = append(seeds, seed)
	}

	for scanner.Scan() {
		line := scanner.Text()

		//marks that we are at new section
		if len(line) == 0 || unicode.IsLetter(rune(line[0])) {
			continue
		}
		mappings := [][]int{}
		newranges := [][]int{}
		for unicode.IsNumber(rune(line[0])) {
			vals := strings.Fields(line)
			nums := convertToIntSlice(vals)
			mappings = append(mappings, nums)
			if scanner.Scan() {
				line = scanner.Text()
				if line == "" {
					break
				}
			} else {
				break
			}
		}
		//sort map by start element
		sort.Slice(mappings, func(i, j int) bool {
			return mappings[i][1] < mappings[j][1]
		})

		for _, seed := range seeds {
			current := seed
			end := 0
			for _, section := range mappings {

				start := section[1]
				end = section[1] + section[2]
				offset := -section[1] + section[0]
				// destination, start, range
				if current[0] < start {
					if current[1] < start {
						newranges = append(newranges, current)
						break
					}
					newranges = append(newranges, []int{current[0], start - 1})
					if current[1] < end {
						newranges = append(newranges, []int{start + offset, current[1] + offset})
						break
					}
					if current[1] >= end {
						newranges = append(newranges, []int{start + offset, end - 1 + offset})
						current[0] = end
					}
				} else if current[0] < end {
					if current[1] < end {
						newranges = append(newranges, []int{current[0] + offset, current[1] + offset})
						break
					}
					if current[1] >= end {
						newranges = append(newranges, []int{current[0] + offset, end - 1 + offset})
						current[0] = end
					}
				}
			}
			if current[0] >= end {
				newranges = append(newranges, []int{current[0], current[1]})
			}
		}

		seeds = newranges

	}

	// Sorting 'seeds' based on the first element of each sublist
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i][0] < seeds[j][0]
	})

	// Accessing the minimum value after sorting
	minValue := seeds[0][0]
	fmt.Println(minValue) // This will print the minimum value based on the first element of each sublist
}

func convertToIntSlice(strSlice []string) []int {
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		num, _ := strconv.Atoi(str)
		intSlice[i] = num
	}
	return intSlice
}

package day7a

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandBidPair struct {
	Hand string
	Bid  int
}

var rankWeights = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func Run() {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	FiveofaKind := make([]HandBidPair, 0, 0)
	FourofaKind := make([]HandBidPair, 0, 0)
	FullHouse := make([]HandBidPair, 0, 0)
	ThreeofaKind := make([]HandBidPair, 0, 0)
	TwoPair := make([]HandBidPair, 0, 0)
	OnePair := make([]HandBidPair, 0, 0)
	HighCard := make([]HandBidPair, 0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		hand := fields[0]
		bid, _ := strconv.Atoi(fields[1])

		pair := HandBidPair{hand, bid}

		charCount := make(map[rune]int)

		// Count occurrences of each character
		for _, char := range hand {
			charCount[char]++
		}

		maxCount := 0

		// Find the maximum count of a character
		for _, count := range charCount {
			if count > maxCount {
				maxCount = count
			}
		}

		//determine hand type
		if maxCount == 5 {
			FiveofaKind = append(FiveofaKind, pair)
		} else if maxCount == 4 {
			FourofaKind = append(FourofaKind, pair)
		} else if maxCount == 3 {
			if len(charCount) == 2 {
				//Full House
				FullHouse = append(FullHouse, pair)
			} else {
				ThreeofaKind = append(ThreeofaKind, pair)
			}
		} else if maxCount == 2 {
			if len(charCount) == 3 {
				//Two Pair
				TwoPair = append(TwoPair, pair)
			} else {
				//One Pair
				OnePair = append(OnePair, pair)
			}
		} else {
			HighCard = append(HighCard, pair)
		}
	}
	rank := 1
	sum := 0
	AllHands := [][]HandBidPair{HighCard, OnePair, TwoPair, ThreeofaKind, FullHouse, FourofaKind, FiveofaKind}
	for _, group := range AllHands {
		group = customSort(group)
		for _, pair := range group {
			sum += rank * pair.Bid
			rank++
		}
	}

	fmt.Println(sum)
}

func customSort(pairs []HandBidPair) []HandBidPair {
	sort.Slice(pairs, func(i, j int) bool {
		s1 := pairs[i].Hand
		s2 := pairs[j].Hand

		// Compare strings character by character
		for k := 0; k < len(s1) && k < len(s2); k++ {
			if rankWeights[rune(s1[k])] != rankWeights[rune(s2[k])] {
				return rankWeights[rune(s1[k])] < rankWeights[rune(s2[k])]
			}
		}

		// If the initial characters are the same, compare string lengths
		return len(s1) < len(s2)
	})
	return pairs
}

package day11a

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func Run() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	galaxyMap := make(map[int]Point)
	noColExpandMap := make(map[int]bool)

	row := 0
	galaxy_num := 0
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		seen := false
		for col, element := range line {
			if element == '#' {
				seen = true
				galaxy_num++
				point := Point{X: col, Y: row}
				galaxyMap[galaxy_num] = point
				//if galaxy found, mark it so that we know which columns to expand later on
				noColExpandMap[col] = true

			}
		}
		if !seen {
			//expanding universe, add another row
			row++
		}
		row++
	}

	//expand out columns
	offset := 0
	for k := 0; k < len(line); k++ {
		if noColExpandMap[k] == true {
			continue
		}
		for n := 1; n <= galaxy_num; n++ {
			if galaxyMap[n].X > k+offset {
				galaxyMap[n] = Point{X: galaxyMap[n].X + 1, Y: galaxyMap[n].Y}
			}
		}
		offset++
	}

	total_distances := 0

	for i := 1; i <= galaxy_num; i++ {
		for j := i + 1; j <= galaxy_num; j++ {
			if i == j {
				break
			}
			total_distances += abs(galaxyMap[i].X-galaxyMap[j].X) + abs(galaxyMap[i].Y-galaxyMap[j].Y)
		}
	}
	fmt.Println(total_distances)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package day10b

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0, 0)
	points := make([][]int, 0, 0)
	pos := []int{-1, 0}

	//read in file and locate 'S'
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	found := false
	for row, line := range lines {
		for idx, element := range line {
			if element == 'S' {
				pos[1] = idx
				found = true
				break
			}
		}
		if found {
			pos[0] = row
			break
		}
	}
	points = append(points, pos)

	//get first direction

	direction := []int{0, 0}
	if pos[0]+1 < len(lines) && contains([]rune{'|', 'L', 'J'}, rune(lines[pos[0]+1][pos[1]])) {
		direction = []int{1, 0} //down
	} else if pos[0]-1 >= 0 && contains([]rune{'|', 'F', '7'}, rune(lines[pos[0]-1][pos[1]])) {
		direction = []int{-1, 0} //up
	} else if pos[1]+1 < len(lines[0]) && contains([]rune{'-', 'J', '7'}, rune(lines[pos[0]][pos[1]+1])) {
		direction = []int{0, 1} //right
	} else if pos[1]-1 >= 0 && contains([]rune{'-', 'F', 'L'}, rune(lines[pos[0]][pos[1]-1])) {
		direction = []int{0, -1} //left
	}

	max := 1
	current_pos := make([]int, len(pos))
	copy(current_pos, pos)
	current_pos[0] += direction[0]
	current_pos[1] += direction[1]

	for current_pos[0] != pos[0] || current_pos[1] != pos[1] {
		cur := lines[current_pos[0]][current_pos[1]]
		point := make([]int, len(pos))
		copy(point, current_pos)
		points = append(points, point)
		if cur == 'L' {
			if direction[0] == 1 { // down
				direction = []int{0, 1} // go right
			} else {
				direction = []int{-1, 0} // go up
			}
		} else if cur == 'J' {
			if direction[0] == 1 { //down
				direction = []int{0, -1} // go left
			} else {
				direction = []int{-1, 0} // go up
			}
		} else if cur == 'F' {
			if direction[0] == -1 { //up
				direction = []int{0, 1} // go right
			} else {
				direction = []int{1, 0} // go down
			}
		} else if cur == '7' {
			if direction[0] == -1 { // up
				direction = []int{0, -1} //go left
			} else {
				direction = []int{1, 0} // go down
			}
		}
		// if cur == '|' or cur == '-' then can keep direction the same (up or down) or (left and rigth)
		current_pos[0] += direction[0]
		current_pos[1] += direction[1]
		max++
	}

	area := shoelace(points)
	//Pick's theorem
	interior_points := area + 1 - float64(max)/2

	// if odd number of steps, add one
	fmt.Println(interior_points)

}

func contains(arr []rune, target rune) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

// formula used to calculate area of polygon with corners defined as integer coordinates
func shoelace(points [][]int) float64 {
	n := len(points)
	if n < 3 {
		return 0.0 // Not a valid polygon
	}

	area := 0.0
	//Calculating determinants between points
	for i := 0; i < n-1; i++ {
		area += float64(points[i][0]*points[i+1][1] - points[i+1][0]*points[i][1])
	}
	area += float64(points[n-1][0]*points[0][1] - points[0][0]*points[n-1][1])

	return 0.5 * abs(area)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

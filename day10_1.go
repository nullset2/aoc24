package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(matrix [][]int, target []int) bool {
	for _, row := range matrix {
		if equalSlices(row, target) {
			return true
		}
	}
	return false
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func dfs(data [][]int, i, j int) int {
	score := 0
	N := len(data)

	stack := make([][]int, 0)
	visited := make([][]int, 0)
	// U D L R
	directions := [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
	}

	current := []int{i, j}
	stack = append(stack, current)

	for len(stack) != 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if data[current[0]][current[1]] == 9 {
			visited = append(visited, current)
			score++
			continue
		}
		if !contains(visited, current) {
			visited = append(visited, current)

			for _, direction := range directions {
				newPoint := []int{current[0] + direction[0],
					current[1] + direction[1]}

				if newPoint[0] >= 0 && newPoint[0] < N &&
					newPoint[1] >= 0 && newPoint[1] < N &&
					!contains(visited, newPoint) &&
					data[newPoint[0]][newPoint[1]] == data[current[0]][current[1]]+1 {

					stack = append(stack, newPoint)
				}
			}
		}
	}
	return score
}

func main() {
	file, _ := os.Open("input10.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		s := scanner.Text()
		var row []int

		for _, c := range s {
			x := int(c - '0')
			row = append(row, x)
		}

		data = append(data, row)
	}

	accum := 0

	for i, row := range data {
		for j, cell := range row {
			if cell == 0 {
				accum += dfs(data, i, j)
			}
		}
	}
	fmt.Println(accum)
}

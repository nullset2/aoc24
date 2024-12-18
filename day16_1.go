package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func contains(s [][]int, t []int) bool {
	for _, a := range s {
		if a[0] == t[0] && a[1] == t[1] && a[2] == t[2] {
			return true
		}
	}
	return false
}

func getSize(s string) int {
	file, _ := os.Open(s)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	file.Close()
	return len(scanner.Text())
}

func getScore(route [][]int) int {
	score := 0
	var lastSteps [][]int
	for _, x := range route {
		lastSteps = append(lastSteps, x)
		if len(lastSteps) == 2 {
			if lastSteps[0][2] != lastSteps[1][2] {
				score += 1000
			}
			lastSteps = lastSteps[1:]
		}
		score += 1
	}

	return score
}

func main() {
	N := getSize("input16.txt")
	file, _ := os.Open("input16.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := make([][]string, N)
	for i, _ := range grid {
		grid[i] = make([]string, N)
	}

	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		chars := strings.Split(s, "")
		grid[i] = chars
		i++
	}

	// N E S W => 0 1 2 3
	//player always starts here FACING EAST
	y, x := N-2, 1

	current := []int{y, x, 1}
	queue := [][]int{current}
	visited := [][]int{}
	possibleRoutes := [][][]int{}
	output := [][]int{}

	for len(queue) != 0 {
		current, queue = queue[0], queue[1:]
		visited = append(visited, current)

		if grid[current[0]][current[1]] == "E" {
			possibleRoutes = append(possibleRoutes, visited)
			continue
		}

		//N
		if (grid[current[0]-1][current[1]] == "." || grid[current[0]-1][current[1]] == "E") && !contains(visited, []int{current[0] - 1, current[1], 0}) {
			queue = append(queue, []int{current[0] - 1, current[1], 0})
		}

		//E
		if (grid[current[0]][current[1]+1] == "." || grid[current[0]][current[1]+1] == "E") && !contains(visited, []int{current[0], current[1] + 1, 1}) {

			queue = append(queue, []int{current[0], current[1] + 1, 1})
		}

		//S
		if (grid[current[0]+1][current[1]] == "." || grid[current[0]+1][current[1]] == "E") && !contains(visited, []int{current[0] + 1, current[1], 2}) {
			queue = append(queue, []int{current[0] + 1, current[1], 2})
		}

		//W
		if (grid[current[0]][current[1]-1] == "." || grid[current[0]][current[1]-1] == "E") && !contains(visited, []int{current[0], current[1] - 1, 3}) {
			queue = append(queue, []int{current[0], current[1] - 1, 3})
		}

		output = append(output, current)
	}

	min := 9999999999999

	for _, route := range possibleRoutes {
		route = route[1:]
		score := getScore(route)
		if score < min {
			min = score
		}
	}

	fmt.Println(min)

}

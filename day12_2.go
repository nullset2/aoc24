package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func bfs(data *[][]string, visited *[][]bool, k, v int) (int, int) {
	a := 0
	corners := 0
	N := len(*data)
	M := len((*data)[0])
	queue := [][]int{}
	queue = append(queue, []int{k, v})

	for len(queue) != 0 {
		var edges []string
		var current []int
		current, queue = queue[0], queue[1:]
		i, j := current[0], current[1]
		if (*visited)[i][j] {
			continue
		}
		c := (*data)[i][j]

		//up
		if i-1 < 0 || (*data)[i-1][j] != c {
			edges = append(edges, "up")
		} else if i-1 >= 0 && !(*visited)[i-1][j] {
			queue = append(queue, []int{i - 1, j})
		}

		//right
		if j+1 == M || (*data)[i][j+1] != c {
			edges = append(edges, "right")
		} else if j+1 <= M && !(*visited)[i][j+1] {
			queue = append(queue, []int{i, j + 1})
		}

		// down
		if i+1 == N || (*data)[i+1][j] != c {
			edges = append(edges, "down")
		} else if i+1 < N && !(*visited)[i+1][j] {
			queue = append(queue, []int{i + 1, j})
		}

		//left
		if j-1 < 0 || (*data)[i][j-1] != c {
			edges = append(edges, "left")
		} else if j-1 >= 0 && !(*visited)[i][j-1] {
			queue = append(queue, []int{i, j - 1})
		}

		(*visited)[i][j] = true
		a++

		if slices.Contains(edges, "up") {
			if slices.Contains(edges, "right") {
				corners++
			}

			if slices.Contains(edges, "left") {
				corners++
			}
		}
		if slices.Contains(edges, "down") {
			if slices.Contains(edges, "right") {
				corners++
			}

			if slices.Contains(edges, "left") {
				corners++
			}
		}

		if slices.Contains(edges, "left") {
			if i-1 >= 0 && j-1 >= 0 &&
				(*data)[i-1][j] == c &&
				(*data)[i-1][j-1] == c {
				corners++
			}

			if i+1 < N && j-1 >= 0 &&
				(*data)[i+1][j] == c &&
				(*data)[i+1][j-1] == c {
				corners++
			}
		}

		if slices.Contains(edges, "right") {
			if i-1 >= 0 && j+1 < M &&
				(*data)[i-1][j] == c &&
				(*data)[i-1][j+1] == c {
				corners++
			}

			if i+1 < N && j+1 < M &&
				(*data)[i+1][j] == c &&
				(*data)[i+1][j+1] == c {
				corners++
			}
		}

	}

	return a, corners
}

func main() {
	file, _ := os.Open("input12.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := [][]string{}

	for scanner.Scan() {
		s := scanner.Text()
		row := strings.Split(s, "")
		data = append(data, row)
	}

	total := 0

	N := len(data)
	M := len(data[0])

	visited := make([][]bool, N)
	for i, _ := range visited {
		visited[i] = make([]bool, M)
	}

	for i, row := range data {
		for j, c := range row {
			if !visited[i][j] {
				a, corners := bfs(&data, &visited, i, j)
				fmt.Println(c, ": ", a, ",", corners)
				total += a * corners
			}
		}
	}
	fmt.Println(total)
}

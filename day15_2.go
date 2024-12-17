package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contains(s [][]int, t []int) bool {
	for _, a := range s {
		if a[0] == t[0] && a[1] == t[1] {
			return true
		}
	}
	return false
}

func printGrid(grid [][]string) {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < len(grid)/2; i++ {
		for j := 0; j < len(grid); j++ {
			fmt.Printf("%s", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	N := 100
	commands := []string{}
	file, _ := os.Open("input15_2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := make([][]string, N)
	for i, _ := range grid {
		grid[i] = make([]string, N/2)
	}

	x, y := 0, 0
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		chars := strings.Split(s, "")

		for j, c := range chars {
			if c == "@" {
				y = i
				x = j
			}
		}
		grid[i] = chars
		i++
	}

	for scanner.Scan() {
		s := scanner.Text()
		commands = append(commands, strings.Split(s, "")...)
	}

	for _, command := range commands {
		lastStep := []int{y, x}
		printGrid(grid)
		fmt.Println("now moving.......", command)
		//bufio.NewReader(os.Stdin).ReadString('\n')

		if command == "^" && grid[y-1][x] != "#" {
			y = y - 1
		} else if command == ">" && grid[y][x+1] != "#" {
			x = x + 1
		} else if command == "v" && grid[y+1][x] != "#" {
			y = y + 1
		} else if command == "<" && grid[y][x-1] != "#" {
			x = x - 1
		}

		var connectedCrates [][]int

		if grid[y][x] == "[" || grid[y][x] == "]" {
			if command == "^" {
				stack := [][]int{}
				visited := [][]int{}
				current := []int{y, x}
				stack = append(stack, current)

				for len(stack) != 0 {
					current, stack = stack[len(stack)-1], stack[:len(stack)-1]
					visited = append(visited, current)

					if current[0] < 0 || current[0] >= N/2 ||
						current[1] < 0 || current[1] >= N ||
						grid[current[0]][current[1]] == "." {
						continue
					}

					if grid[current[0]][current[1]] == "[" && !contains(visited, []int{current[0], current[1] + 1}) {
						stack = append(stack, []int{current[0], current[1] + 1})
					}

					if grid[current[0]][current[1]] == "]" && !contains(visited, []int{current[0], current[1] - 1}) {
						stack = append(stack, []int{current[0], current[1] - 1})
					}

					if (grid[current[0]][current[1]] == "[" || grid[current[0]][current[1]] == "]") &&
						!contains(visited, []int{current[0] - 1, current[1]}) {
						stack = append(stack, []int{current[0] - 1, current[1]})
					}
					connectedCrates = append(connectedCrates, current)
				}
			} else if command == ">" {
				for k := 0; x+k < N && (grid[y][x+k] == "[" || grid[y][x+k] == "]"); k++ {
					connectedCrates = append(connectedCrates, []int{y, x + k})
				}
			} else if command == "v" {
				stack := [][]int{}
				visited := [][]int{}
				current := []int{y, x}
				stack = append(stack, current)

				for len(stack) != 0 {
					current, stack = stack[len(stack)-1], stack[:len(stack)-1]
					visited = append(visited, current)

					if current[0] < 0 || current[0] >= N/2 ||
						current[1] < 0 || current[1] >= N ||
						grid[current[0]][current[1]] == "." {
						continue
					}

					if grid[current[0]][current[1]] == "[" && !contains(visited, []int{current[0], current[1] + 1}) {
						stack = append(stack, []int{current[0], current[1] + 1})
					}

					if grid[current[0]][current[1]] == "]" && !contains(visited, []int{current[0], current[1] - 1}) {
						stack = append(stack, []int{current[0], current[1] - 1})
					}

					if (grid[current[0]][current[1]] == "[" || grid[current[0]][current[1]] == "]") &&
						!contains(visited, []int{current[0] + 1, current[1]}) {
						stack = append(stack, []int{current[0] + 1, current[1]})
					}
					connectedCrates = append(connectedCrates, current)
				}
			} else if command == "<" {
				for k := 0; x-k >= 0 && (grid[y][x-k] == "[" || grid[y][x-k] == "]"); k++ {
					connectedCrates = append(connectedCrates, []int{y, x - k})
				}
			}

			lastCrate := connectedCrates[len(connectedCrates)-1]
			moves := false

			if command == "^" {
				for _, crate := range connectedCrates {
					if crate[0]-1 < 0 || grid[crate[0]-1][crate[1]] == "#" {
						moves = false
						break
					}
					moves = true
				}

				if !moves {
					y = lastStep[0]
					x = lastStep[1]
					continue
				}

				content := make(map[int]string)

				for i, crate := range connectedCrates {
					content[i] = grid[crate[0]][crate[1]]
				}

				for _, crate := range connectedCrates {
					grid[crate[0]][crate[1]] = "."
				}

				for i, crate := range connectedCrates {
					grid[crate[0]-1][crate[1]] = content[i]
				}

			} else if command == ">" && grid[lastCrate[0]][lastCrate[1]+1] != "#" {
				for k := len(connectedCrates) - 1; k >= 0; k-- {
					grid[y][x+k+1] = grid[y][x+k]
				}
			} else if command == "v" && grid[lastCrate[0]+1][lastCrate[1]] != "#" {
				for _, crate := range connectedCrates {
					if crate[0]+1 == N/2 || grid[crate[0]+1][crate[1]] == "#" {
						moves = false
						break
					}
					moves = true
				}

				if !moves {
					y = lastStep[0]
					x = lastStep[1]
					continue
				}

				content := make(map[int]string)

				for i, crate := range connectedCrates {
					content[i] = grid[crate[0]][crate[1]]
				}

				for _, crate := range connectedCrates {
					grid[crate[0]][crate[1]] = "."
				}

				for i, crate := range connectedCrates {
					grid[crate[0]+1][crate[1]] = content[i]
				}
			} else if command == "<" && grid[lastCrate[0]][lastCrate[1]-1] != "#" {
				for k := len(connectedCrates) - 1; k >= 0; k-- {
					grid[y][x-k-1] = grid[y][x-k]
				}
			} else {
				x, y = lastStep[1], lastStep[0]
			}
		}

		grid[lastStep[0]][lastStep[1]] = "."
		grid[y][x] = "@"
	}

	printGrid(grid)

	sum := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == "[" {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)
}

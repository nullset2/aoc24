package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printGrid(grid [][]string) {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			fmt.Printf("%s", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	N := 50
	commands := []string{}
	file, _ := os.Open("input15.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := make([][]string, N)
	for i, _ := range grid {
		grid[i] = make([]string, N)
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

		if grid[y][x] == "O" {
			if command == "^" {
				for k := 0; y-k >= 0 && grid[y-k][x] == "O"; k++ {
					connectedCrates = append(connectedCrates, []int{y - k, x})
				}
			} else if command == ">" {
				for k := 0; x+k < N && grid[y][x+k] == "O"; k++ {
					connectedCrates = append(connectedCrates, []int{y, x + k})
				}
			} else if command == "v" {
				for k := 0; y+k < N && grid[y+k][x] == "O"; k++ {
					connectedCrates = append(connectedCrates, []int{y + k, x})
				}
			} else if command == "<" {
				for k := 0; x-k >= 0 && grid[y][x-k] == "O"; k++ {
					connectedCrates = append(connectedCrates, []int{y, x - k})
				}
			}

			lastCrate := connectedCrates[len(connectedCrates)-1]
			if command == "^" && grid[lastCrate[0]-1][lastCrate[1]] != "#" {
				for k := len(connectedCrates) - 1; k >= 0; k-- {
					grid[y-k-1][x] = grid[y-k][x]
				}
			} else if command == ">" && grid[lastCrate[0]][lastCrate[1]+1] != "#" {
				for k := len(connectedCrates) - 1; k >= 0; k-- {
					grid[y][x+k+1] = grid[y][x+k]
				}
			} else if command == "v" && grid[lastCrate[0]+1][lastCrate[1]] != "#" {
				for k := len(connectedCrates) - 1; k >= 0; k-- {
					grid[y+k+1][x] = grid[y+k][x]
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

	sum := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == "O" {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)
}

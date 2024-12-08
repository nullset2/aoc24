package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func getSize(s string) int {
	file, _ := os.Open(s)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	file.Close()
	return len(scanner.Text())
}

func printMatrix(mat [][]rune) {
	exec.Command("clear").Run()
	for _, row := range mat {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	N := getSize("input6.txt")

	file, _ := os.Open("input6.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mat := make([][]rune, N)
	for i := range mat {
		mat[i] = make([]rune, N)
	}

	starty := 0
	startx := 0

	i := 0
	for scanner.Scan() && i < N {
		s := scanner.Text()
		for j, c := range s {
			mat[i][j] = c
			if c == '^' {
				starty = i
				startx = j
			}
		}
		i++
	}

	direction := 0

	x := startx
	y := starty

	count := 0

	for x < N && y < N {
		if mat[y][x] != 'X' {
			count++
		}
		mat[y][x] = 'X'

		if direction == 0 {
			if y-1 == N {
				break
			}
			if mat[y-1][x] == '#' {
				direction = 1
				continue
			}
			y = y - 1
		} else if direction == 1 {
			if x+1 == N {
				break
			}
			if mat[y][x+1] == '#' {
				direction = 2
				continue
			}
			x = x + 1
		} else if direction == 2 {
			if y+1 == N {
				break
			}
			if mat[y+1][x] == '#' {
				direction = 3
				continue
			}
			y = y + 1
		} else if direction == 3 {
			if x-1 == N {
				break
			}
			if mat[y][x-1] == '#' {
				direction = 0
				continue
			}
			x = x - 1
		}
		printMatrix(mat)
	}

	fmt.Println(count)
}

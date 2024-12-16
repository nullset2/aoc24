package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Robot struct {
	vx int
	vy int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func saveImage(matrix [][][]Robot, filename string) error {
	height := len(matrix)
	if height == 0 {
		return nil
	}
	width := len(matrix[0])

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if len(matrix[y][x]) > 0 {
				img.SetGray(x, y, color.Gray{Y: 0})
			} else {
				img.SetGray(x, y, color.Gray{Y: 255})
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func printGrid(grid [][][]Robot) {
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if len(grid[i][j]) != 0 {
				fmt.Printf("%d", len(grid[i][j]))
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func main() {
	N, M := 101, 103

	file, _ := os.Open("input14.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := make([][][]Robot, N)
	for i, _ := range grid {
		grid[i] = make([][]Robot, M)
		for j, _ := range grid[i] {
			grid[i][j] = make([]Robot, 0)
		}
	}

	for scanner.Scan() {
		s := scanner.Text()
		x, y, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(s, "p=%v,%v v=%v,%v", &x, &y, &vx, &vy)

		grid[x][y] = append(grid[x][y], Robot{vx, vy})
	}

	for i := 0; i < 10000; i++ {
		saveImage(grid, fmt.Sprintf("%d.png", i))
		tmp := make([][][]Robot, N)
		for i, _ := range tmp {
			tmp[i] = make([][]Robot, M)
			for j, _ := range tmp[i] {
				tmp[i][j] = make([]Robot, 0)
			}
		}

		for x, row := range grid {
			for y, cell := range row {
				for _, robot := range cell {
					newX := x + robot.vx
					if newX < 0 {
						newX = N + newX
					}
					newY := y + robot.vy
					if newY < 0 {
						newY = M + newY
					}
					tmp[newX%N][newY%M] = append(tmp[newX%N][newY%M], robot)
				}
			}
		}

		for i, row := range tmp {
			for j, cell := range row {
				grid[i][j] = cell
			}
		}
	}

	scores := []int{}
	score := 0
	for i := 0; i < N/2; i++ {
		for j := 0; j < M/2; j++ {
			score += len(grid[i][j])
		}
	}
	scores = append(scores, score)
	score = 0
	for i := (N / 2) + 1; i < N; i++ {
		for j := 0; j < M/2; j++ {
			score += len(grid[i][j])
		}
	}
	scores = append(scores, score)

	score = 0
	for i := 0; i < N/2; i++ {
		for j := (M / 2) + 1; j < M; j++ {
			score += len(grid[i][j])
		}
	}
	scores = append(scores, score)
	score = 0
	for i := (N / 2) + 1; i < N; i++ {
		for j := (M / 2) + 1; j < M; j++ {
			score += len(grid[i][j])
		}
	}
	scores = append(scores, score)
	score = 0

	product := 1

	for _, x := range scores {
		product *= x
	}

	fmt.Println(scores)
	fmt.Println(product)
}

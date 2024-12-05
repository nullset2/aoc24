package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	WORD := "XMAS"
	count := 0

	file, _ := os.Open("input4.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()
	N := len(s)
	mat := make([][]rune, N)
	for row := range mat {
		mat[row] = make([]rune, N)
	}
	file.Close()

	file, _ = os.Open("input4.txt")
	defer file.Close()
	scanner = bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		s = scanner.Text()
		for j, c := range s {
			mat[i][j] = rune(c)
		}
		i++
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if mat[i][j] == 'X' {
				// n
				for k := 0; i-k >= 0 && k < len(WORD); k++ {
					if mat[i-k][j] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i-k][j] == 'S' {
						count++
					}
				}

				// ne
				for k := 0; i-k >= 0 && j+k < N && k < len(WORD); k++ {
					if mat[i-k][j+k] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i-k][j+k] == 'S' {
						count++
					}
				}
				// e
				for k := 0; j+k < N && k < len(WORD); k++ {
					if mat[i][j+k] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i][j+k] == 'S' {
						count++
					}
				}
				// se
				for k := 0; i+k < N && j+k < N && k < len(WORD); k++ {
					if mat[i+k][j+k] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i+k][j+k] == 'S' {
						count++
					}
				}

				// s
				for k := 0; i+k < N && k < len(WORD); k++ {
					if mat[i+k][j] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i+k][j] == 'S' {
						count++
					}
				}

				// sw
				for k := 0; i+k < N && j-k >= 0 && k < len(WORD); k++ {
					if mat[i+k][j-k] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i+k][j-k] == 'S' {
						count++
					}
				}

				// w
				for k := 0; j-k >= 0 && k < len(WORD); k++ {
					if mat[i][j-k] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i][j-k] == 'S' {
						count++
					}
				}

				// nw
				for k := 0; i-k >= 0 && j-k >= 0 && k < len(WORD); k++ {
					if mat[i-k][j-k] != rune(WORD[k]) {
						break
					}
					if k == 3 && mat[i-k][j-k] == 'S' {
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)
}

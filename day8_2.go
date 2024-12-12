package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func combinations(pairs [][]int) [][][]int {
	result := [][][]int{}

	for i := 0; i < len(pairs)-1; i++ {
		for j := i + 1; j < len(pairs); j++ {
			combination := [][]int{pairs[i], pairs[j]}
			result = append(result, combination)
		}
	}

	return result
}

func main() {
	N := 0
	list := make(map[string]int)
	points := make(map[string][][]int)
	file, _ := os.Open("input8.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		s := scanner.Text()
		N = len(s)
		for x, c := range s {
			_, ok := points[string(c)]
			if c != '.' {
				if !ok {
					points[string(c)] = make([][]int, 0)
				}
				points[string(c)] = append(points[string(c)], []int{x, y})
			}
		}
		y++
	}

	for _, v := range points {
		combos := combinations(v)
		for _, c := range combos {
			p := c[0]
			q := c[1]
			dx := q[0] - p[0]
			dy := q[1] - p[1]

			var sb strings.Builder
			sb.WriteString(strconv.Itoa(p[0]))
			sb.WriteString(",")
			sb.WriteString(strconv.Itoa(p[1]))
			tmp := sb.String()
			list[tmp] = 1

			var sh strings.Builder
			sh.WriteString(strconv.Itoa(q[0]))
			sh.WriteString(",")
			sh.WriteString(strconv.Itoa(q[1]))
			tmp = sh.String()
			list[tmp] = 1

			for i, j := dx, dy; i < N && j < N; i, j = i + dx, j + dy {
				antinode1 := []int{p[0] - i, p[1] - j}
				antinode2 := []int{q[0] + i, q[1] + j}

				if antinode1[0] >= 0 && antinode1[0] < N &&
					antinode1[1] >= 0 && antinode1[1] < N {
					var sb strings.Builder
					sb.WriteString(strconv.Itoa(antinode1[0]))
					sb.WriteString(",")
					sb.WriteString(strconv.Itoa(antinode1[1]))
					tmp := sb.String()
					list[tmp] = 1
				}

				if antinode2[0] >= 0 && antinode2[0] < N &&
					antinode2[1] >= 0 && antinode2[1] < N {
					var sb strings.Builder
					sb.WriteString(strconv.Itoa(antinode2[0]))
					sb.WriteString(",")
					sb.WriteString(strconv.Itoa(antinode2[1]))
					tmp := sb.String()
					list[tmp] = 1
				}
			}
		}
	}

	fmt.Println(len(list))
}

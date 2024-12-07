package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	rules := make([][]int, 0)
	orders := make([][]int, 0)
	file, _ := os.Open("input5.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, "|")
		rule := make([]int, 0)
		for _, p := range parts {
			tmp, _ := strconv.Atoi(p)
			rule = append(rule, tmp)
		}
		rules = append(rules, rule)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		order := make([]int, 0)
		for _, p := range parts {
			tmp, _ := strconv.Atoi(p)
			order = append(order, tmp)
		}
		orders = append(orders, order)
	}

	accum := 0
	for _, order := range orders {
		valid := true
		a := order[:len(order)-1]
		b := order[1:]
		for i := 0; i < len(a); i++ {
			cur := []int{a[i], b[i]}
			if !contains(rules, cur) {
				valid = false
				break
			}
		}
		if valid {
			accum += order[len(order)/2]
		}
	}
	fmt.Println(accum)
}

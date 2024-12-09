package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func aux(i int, size int, accumulator string, set []string) []string {
	if size == len(accumulator) {
		set = append(set, accumulator)
		return set
	}

	for _, l := range "+*" {
		current := accumulator + string(l)
		set = aux(i+1, size, current, set)
		current = current[:len(current)-1]
	}
	return set
}

func generateOperators(size int) []string {
	return aux(0, size, "", []string{})
}

func main() {
	passes := make(map[int]int)
	accum := 0
	file, _ := os.Open("input7.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		numbers := strings.Split(s, ": ")
		result, _ := strconv.Atoi(numbers[0])
		tmp := strings.Split(numbers[1], " ")
		operands := make([]int, len(tmp))
		for i, n := range tmp {
			operands[i], _ = strconv.Atoi(n)
		}

		operators := generateOperators(len(operands) - 1)

		for _, op := range operators {
			k := 0
			total := operands[k]
			for _, o := range op {
				k = k + 1
				if k == len(operands) {
					break
				} else if o == '+' {
					total += operands[k]
				} else if o == '*' {
					total *= operands[k]
				}
			}
			if total == result {
				passes[result] = 1
			}
		}
	}

	for k, _ := range passes {
		accum += k
	}

	fmt.Println(accum)
}

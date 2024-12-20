package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var A, B, C int
	var program string
	var output []string
	combo := func(op int) int {
		if op <= 3 {
			return op
		} else if op == 4 {
			return A
		} else if op == 5 {
			return B
		}
		return C
	}

	file, _ := os.Open("input17.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()
	fmt.Sscanf(s, "Register A: %d", &A)
	scanner.Scan()
	s = scanner.Text()
	fmt.Sscanf(s, "Register B: %d", &B)
	scanner.Scan()
	s = scanner.Text()
	fmt.Sscanf(s, "Register C: %d", &C)
	scanner.Scan()
	scanner.Scan()
	s = scanner.Text()
	fmt.Sscanf(s, "Program: %s", &program)

	parts := strings.Split(program, ",")
	c := make([]int, len(parts))
	for i, p := range parts {
		c[i], _ = strconv.Atoi(p)
	}
	i := 0
	for i < len(c) {
		fmt.Println("Current", c[i])
		fmt.Println(A)
		fmt.Println(B)
		fmt.Println(C)
		switch c[i] {
		case 0: //adv
			A = int(math.Round(float64(A / int(math.Pow(2, float64(combo(c[i+1])))))))
			i += 2
		case 1: //bxl
			B = B ^ c[i+1]
			i += 2
		case 2: //bst
			B = combo(c[i+1]) % 8
			i += 2
		case 3: //jnz
			if A == 0 {
				i += 2
				continue
			}
			i = c[i+1]
		case 4: //bxc
			B = B ^ C
			i += 2
		case 5: //out
			out := combo(c[i+1]) % 8
			output = append(output, strconv.Itoa(out))
			i += 2
		case 6: //bdv
			B = int(math.Round(float64(A / int(math.Pow(2, float64(combo(c[i+1])))))))
			i += 2
		case 7: //cdv
			C = int(math.Round(float64(A / int(math.Pow(2, float64(combo(c[i+1])))))))
			i += 2
		}
	}
	fmt.Println(strings.Join(output, ","))
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left, right []int

	input, _ := os.Open("input1.txt")
	defer input.Close()

	reader := bufio.NewReader(input)
	for {
		line, _, _ := reader.ReadLine()
		s := string(line)
		if len(s) == 0 {
			break
		}

		parts := strings.Split(s, "   ")
		leftval, _ := strconv.Atoi(parts[0])
		left = append(left, leftval)
		rightval, _ := strconv.Atoi(parts[1])
		right = append(right, rightval)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(sum)
}

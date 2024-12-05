package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	safe := 0
	input, _ := os.Open("input2.txt")
	defer input.Close()
	reader := bufio.NewReader(input)
	for {
		var levels []int
		line, _, err := reader.ReadLine()
		s := string(line)
		if len(s) == 0 || err != nil {
			break
		}
		parts := strings.Split(s, " ")
		for _, n := range parts {
			x, _ := strconv.Atoi(n)
			levels = append(levels, x)
		}

		increasing := false
		sf := false
		increasing = levels[0]-levels[1] < 0
		for i, x := range levels {
			if i+1 == len(levels) {
				break
			}
			diff := x - levels[i+1]
			sf = (increasing && diff < 0 && diff >= -3) || (!increasing && diff > 0 && diff <= 3)
			if !sf {
				break
			}
		}

		if sf {
			safe++
		}
	}
	fmt.Println(safe)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.Open("input3.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	accum := 0
	r := regexp.MustCompile(`mul\((?P<n1>\d+),(?P<n2>\d+)\)`)
	for {
		line, _, err := reader.ReadLine()
		if len(line) == 0 || err != nil {
			break
		}
		s := string(line)
		m := r.FindAllStringSubmatch(s, -1)
		for _, cur := range m {
			p, _ := strconv.Atoi(cur[1])
			q, _ := strconv.Atoi(cur[2])
			accum += p * q
		}
	}
	fmt.Println(accum)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input22.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var output []int

	for scanner.Scan() {
		s := scanner.Text()
		n, _ := strconv.Atoi(s)
		for i := 0; i < 2000; i++ {
			s1 := ((n * 64) ^ n) % 16777216
			s2 := ((int(math.Round(float64(s1 / 32)))) ^ s1) % 16777216
			s3 := (s2*2048 ^ s2) % 16777216
			fmt.Println(s3)
			n = s3
		}

		output = append(output, n)
	}

	sum := 0
	for _, x := range output {
		sum += x
	}

	fmt.Println(sum)

}

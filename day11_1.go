package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input11.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stones := make(map[int]int)

	for scanner.Scan() {
		s := scanner.Text()
		tmp := strings.Split(s, " ")
		for _, c := range tmp {
			parsed, _ := strconv.Atoi(c)
			stones[parsed] = 1
		}
	}

	temp := make(map[int]int)

	for k, v := range stones {
		temp[k] = v
	}

	for i := 0; i < 25; i++ {
		for k, v := range stones {
			if k == 0 {
				temp[1] += v
			} else if len(strconv.Itoa(k))%2 == 0 {
				str := strconv.Itoa(k)
				l, _ := strconv.Atoi(str[:len(str)/2])
				r, _ := strconv.Atoi(str[len(str)/2:])
				temp[l] += v
				temp[r] += v
			} else {
				temp[k*2024] += v
			}
			temp[k] -= v
		}

		for k, v := range temp {
			stones[k] = v
		}
	}

	accum := 0
	for _, v := range stones {
		accum += v
	}
	fmt.Println(accum)
}

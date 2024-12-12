package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input9.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()

	var data []int
	id := 0

	for i, c := range s {
		info := int(c - '0')
		val := id

		if i%2 != 0 {
			val = -1
		} else {
			id++
		}

		for j := 0; j < info; j++ {
			data = append(data, val)
		}
	}

	N := len(data)
	i := 0
	j := N - 1

	for i < N && j > i {
		for i < N && data[i] != -1 {
			i++
		}

		for j > i && data[j] == -1 {
			j--
		}

		data[i], data[j] = data[j], -1
		j--
	}

	checksum := 0
	for i, x := range data {
		if x != -1 {
			checksum += x * i
		}
	}
	fmt.Println(checksum)

}

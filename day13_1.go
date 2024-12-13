package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

func isInteger(f float64) bool {
	rounded := math.Round(f)
	return math.Abs(f-rounded) < 1e-9
}

func main() {
	file, _ := os.Open("input13.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tokens := 0

	for scanner.Scan() {
		var a, b, c, d, e, f float64
		fmt.Sscanf(scanner.Text(), "Button A: X+%v, Y+%v", &a, &c)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%v, Y+%v", &b, &d)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%v, Y=%v", &e, &f)
		scanner.Scan()

		coef := mat.NewDense(2, 2, []float64{a, b, c, d})
		res := mat.NewVecDense(2, []float64{e, f})

		var x mat.VecDense
		err := x.SolveVec(coef, res)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var ans []int
		failed := false
		for _, root := range x.RawVector().Data {
			if isInteger(root) {
				ans = append(ans, int(math.Round(root)))
			} else {
				failed = true
				break
			}
		}
		if failed {
			continue
		}

		tokens += 3 * ans[0]
		tokens += ans[1]
	}

	fmt.Println(tokens)

}

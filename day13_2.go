package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func solveSystemOfEquations(a, b, c, d, e, f *big.Float) (*big.Float, *big.Float, error) {
	det := new(big.Float).Sub(new(big.Float).Mul(a, d), new(big.Float).Mul(b, c))

	if det.Cmp(big.NewFloat(0)) == 0 {
		return nil, nil, fmt.Errorf("")
	}

	detX := new(big.Float).Sub(new(big.Float).Mul(e, d), new(big.Float).Mul(b, f))
	detY := new(big.Float).Sub(new(big.Float).Mul(a, f), new(big.Float).Mul(e, c))

	x := new(big.Float).Quo(detX, det)
	y := new(big.Float).Quo(detY, det)

	return x, y, nil
}

func main() {
	file, _ := os.Open("input13.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tokens := int64(0)

	for scanner.Scan() {
		var a, b, c, d, e, f float64
		fmt.Sscanf(scanner.Text(), "Button A: X+%v, Y+%v", &a, &c)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%v, Y+%v", &b, &d)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%v, Y=%v", &e, &f)
		scanner.Scan()

		e = e + 10000000000000
		f = f + 10000000000000

		ab := big.NewFloat(a)
		bb := big.NewFloat(b)
		cb := big.NewFloat(c)
		db := big.NewFloat(d)
		eb := big.NewFloat(e)
		fb := big.NewFloat(f)

		x, y, err := solveSystemOfEquations(ab, bb, cb, db, eb, fb)
		if err != nil {
			fmt.Println(err)
			return
		}

		if x.IsInt() && y.IsInt() {
			i, _ := x.Int64()
			tokens += 3 * i
			j, _ := y.Int64()
			tokens += j
		}
	}

	fmt.Println(tokens)

}

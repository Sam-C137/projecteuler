package power_digit_sum

import (
	"math/big"
	"strconv"
)

func Run(maxPow int) int {
	var acc int
	exp := new(big.Int)
	exp.SetString("1", 10)
	exp.Lsh(exp, uint(maxPow))

	for _, n := range exp.String() {
		num, _ := strconv.Atoi(string(n))
		acc += num
	}

	return acc
}

package factorial_digit_sum

import (
	"math/big"
	"strconv"
)

func Run(fac *big.Int) int {
	f := factorial(fac, make(map[string]*big.Int))
	var acc int

	for _, s := range f.String() {
		n, _ := strconv.Atoi(string(s))
		acc += n
	}

	return acc
}

func factorial(num *big.Int, dict map[string]*big.Int) *big.Int {
	base := big.NewInt(1)

	if num.Cmp(base) <= 0 {
		return big.NewInt(1)
	}

	if val, ok := dict[num.String()]; ok {
		return val
	}

	one := big.NewInt(1)
	sub := new(big.Int).Sub(num, one)
	val := new(big.Int).Mul(num, factorial(sub, dict))
	dict[num.String()] = val

	return val
}

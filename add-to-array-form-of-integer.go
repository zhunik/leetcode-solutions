package solutions

import (
	"bytes"
	"math/big"
	"strconv"
)

func AddToArrayForm(A []int, K int) []int {

	var a bytes.Buffer

	for i := 0; i < len(A); i++ {
		a.WriteString(strconv.Itoa(A[i]))
	}

	aBInt := new(big.Int)
	aBInt, _ = aBInt.SetString(a.String(), 10)
	res := aBInt.Add(aBInt, big.NewInt(int64(K)))

	resString := res.String()

	var result = make([]int, len(resString))

	for i, b := range resString {
		tmp, _ := strconv.Atoi(string(b))
		result[i] = tmp
	}

	return result
}

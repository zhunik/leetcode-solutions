package problems

import (
	"bytes"
	"math/big"
	"strconv"
)

func addToArrayForm(A []int, K int) []int {

	var a bytes.Buffer
	var result []int

	for i := 0; i < len(A); i++ {
		a.WriteString(strconv.Itoa(A[i]))
	}

	aBInt := new(big.Int)
	aBInt, _ = aBInt.SetString(a.String(), 10)
	res := aBInt.Add(aBInt, big.NewInt(int64(K)))

	for _, b := range res.String() {
		tmp, _ := strconv.Atoi(string(b))
		result = append(result, tmp)
	}

	return result
}

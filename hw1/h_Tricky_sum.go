//https://contest.yandex.ru/contest/25606/problems/H/

package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		A      float64
		B      float64
		SumSqr float64
	)

	fmt.Scan(&A)
	fmt.Scan(&B)

	SumSqr = math.Pow((A + B), 2)

	fmt.Println(SumSqr)

}

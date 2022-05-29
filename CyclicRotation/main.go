package main

import (
	"fmt"
)

func Solution(A []int, k int) []int {
	if k < 0 || len(A) == 0 || k > 100 {
		return A
	}
	r := len(A) - k%len(A)
	A = append(A[r:], A[:r]...)

	fmt.Println(A)
	return A
}

func main() {
	var N int = 4
	var K int
	A := make([]int, N)

	A[0] = 4
	A[1] = 2
	A[2] = 1
	A[3] = 3

	fmt.Println("На сколько элементов вы хотите сместить элементы массива вправо?")
	fmt.Scan(&K)
	fmt.Println("Полученный массив:")
	Solution(A, K)
}

package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) bool {
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i-1] != (A[i] - 1) {
			return false
		}
	}
	return true
}

func main() {
	var N int = 4
	A := make([]int, N)

	A[0] = 4
	A[1] = 2
	//A[2] = 1
	A[3] = 3

	fmt.Println("Является ли массив последовательностью -", Solution(A))
}

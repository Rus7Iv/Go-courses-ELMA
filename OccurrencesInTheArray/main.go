package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)
	var A1 int

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] != A[j] {
				A1 = A[i]
			}
		}
	}
	return A1
}

func main() {
	var N int = 7
	A := make([]int, N)

	A[0] = 9
	A[1] = 3
	A[2] = 9
	A[3] = 3
	A[4] = 9
	A[5] = 7
	A[6] = 9

	fmt.Println("Элемент, встречающийся один раз: ", Solution(A))
}

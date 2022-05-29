package main

import "fmt"

func Solution(A []int, N int) int {
	var sum int = 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	var search int = N*(N+1)/2 - sum
	return search
}

func main() {
	var N int = 4
	A := make([]int, N+1)

	A[0] = 0
	//A[1] = 1
	A[2] = 2
	A[3] = 3
	A[4] = 4

	fmt.Println("Индекс отсутсвующего элемента равен ", Solution(A, N))
}

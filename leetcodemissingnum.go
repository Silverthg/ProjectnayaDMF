package main

import (
	"fmt"
)

func find_missing_number(arr []int, n int) int {
	sum := 0
	sum1 := 0
	for i := 0; i < n; i++ {
		sum += arr[i] //
	}
	sum1 = ((n + 1) * (n + 2)) / 2
	return (sum1 - sum)
}
func main() {
	arr := []int{5, 6, 8, 1, 2, 9, 4, 3}
	size1 := len(arr)
	missing_no := find_missing_number(arr, size1)
	fmt.Printf("В массиве с размером %d пропущена %d", size1+1, missing_no)
}

package main

import (
	"fmt"
)

func main() {
	a := []int32{3, 2, 1}
	bubbleSort(a)
}

func bubbleSort(a []int32) {
	var totalSwap int32
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j+1]
				move := a[j]
				a[j] = temp
				a[j+1] = move
				totalSwap += 1
			}
		}
	}
	fmt.Printf("Array is sorted in %d swaps. \n", totalSwap)
	fmt.Printf("First Element: %d  \n", a[0])
	fmt.Printf("Last Element: %d  \n", a[len(a)-1])
}

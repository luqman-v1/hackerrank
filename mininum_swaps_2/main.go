package main

import (
	"log"
)

func main() {
	log.Println(minimumSwaps([]int32{2, 31, 1, 38, 29, 5, 44, 6, 12, 18, 39, 9, 48, 49, 13, 11, 7, 27, 14, 33, 50, 21, 46, 23, 15, 26, 8, 47, 40, 3, 32, 22, 34, 42, 16, 41, 24, 10, 4, 28, 36, 30, 37, 35, 20, 17, 45, 43, 25, 19}))

}

//func optimizeMinSwap(arr []int32) int32 {
//	m := make(map[int]int, 0)
//	for _, v := range arr {
//		m[int(v)] = int(v)
//	}
//
//	totalSwap := 0
//	for i := 0; i < len(arr); i++ {
//		smallestValueIndex := i
//		for arr[]{
//
//		}
//
//		for j := i; j < len(arr); j++ {
//			if arr[j] < arr[smallestValueIndex] {
//				smallestValueIndex = j
//			}
//		}
//		//swap
//		temp := arr[i]
//		if arr[smallestValueIndex] != temp {
//			arr[i] = arr[smallestValueIndex]
//			arr[smallestValueIndex] = temp
//			totalSwap += 1
//		}
//	}
//	return int32(totalSwap)
//}

//minimumSwaps use algorithm selection sort
func minimumSwaps(arr []int32) int32 {
	totalSwap := 0
	for i := 0; i < len(arr); i++ {
		smallestValueIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[smallestValueIndex] {
				smallestValueIndex = j
			}
		}
		//swap
		temp := arr[i]
		if arr[smallestValueIndex] != temp {
			arr[i] = arr[smallestValueIndex]
			arr[smallestValueIndex] = temp
			totalSwap += 1
		}
	}

	return int32(totalSwap)
}

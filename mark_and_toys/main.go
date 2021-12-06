package main

import (
	"log"
	"sort"
)

func NewMax(prices []int32, k int32) int32 {
	var sorted []int
	p := prices
	for _, v := range p {
		sorted = append(sorted, int(v))
	}
	sort.Ints(sorted)
	var totalMaxToys int32

	for _, v := range sorted {
		k = k - int32(v)
		if k >= 0 {
			totalMaxToys += 1
		}
	}
	return totalMaxToys
}

func maximumToys(prices []int32, k int32) int32 {
	var sorted []int
	p := prices
	for _, v := range p {
		sorted = append(sorted, int(v))
	}
	sort.Ints(sorted)
	var totalMaxToys int
	originalK := k
	for i := 0; i < len(sorted); i++ {
		var tempToys int
		var totalBuy int
		for j := 0; j <= len(sorted)-1; j++ {
			nowK := int(k) - sorted[j]
			if nowK <= 0 {
				continue
			}
			k = int32(int(k) - sorted[j])
			totalBuy += sorted[j]
			if int(originalK) >= totalBuy {
				tempToys += 1
			}
		}
		if tempToys > totalMaxToys {
			totalMaxToys = tempToys
		}
	}

	return int32(totalMaxToys)
}

func main() {
	prices := []int32{3, 7, 2, 9, 4}
	var k int32 = 15
	log.Println(NewMax(prices, k))
	//log.Println(maximumToys(prices, k))
}

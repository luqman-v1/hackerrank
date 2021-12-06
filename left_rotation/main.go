package main

import "log"

func main() {
	log.Println(rotLeft([]int32{1, 2, 3, 4, 5}, 4))
}

func rotLeft(a []int32, d int32) []int32 {
	var first []int32
	var last []int32
	n := len(a)
	for i := int(d); i < n; i++ {
		first = append(first, a[i])
	}
	for i := 0; i < int(d); i++ {
		last = append(last, a[i])
	}
	return append(first, last...)
}

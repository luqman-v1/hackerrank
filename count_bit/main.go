package main

import (
	"log"
	"strconv"
	"strings"
)

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
	bit := strconv.FormatInt(int64(num), 2)
	ce := strings.Split(bit, "")
	var result []int32
	for _, v := range ce {
		i, _ := strconv.Atoi(v)
		if i <= 0 {
			continue
		}
		result = append(result, int32(i))
	}
	return int32(len(result))
}

func main() {
	log.Println(countBits(10))
}

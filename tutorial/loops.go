package main

import (
	"fmt"
)

func loopAsetup() []int {
	array := make([]int, 10, 20)
	for i := 0; i < 10; i++ {
		array[i] = 1970 + i * 10
	}
	slice := array[:]
	return slice
}

func loopA(slice []int) {
	for i,v := range slice {
		fmt.Printf("[%d] %d\n", i, v)
	}
}


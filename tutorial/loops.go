package main

import (
	"fmt"
)

func loopAsetup() []int {
	// Make slice of int
	var s1 = []int{2, 4, 6, 8}
	fmt.Printf("Type %T, len %d, cap %d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 10)
	fmt.Printf("Type %T, len %d, cap %d\n", s1, len(s1), cap(s1))
	// Make slice of int
	s := make([]int, 10, 20)
	fmt.Printf("Type %T, len %d, cap %d\n", s, len(s), cap(s))
	for i := 0; i < 10; i++ {
		s[i] = 1970 + i * 10
	}
	slice := s[:]
	return slice
}

func loopA(slice []int) {
	for i,v := range slice {
		fmt.Printf("[%d] %d\n", i, v)
	}
	fmt.Printf("slice %T: len %d, cap %d\n", slice, len(slice), cap(slice))
}


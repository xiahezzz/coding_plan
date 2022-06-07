package main

import (
	"fmt"
	"math/rand"
	"time"

	"CodingPlan/coding"
)

func main() {
	/* 	s := user.Hello()
	   	fmt.Printf("s: %v\n", s) */
	/* 	fmt.Printf("coding.SelectionSort([]int{2, 3, 1}): %v\n", coding.InsertSort([]int{7, 2, 3, 1, 1, 3, 7, 0}))
	   	fmt.Println(coding.CalSumLR(1, 2)) */

	rand.Seed(time.Now().UnixNano())
	count := 0
	for i := 0; i < 100000000; i++ {
		if coding.RandChange() == 1 {
			count++
		}
	}
	fmt.Printf("count: %v\n", count)
	fmt.Printf("(count / 100000000): %v\n", float64(count)/100000000.0)

}

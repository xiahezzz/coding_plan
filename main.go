package main

import (
	"fmt"
	"math/rand"
	"time"

	"CodingPlan/coding"
	"CodingPlan/user"
)

func test(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] >= num {
			return i
		}
	}
	return -1
}

func main() {
	/* 	s := user.Hello()
	   	fmt.Printf("s: %v\n", s) */
	/* 	fmt.Printf("coding.SelectionSort([]int{2, 3, 1}): %v\n", coding.InsertSort([]int{7, 2, 3, 1, 1, 3, 7, 0}))
	   	fmt.Println(coding.CalSumLR(1, 2)) */

	rand.Seed(time.Now().UnixNano())
	/* count := 0
	for i := 0; i < 100000000; i++ {
		if coding.GetRandNum01f() == 1 {
			count++
		}
	}
	fmt.Printf("count: %v\n", count)
	fmt.Printf("(count / 100000000): %v\n", float64(count)/100000000.0) */
	test_times := 1
	maxValue := 1000
	maxLength := 100
	for i := 0; i < test_times; i++ {
		arr_use := coding.GetArrRandom(maxLength, maxValue)
		arr_sorted := make([]int, len(arr_use))
		copy(arr_sorted, arr_use)
		coding.InsertSort(arr_sorted)
		/* 		if !user.IsSortedArr(arr_sorted) {
			fmt.Println("arr_use:", arr_use)
			fmt.Println("arr_sorted:", arr_sorted)
			fmt.Println("bubble sort false")
			break
		} */
		rand_value := int(rand.Float64() * float64(maxValue))
		index, _ := coding.FindNumBLeft(arr_sorted, rand_value)
		if !(index == test(arr_sorted, rand_value)) {
			fmt.Println("Fucked!")
		}
		t := coding.FindMinIndex([]int{3, 2, 3, 2, 3})
		fmt.Printf("t: %v\n", t)
	}
	// var a *user.Node
	a := &user.Node{Data: 1, Next: nil}
	a.Next = &user.Node{Data: 2, Next: nil}
	a.Next.Next = &user.Node{Data: 3, Next: nil}

	a.ReverseList().PrintList()
	a.PrintList()

}

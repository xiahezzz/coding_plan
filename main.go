package main

import (
	"math/rand"
	"time"

	"CodingPlan/coding"
	"CodingPlan/user"
)

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
	test_times := 1000
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
		_, ok := coding.BinarySearch(arr_sorted, rand_value)
		if !(user.IsExistElem(arr_sorted, rand_value) == ok) {
			print("search false")
			break
		}
	}

}

package coding

import (
	"CodingPlan/user"
	"fmt"
)

func ReverseByGroup(list *user.Node, k int) *user.Node {
	length := list.Length()
	data := list.GetArrList()
	fmt.Printf("data: %v\n", data)

	newList := list

	for i := 0; i < length/k+1; i++ {
		temp := data[i*k:]
		if (i+1)*k <= length {
			temp = data[i*k : (i+1)*k]
		}
		BubbleSort(temp)
		for i2, j2 := 0, len(temp)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
			temp[i2], temp[j2] = temp[j2], temp[i2]
		}

		j := 0
		for list != nil {
			list.Data = temp[j]
			fmt.Printf("temp[j]: %v\n", temp[j])
			j++
			list = list.Next
			if j == k {
				break
			}
		}
	}
	return newList
}

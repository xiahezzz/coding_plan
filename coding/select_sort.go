package coding

func SelectionSort(arr []int) []int {
	length := len(arr)

	if arr == nil || length < 2 {
		return arr
	}

	for i := 0; i < length-1; i++ {
		min_value_index := i

		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min_value_index] {
				min_value_index = j
			}
		}

		arr[i], arr[min_value_index] = arr[min_value_index], arr[i]
	}

	return arr
}

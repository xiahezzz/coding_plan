package coding

func BinarySearch(arr []int, num int) (int, bool) {
	if len(arr) == 0 || arr == nil {
		return 0, false
	}

	start := 0
	end := len(arr) - 1
	mid := (start + end) / 2

	for start <= end {
		if arr[mid] == num {
			return mid, true
		}

		if arr[mid] > num {
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = (start + end) / 2
	}
	return 0, false
}

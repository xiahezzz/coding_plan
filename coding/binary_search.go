package coding

func BinarySearch(arr []int, num int) (int, bool) {
	if len(arr) == 0 || arr == nil {
		return -1, false
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
	return -1, false
}

func FindNumBLeft(arr []int, num int) (int, bool) {
	length := len(arr)
	if arr == nil || length == 0 {
		return -1, false
	}

	start := 0
	end := length - 1
	mid := (start + end) / 2
	index := -1

	for start <= end {
		if arr[mid] >= num {
			index = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}

	return index, true
}

func FindMinIndex(arr []int) int {
	//[]int 长度为0
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[length-2] > arr[length-1] {
		return length - 1
	}

	//剩下的情况只可能是，0-1递减，N-2到N-1递增，由于相邻项不相等0到N-1必定存在局部最小
	start := 0
	end := length - 1
	mid := (start + end) / 2
	for start < end-1 {
		if arr[mid-1] > arr[mid] && arr[mid] < arr[mid+1] {
			return mid
		}

		if arr[mid-1] < arr[mid] {
			start = mid + 1
		} else if arr[mid+1] < arr[mid] {
			end = mid - 1
		}

		mid = (start + end) / 2
	}

	if arr[start] < arr[end] {
		return start
	}
	return end
}

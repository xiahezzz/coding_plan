package user

func CompArr(arr1 []int, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func IsSortedArr(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < maxValue {
			return false
		}
		maxValue = arr[i]
	}
	return true
}

func IsExistElem(arr []int, num int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return true
		}
	}
	return false
}

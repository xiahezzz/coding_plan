package coding

func InsertSort(arr []int) {
	length := len(arr)
	if arr == nil || length <= 1 {
		return
	}

	for i := 1; i < length; i++ {

		for pre := i - 1; pre >= 0 && arr[pre+1] < arr[pre]; pre-- {
			temp := arr[pre+1]
			arr[pre+1] = arr[pre]
			arr[pre] = temp
		}
	}
}

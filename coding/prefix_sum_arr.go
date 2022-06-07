package coding

import "errors"

func CalSumLR(left int, right int) (int, error) {
	if left > right {
		return 0, errors.New("invalid input : left bigger than right")
	}

	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	pre_fix_sum_arr := make([]int, 7)
	pre_fix_sum_arr[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		pre_fix_sum_arr[i] = pre_fix_sum_arr[i-1] + arr[i]
	}

	if left == 0 {
		return pre_fix_sum_arr[right], nil
	}

	return pre_fix_sum_arr[right] - pre_fix_sum_arr[left-1], nil
}

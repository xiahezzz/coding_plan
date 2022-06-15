package coding

func TwoSum(nums []int, target int) []int {
	length := len(nums)
	numMap := make(map[int]int, length)
	for i := 0; i < len(nums); i++ {
		second := target - nums[i]
		if v, ok := numMap[second]; !ok {
			numMap[nums[i]] = i
		} else {
			return []int{v, i}
		}
	}
	return []int{0, 0}
}

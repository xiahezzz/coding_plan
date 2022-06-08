package coding

import "math/rand"

func GetArrRandom(maxLen int, maxValue int) []int {
	length := int(rand.Float32() * float32(maxLen))

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = int(rand.Float64() * float64(maxValue))
	}

	return result
}

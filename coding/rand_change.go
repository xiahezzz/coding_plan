package coding

import (
	"math/rand"
)

func getRandNum01() int {
	for {
		base := int(rand.Float64()*5) + 1
		switch base {
		case 1, 2:
			return 0
		case 4, 5:
			return 1
		case 3:
			continue
		}
	}
}

func RandChange() int {
	result := 0
	for {
		for i := 0; i < 3; i++ {
			result += (getRandNum01() << (2 - i))
		}
		if result != 0 {
			return result
		}
	}
}

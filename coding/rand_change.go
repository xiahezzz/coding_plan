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

func f() int {
	if rand.Float64() < 0.6 {
		return 0
	}
	return 1
}

func GetRandNum01f() int {
	result := f()
	for result == f() {
		result = f()
	}
	return result
	/* 	for {
		result = 0
		for i := 0; i < 2; i++ {
			result += (f() << (1 - i))
		}

		if result != 0 && result != 3 {
			return result - 1
		}
	} */
}

func RandChange() int {
	var result int
	for {
		result = 0
		for i := 0; i < 3; i++ {
			result += (getRandNum01() << (2 - i))
		}
		if result != 0 {
			return result
		}
	}
}

func RandChange2() int {
	var result int
	for {
		result = 0
		for i := 0; i < 6; i++ {
			result += (getRandNum01() << (5 - i))
		}
		if result <= 39 {
			return result + 17
		}
	}
}

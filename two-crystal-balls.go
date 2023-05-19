package algorithms

import "math"

func TwoCrystalBalls(floors *[]bool) int {

	low := 0
	high := len(*floors)
	step := int(math.Sqrt(float64(high)))

	for low < high {
		fl := low + step
		if !(*floors)[fl] {
			break
		} else {
			low = low + step
		}
	}

	for i := low; i < high; i++ {
		if (*floors)[i] {
			return low + i
		}
	}

	return -1

}

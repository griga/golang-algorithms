package algorithms

func BinarySearch(haystack *[]int, needle int) bool {

	l := 0
	h := len(*haystack)

	for l < h {
		m := l + (h-l)/2
		v := (*haystack)[m]
		if v < needle {
			l = m + 1
		} else if v > needle {
			h = m
		} else {
			return true
		}

	}
	return false
}

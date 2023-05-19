package algorithms

func BubbleSort(data *[]int) {

	iterations := 0
	swapFlag := true
	for swapFlag {
		swapFlag = false
		for i := 0; i < len(*data)-1-iterations; i++ {
			if (*data)[i] > (*data)[i+1] {
				(*data)[i], (*data)[i+1] = (*data)[i+1], (*data)[i]
				swapFlag = true
			}
		}
		iterations++
	}
}

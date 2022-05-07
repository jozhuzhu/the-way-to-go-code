package _slices

func MinSlice(items []int) (min int) {
	min = items[0]

	for _, item := range items {
		if min >= item {
			min = item
		}
	}
	return
}

func MaxSlice(items []int) (max int) {
	max = items[0]

	for _, item := range items {
		if max <= item {
			max = item
		}
	}

	return
}

package _slices

func Sum(items [5]float32) float32 {
	var sum float32 = 0.0
	for _, item := range items {
		sum += item
	}

	return sum
}

func SumAndAverage(items []int) (sum int, avg float32) {
	for _, item := range items {
		sum += item
	}

	avg = float32(sum / len(items))
	return
}

package math

// Average finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Min finds the minimum number in a series of numbers
func Min(xs []float64) float64 {
	var minimum float64
	for i, x := range xs {
		if i == 0 || x < minimum {
			minimum = x
		}
	}
	return minimum
}

// Max finds the largest number in a series of numbers
func Max(xs []float64) float64 {
	var max float64
	for i, x := range xs {
		if i == 0 || x > max {
			max = x
		}
	}
	return max
}

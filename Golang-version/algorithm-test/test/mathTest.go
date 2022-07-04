package test

func MathTest(ratio []int) []float64 {
	res := make([]float64, len(ratio))
	for i := 0; i < len(ratio); i++ {
		res[i] = float64(ratio[i]) / float64(ratio[0])
	}
	return res
}

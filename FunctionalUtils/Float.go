package FunctionalUtils

func MapFloat(f func(float64) float64, input []float64) (output []float64) {
	output = make([]float64, 0)
	for _, x := range input {
		output = append(output, f(x))
	}
	return
}

func FilterFloat(f func(float64) bool, input []float64) (output []float64) {
	output = make([]float64, 0)
	for _, x := range input {
		if f(x) {
			output = append(output, x)
		}
	}
	return
}

func ReduceFloat(f func(float64) float64, input []float64) (output float64) {
	for _, x := range input {
		output += f(x)
	}
	return
}
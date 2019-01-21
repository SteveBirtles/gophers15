package FunctionalUtils

func MapString(f func(string) string, input []string) (output []string) {
	output = make([]string, 0)
	for _, x := range input {
		output = append(output, f(x))
	}
	return
}

func FilterString(f func(string) bool, input []string) (output []string) {
	output = make([]string, 0)
	for _, x := range input {
		if f(x) {
			output = append(output, x)
		}
	}
	return
}

func ReduceString(f func(string) string, input []string) (output string) {
	for _, x := range input {
		output += f(x)
	}
	return
}
package FunctionalUtils

func MapInt(f func(int) int, input []int) (output []int) {
	output = make([]int, 0)
	for _, x := range input {
		output = append(output, f(x))
	}
	return
}

func FilterInt(f func(int) bool, input []int) (output []int) {
	output = make([]int, 0)
	for _, x := range input {
		if f(x) {
			output = append(output, x)
		}
	}
	return
}

func EveryInt(f func(int) bool, input []int) bool {
	for _, x := range input {
		if !f(x) { return false }
	}
	return true
}

func SomeInt(f func(int) bool, input []int) bool {
	for _, x := range input {
		if f(x) { return true }
	}
	return false
}

func RemoveInt(f func(int) bool, input []int) (output []int) {
	output = make([]int, 0)
	for _, x := range input {
		if !f(x) {
			output = append(output, x)
		}
	}
	return
}

func PartitionInt(f func(int) bool, input []int) (yes, no []int) {
	yes = make([]int, 0)
	no = make([]int, 0)
	for _, x := range input {
		if f(x) {
			yes = append(yes, x)
		} else {
			no = append(no, x)
		}
	}
	return
}

func SortInt(f func(int, int) bool, input []int) (output []int) {
	output = make([]int, len(input))
	copy(output, input)
	for i := 1; i < len(output); i++ {
		key := output[i]
		j := i - 1
		for j >= 0 && f(key, output[j]) {
			output[j+1] = output[j]
			j = j - 1
		}
		output[j+1] = key
	}
	return
}

func ReduceInt(f func(int) int, input []int) (output int) {
	for _, x := range input {
		output += f(x)
	}
	return
}

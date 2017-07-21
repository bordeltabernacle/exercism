package accumulate

const testVersion = 1

func Accumulate(input []string, f func(string) string) []string {
	for k, v := range input {
		input[k] = f(v)
	}
	return input
}

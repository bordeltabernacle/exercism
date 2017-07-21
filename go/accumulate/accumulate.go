package accumulate

const testVersion = 1

func Accumulate(input []string, f func(string) string) []string {
	output := make([]string, len(input))
	for k, v := range input {
		output[k] = f(v)
	}
	return output
}

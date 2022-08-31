package iteration

// Repeat takes in a character and returns a string by repeating it n times.
func Repeat(char string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += char
	}
	return
}

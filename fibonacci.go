package tour

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	n1 := -1
	n2 := 1
	// Fibonacci's sequence is (-1, 1), 0, 1, 1, 2, 3 ...
	// The n th value is sum of values of n-1 th and n-2 th.
	return func() int {
		v := n1 + n2
		n1 = n2
		n2 = v
		return v
	}
}

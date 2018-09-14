package tour

// Sqrt computes square root of the value by Newton' method.
// If the given value is not positive, then returns 0.
func Sqrt(x float64) float64 {
	if x <= 0 {
		return 0
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

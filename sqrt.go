package tour

import "fmt"

// ErrNegativeSqrt is used instead of complex number
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not compute square root of %v as real number.", float64(e))
}

// Sqrt computes square root of the value by Newton' method.
// If the given value is not positive, then returns 0.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return x, e
	}
	if x == 0 {
		return 0, nil
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

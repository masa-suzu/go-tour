package tour

// Pic creates a slice(len = dy) of slice(len = dx).
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
	}
	return pic
}

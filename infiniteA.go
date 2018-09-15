package tour

// InfiniteA is an infinite stream.
type InfiniteA struct{}

// Read always reads 'A'.
func (r InfiniteA) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

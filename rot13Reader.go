package tour

import "io"

type Rot13Reader struct {
	r io.Reader
}

func (rot13 *Rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	for i := 0; i < len(b); i++ {
		v := b[i]
		if 'a' <= v && v <= 'z' {
			b[i] = 'a' + (v-'a'+13)%26
		} else if 'A' <= v && v <= 'Z' {
			b[i] = 'A' + (v-'A'+13)%26
		}
	}
	return
}

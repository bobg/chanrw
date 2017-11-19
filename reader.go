package chanrw

import "io"

// Reader turns a []byte channel into an io.Reader.
type Reader struct {
	ch  <-chan []byte
	buf []byte
}

// NewReader produces a new Reader from a []byte channel.
func NewReader(ch <-chan []byte) *Reader {
	return &Reader{ch: ch}
}

// Read satisfies the io.Reader interface.
func (r *Reader) Read(out []byte) (int, error) {
	for len(r.buf) < len(out) {
		b, ok := <-r.ch
		if !ok {
			copy(out, r.buf)
			return len(r.buf), io.EOF
		}
		r.buf = append(r.buf, b...)
	}
	copy(out, r.buf)
	r.buf = r.buf[len(out):]
	return len(out), nil
}

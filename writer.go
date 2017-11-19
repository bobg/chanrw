package chanrw

// Writer turns a []byte channel into an io.WriteCloser.
type Writer struct {
	ch chan<- []byte
}

// NewWriter produces a new Writer from a []byte channel.
func NewWriter(ch chan<- []byte) *Writer {
	return &Writer{ch: ch}
}

// Write satisfies the io.Writer interface.
func (w *Writer) Write(inp []byte) (int, error) {
	w.ch <- inp
	return len(inp), nil
}

// Close satisfies the io.Closer interface.
func (w *Writer) Close() error {
	close(w.ch)
	return nil
}

package chanrw

import (
	"bytes"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	ch := make(chan []byte)
	go func() {
		ch <- []byte("foo")
		ch <- []byte("bar")
		close(ch)
	}()
	r := NewReader(ch)

	var b1 [1]byte
	n, err := r.Read(b1[:])

	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("got n == %d after first read, want 1", n)
	}
	if b1[0] != 'f' {
		t.Fatalf("got first character %c, want 'f'", b1[0])
	}

	var b2 [3]byte
	n, err = r.Read(b2[:])

	if err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("got n == %d after second read, want 3", n)
	}
	if !bytes.Equal(b2[:], []byte("oob")) {
		t.Fatalf("got \"%s\" after second read, want \"oob\"", string(b2[:]))
	}

	n, err = r.Read(b2[:])
	if err != io.EOF {
		t.Fatalf("got error %s after third read, want EOF", err)
	}
	if n != 2 {
		t.Fatalf("got n == %d after third read, want 2", n)
	}
	if !bytes.Equal(b2[:2], []byte("ar")) {
		t.Fatalf("got \"%s\" after third read, want \"ar\"", string(b2[:2]))
	}
}

func TestWriter(t *testing.T) {
	ch := make(chan []byte)
	w := NewWriter(ch)
	go func() {
		w.Write([]byte("f"))
		w.Write([]byte("oob"))
		w.Write([]byte("ar"))
		w.Close()
	}()
	var buf []byte
	for {
		b, ok := <-ch
		if !ok {
			if !bytes.Equal(buf, []byte("foobar")) {
				t.Fatalf("got \"%s\", want \"foobar\"", string(buf))
			}
			break
		}
		buf = append(buf, b...)
	}
}

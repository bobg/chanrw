# Chanrw

This Go package allows you to wrap a `chan []byte` as an `io.Reader`
or as an `io.WriteCloser`.

## Example

Suppose `produceBytes(ch chan<- []byte)` runs as a goroutine and
produces bytes that you wish to consume line-by-line. You can wrap the
reading end of the channel as a `chanrw.Reader` in order to use it in
`bufio.NewScanner`.

```
import (
  "bufio"
  "chanrw"
)

func lineByLine() {
  ch := make(chan []byte)
  go produceBytes(ch)
  r := chanrw.NewReader(ch)
  s := bufio.NewScanner(r)
  for s.Scan() {
    // ...consume a line from produceBytes as s.Text()...
  }
}
```

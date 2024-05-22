// In Go, errors are values. That means they are not special and you can program
// them. Here is one programming technique for avoiding repetitive error
// handling. Adapted from https://go.dev/blog/errors-are-values.
package main

import (
	"fmt"
	"io"
	"os"
)

type errReader struct {
	r   io.Reader
	err error
}

// read becomes a no-op as soon as an error
// occurs but the error value gets saved
func (er *errReader) read(buf []byte) {
	if er.err != nil {
		return
	}
	_, er.err = er.r.Read(buf)
}

func main() {
	filename := "/etc/passwd"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "errgo: %v\n", err)
		os.Exit(1)
	}

	buf := make([]byte, 9)

	er := &errReader{r: f}
	er.read(buf[0:3]) // We do not
	er.read(buf[3:6]) // handle error
	er.read(buf[6:9]) // for each call.
	if er.err != nil {
		fmt.Fprintf(os.Stderr, "errgo: reading %s: %v\n", filename, er.err)
		os.Exit(1)
	}
}

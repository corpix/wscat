package websocket

import (
	"io"
)

type ReadWriteCloser struct {
	*Reader
	*Writer
	io.Closer
}

func NewReadWriteCloser(r *Reader, w *Writer, c io.Closer) ReadWriteCloser {
	return ReadWriteCloser{
		Reader: r,
		Writer: w,
		Closer: c,
	}
}

package websocket

import (
	"io"
	"sync"

	"github.com/gorilla/websocket"
)

type Reader struct {
	*websocket.Conn
	current io.Reader
	lock    *sync.Mutex
}

func (w *Reader) Read(buf []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	var (
		t   int
		r   io.Reader
		n   int
		err error
	)

	if w.current != nil {
		goto read
	}

	for {
		t, r, err = w.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		switch t {
		case websocket.TextMessage:
		case websocket.BinaryMessage:
		default:
			// XXX: Unsupported message
			continue
		}

		break
	}

	w.current = r
read:
	n, err = w.current.Read(buf)
	if err == io.EOF {
		w.current = nil
	}

	return n, err
}

func NewReader(c *websocket.Conn) *Reader {
	return &Reader{
		Conn: c,
		lock: &sync.Mutex{},
	}
}

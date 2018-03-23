package websocket

import (
	"github.com/gorilla/websocket"
)

type Writer struct {
	*websocket.Conn
}

func (w *Writer) Write(buf []byte) (int, error) {
	return len(buf), w.Conn.WriteMessage(websocket.BinaryMessage, buf)
}

func NewWriter(c *websocket.Conn) *Writer {
	return &Writer{Conn: c}
}

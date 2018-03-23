package websocket

import (
	"io"
	"net/http"
)

type Handler interface {
	ServeWebsocket(w io.ReadWriteCloser, r *http.Request)
}

type HandlerFunc func(rwc io.ReadWriteCloser, r *http.Request)

func (hf HandlerFunc) ServeWebsocket(rwc io.ReadWriteCloser, r *http.Request) {
	hf(rwc, r)
}

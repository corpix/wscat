package websocket

import (
	"net/http"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"
	"github.com/gorilla/websocket"
)

var (
	DefaultHeaders = http.Header{}
)

// HTTPUpgradeHandler is a net/http.Handler which is responsible
// for incomming HTTP request upgrade process and handling the
// upgraded request with some Handler which works only with
// websockets.
type HTTPUpgradeHandler struct {
	Handler
	log loggers.Logger
}

func (h *HTTPUpgradeHandler) ServeHTTP(rwc http.ResponseWriter, r *http.Request) {
	r.Body.Close()

	var (
		c   *websocket.Conn
		err error
	)

	c, err = websocket.Upgrade(rwc, r, DefaultHeaders, 1024, 1024)
	if err != nil {
		h.log.Error(err)
		return
	}

	h.ServeWebsocket(
		NewReadWriteCloser(
			NewReader(c),
			NewWriter(c),
			c,
		),
		r,
	)
}

func NewHTTPUpgradeHandler(h Handler, l loggers.Logger) *HTTPUpgradeHandler {
	return &HTTPUpgradeHandler{
		Handler: h,
		log: prefixwrapper.New(
			"HTTPUpgradeHandler: ",
			l,
		),
	}
}

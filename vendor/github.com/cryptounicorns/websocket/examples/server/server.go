package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/logrus"
	"github.com/corpix/loggers/logger/prefixwrapper"
	"github.com/gorilla/mux"
	logrusLogger "github.com/sirupsen/logrus"

	"github.com/cryptounicorns/websocket"
)

type Message struct {
	Text   string
	Number int
}

func NewMessage(text string, num int) *Message {
	return &Message{
		Text:   text,
		Number: num,
	}
}

type Handler struct {
	log loggers.Logger
}

func (h Handler) logError(err error) {
	h.log.Errorf("%T: %s", err, err)
}

func (h Handler) handleError(err error, rwc io.ReadWriteCloser) bool {
	if err == nil {
		return false
	}

	// XXX: Because gorilla's error is private we can't
	// type assert. But we can apply this "hack" to obtain type name
	// still better than strings.HasPrefix and others in my opinion.
	switch fmt.Sprintf("%T", err) {
	case "*net.OpError":
	case "*websocket.netError":
	case "*websocket.CloseError":
	case "*errors.errorString":
		if !strings.HasSuffix(err.Error(), "close sent") {
			h.logError(err)
		}
	default:
		h.logError(err)
	}

	h.log.Debugf(
		"Terminating loop for %#v with %T: %#v",
		rwc,
		err,
		err,
	)

	return true
}

func (h Handler) ServeWebsocket(rwc io.ReadWriteCloser, r *http.Request) {
	defer rwc.Close()
	var (
		wg = &sync.WaitGroup{}
	)

	go func() {
		defer wg.Done()
		var (
			b   []byte
			err error
		)
		for {
			b, err = ioutil.ReadAll(rwc)
			if h.handleError(err, rwc) {
				return
			}

			h.log.Printf("Received %d: %s", len(b), b)
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		var (
			m   []byte
			n   = 0
			err error
		)
		for {
			m, err = json.Marshal(NewMessage("hello", n))
			if h.handleError(err, rwc) {
				return
			}

			_, err = rwc.Write(m)
			if h.handleError(err, rwc) {
				return
			}

			n++
			time.Sleep(10 * time.Millisecond)
		}
	}()
	wg.Add(1)

	wg.Wait()
}

func Mount(r *mux.Router, l loggers.Logger) {
	var (
		log = prefixwrapper.New("GET /: ", l)
	)
	r.Handle(
		"/",
		websocket.NewHTTPUpgradeHandler(
			Handler{log: log},
			log,
		),
	)
}

func main() {
	var (
		lg = logrusLogger.New()
	)

	lg.Level = logrusLogger.DebugLevel

	var (
		r    = mux.NewRouter()
		l    = logrus.New(lg)
		addr = "127.0.0.1:3333"
	)
	Mount(r, l)

	l.Print("Listening ", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		l.Fatal(err)
	}

	select {}
}

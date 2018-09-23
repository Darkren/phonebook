// Package wsjsonrpc contains all the stuff used to pass
// JSONRPC requests through gorilla's websocket implementation
package wsjsonrpc

import (
	"io"

	"github.com/gorilla/websocket"
)

// ReadWriteCloser is an implementation of standard ReadWriteCloser interface.
// It encapsulates gorilla's websocket and provides Read, Write and Close
// methods as wrappers above websocket. The instance of this type is to be used
// in methods provided by net/rpc package
type ReadWriteCloser struct {
	WS *websocket.Conn
	r  io.Reader
	w  io.WriteCloser
}

// Read reads up to len(p) bytes from websocket into p. Returns
// the number of bytes read and error if any
func (rwc *ReadWriteCloser) Read(p []byte) (int, error) {
	if rwc.r == nil {
		var err error
		_, rwc.r, err = rwc.WS.NextReader()

		if err != nil {
			return 0, err
		}
	}

	var n int
	for n < len(p) {
		m, err := rwc.r.Read(p[n:])

		n += m

		if err == io.EOF {
			rwc.r = nil

			return n, err
		}

		if err != nil {
			return n, err
		}
	}

	return n, nil
}

// Write writes bytes from p into websocket. Returns the number
// of bytes written and error if any
func (rwc *ReadWriteCloser) Write(p []byte) (int, error) {
	if rwc.w == nil {
		var err error
		rwc.w, err = rwc.WS.NextWriter(websocket.TextMessage)

		if err != nil {
			return 0, err
		}
	}

	var n int
	for n < len(p) {
		m, err := rwc.w.Write(p)

		n += m

		if err != nil {
			rwc.Close()

			return n, err
		}
	}

	if n == len(p) {
		rwc.Close()
	}

	return n, nil
}

// Close closes websocket connection
func (rwc *ReadWriteCloser) Close() error {
	if rwc.w != nil {
		err := rwc.w.Close()
		rwc.w = nil

		if err != nil {
			return err
		}
	}

	return nil
}

package server

import (
	"io"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

const bufferSize = 32 * 1024

var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, bufferSize)
	},
}

func copyTCPToWS(ws *websocket.Conn, tcp net.Conn) {

	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf)

	for {
		n, err := tcp.Read(buf)
		if err != nil {
			return
		}

		if err := ws.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
			return
		}
	}
}

func copyWSToTCP(ws *websocket.Conn, tcp net.Conn) {

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			return
		}

		if _, err := io.Copy(tcp, bytesReader(data)); err != nil {
			return
		}
	}
}

func bytesReader(b []byte) io.Reader {
	return &sliceReader{b: b}
}

type sliceReader struct {
	b []byte
}

func (r *sliceReader) Read(p []byte) (int, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}

	n := copy(p, r.b)
	r.b = r.b[n:]
	return n, nil
}

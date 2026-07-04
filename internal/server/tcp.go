package server

import (
	"net"

	"github.com/gorilla/websocket"
)

// dialBackend membuka koneksi ke backend SSH.
func dialBackend(addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}

// copyTCPToWS meneruskan data dari SSH -> WebSocket.
func copyTCPToWS(ws *websocket.Conn, tcp net.Conn) {
	buf := make([]byte, 32*1024)

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

// copyWSToTCP meneruskan data dari WebSocket -> SSH.
func copyWSToTCP(ws *websocket.Conn, tcp net.Conn) {
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			return
		}

		if _, err := tcp.Write(data); err != nil {
			return
		}
	}
}

// relay menjalankan proxy dua arah.
func relay(ws *websocket.Conn, tcp net.Conn) {
	done := make(chan struct{}, 2)

	go func() {
		copyTCPToWS(ws, tcp)
		done <- struct{}{}
	}()

	go func() {
		copyWSToTCP(ws, tcp)
		done <- struct{}{}
	}()

	<-done
}

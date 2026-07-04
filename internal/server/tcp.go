package server

import (
	"net"

	"github.com/gorilla/websocket"
)

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

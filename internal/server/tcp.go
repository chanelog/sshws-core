package server

import (
	"io"
	"net"

	"github.com/coder/websocket"
)

func copyTCPToWS(ws *websocket.Conn, tcp net.Conn) error {

	buf := make([]byte, 32*1024)

	for {

		n, err := tcp.Read(buf)
		if err != nil {
			return err
		}

		err = ws.Write(
			ws.Context(),
			websocket.MessageBinary,
			buf[:n],
		)

		if err != nil {
			return err
		}

	}

}

func copyWSToTCP(ws *websocket.Conn, tcp net.Conn) error {

	for {

		_, data, err := ws.Read(ws.Context())
		if err != nil {
			return err
		}

		_, err = tcp.Write(data)
		if err != nil {
			return err
		}

	}

}

func relay(ws *websocket.Conn, tcp net.Conn) {

	done := make(chan error, 2)

	go func() {
		done <- copyTCPToWS(ws, tcp)
	}()

	go func() {
		done <- copyWSToTCP(ws, tcp)
	}()

	<-done
}

package server

import (
	"net/http"

	"github.com/chanelog/sshws-core/internal/logger"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {

	// Upgrade HTTP -> WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer ws.Close()

	logger.Info.Printf("New client: %s\n", r.RemoteAddr)

	// Connect ke backend SSH
	tcp, err := dialBackend(s.cfg.Backend)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer tcp.Close()

	// Jalankan proxy dua arah
	done := make(chan struct{}, 2)

	go func() {
		copyTCPToWS(ws, tcp)
		done <- struct{}{}
	}()

	go func() {
		copyWSToTCP(ws, tcp)
		done <- struct{}{}
	}()

	// Tunggu salah satu koneksi selesai
	<-done
}

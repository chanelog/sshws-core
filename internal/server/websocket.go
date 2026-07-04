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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer conn.Close()

	logger.Info.Println("New WebSocket:", r.RemoteAddr)

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		if err := conn.WriteMessage(mt, msg); err != nil {
			break
		}
	}
}

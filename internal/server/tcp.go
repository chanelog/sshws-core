package server

import (
	"net"
	"time"
)

func dialBackend(addr string) (net.Conn, error) {

	conn, err := net.DialTimeout(
		"tcp",
		addr,
		10*time.Second,
	)
	if err != nil {
		return nil, err
	}

	if tcp, ok := conn.(*net.TCPConn); ok {
		tcp.SetKeepAlive(true)
		tcp.SetKeepAlivePeriod(30 * time.Second)
		tcp.SetNoDelay(true)
	}

	return conn, nil
}

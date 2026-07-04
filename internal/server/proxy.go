package server

import (
	"net"

	"github.com/chanelog/sshws-core/internal/config"
	"github.com/chanelog/sshws-core/internal/logger"
)

type Proxy struct {
	cfg *config.Config
}

func NewProxy(cfg *config.Config) *Proxy {
	return &Proxy{
		cfg: cfg,
	}
}

func (p *Proxy) Dial() (net.Conn, error) {

	logger.Info.Printf("Connecting to backend %s\n", p.cfg.Backend)

	conn, err := net.Dial("tcp", p.cfg.Backend)
	if err != nil {
		logger.Error.Println(err)
		return nil, err
	}

	logger.Info.Println("Backend connected")

	return conn, nil
}

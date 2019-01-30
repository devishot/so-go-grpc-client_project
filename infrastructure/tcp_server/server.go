package tcp_server

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	Cfg Config
}

func (s *TCPServer) Listen() net.Listener {
	addr := fmt.Sprintf(":%d", s.Cfg.Port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("tcp_server Listen: error=%v", err)
		return nil
	}

	return lis
}

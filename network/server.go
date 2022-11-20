package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts ServerOpts
	rpcCh      chan RPC
	quitCh     chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC, 1),
		quitCh:     make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransport()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%s\n", rpc)
		case <-ticker.C:
			fmt.Println("tick x seconds")
		case <-s.quitCh:
			break free
		}
	}

	fmt.Println("Server stopped")
}

func (s *Server) initTransport() {
	for _, tr := range s.ServerOpts.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}

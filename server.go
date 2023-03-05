package main

import (
	"fmt"
	"github.com/shu2k/meme-cache/cache"
	"log"
	"net"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts
	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}

	fmt.Printf("server startingg on port [%s]\n", s.ListenAddr)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error %s\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(connection net.Conn) {
	defer func() {
		connection.Close()
	}()
	buffer := make([]byte, 2048)
	for {
		n, err := connection.Read(buffer)

		if err != nil {
			log.Printf("connection read error: %s\n", err)
			break
		}

		msg := buffer[:n]
		fmt.Println(string(msg))
	}
}

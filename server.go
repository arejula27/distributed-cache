package main

import (
	"log"
	"net"
)

// Server handles the cache connections
type Server struct {
	cache      Cache
	listenAddr string
}

func (s *Server) handleConnection(conn net.Conn) {
	//It is possibe to read and write from the connection
	defer conn.Close()
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		log.Printf("error reading buffer: %s\n", err)
		return
	}
	log.Println("Message received ", string(buff))
}

func (s *Server) Start() error {

	listen, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer listen.Close()
	log.Printf("server starting on port [%s]\n", s.listenAddr)
	for {
		conn, err := listen.Accept()
		log.Println("Connection accepted ", conn.RemoteAddr())
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// NewServer creates a server
func NewServer() *Server {
	server := Server{
		listenAddr: "127.0.0.1:8000",
		cache:      NewMemoryCache(),
	}
	return &server
}

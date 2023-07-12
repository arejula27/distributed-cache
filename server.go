package main

import (
	"io"
	"log"
	"net"

	"github.com/arejula27/distributed-cache/protocol"
)

// Server handles the cache connections
type Server struct {
	cache      Cache
	listenAddr string
}

func (s *Server) handleConnection(conn net.Conn) {
	//It is possibe to read and write from the connection
	defer conn.Close()
	exit := false
	for !exit {

		rqt, err := protocol.ParseAction(conn)
		if err != nil {

			switch err {
			case io.EOF:
				log.Println("connection closed by client")
				exit = true
			default:
				log.Println("parse command error:", err)
			}

		}

		switch action := rqt.(type) {
		case *protocol.GetRequest:
			response, err := s.handleGetAction(action)
			if err != nil {
				errResponse := protocol.ErrorResponse{Error: err}
				conn.Write(errResponse.Serialize())
				continue
			}
			conn.Write(response.Serialize())
		case *protocol.SetRequest:
			s.handleSetAction(action)
		case *protocol.ExitRequest:
			exit = true
		}
	}

}

func (s *Server) handleGetAction(getRqt *protocol.GetRequest) (protocol.GetResponse, error) {
	value, err := s.cache.Get(getRqt.Key())
	if err != nil {
		return protocol.GetResponse{}, err
	}
	return protocol.GetResponse{Value: value}, nil

}

func (s *Server) handleSetAction(setRqt *protocol.SetRequest) {
	log.Println("Handling SET. key: ", setRqt.Key())
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

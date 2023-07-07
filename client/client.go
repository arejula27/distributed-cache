package main

import (
	"log"
	"net"

	"github.com/arejula27/distributed-cache/protocol"
)

const TYPE = "tcp"

func main() {

	tcpServer, err := net.ResolveTCPAddr(TYPE, "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("Resolving tcp addres failed:", err.Error())

	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Fatalln("Dial failed:", err.Error())

	}
	log.Println("connection succes")
	msg := protocol.CreateGetRequest([]byte("hola"))
	_, err = conn.Write(msg.Serialize())
	if err != nil {
		log.Fatalln("Write data failed:", err.Error())

	}
	log.Println("message sended")
}

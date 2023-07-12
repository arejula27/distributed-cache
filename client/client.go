package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

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
	reader := bufio.NewReader(os.Stdin)

	exit := false
	for !exit {
		fmt.Print("Enter command: ")
		rawAction, _ := reader.ReadString('\n')
		action := strings.Replace(rawAction, "\n", "", 1)
		words := strings.Split(action, " ")

		switch command := words[0]; command {
		case "GET":

			msg := protocol.CreateGetRequest([]byte("hola"))
			_, err = conn.Write(msg.Serialize())
			if err != nil {
				log.Fatalln("Write data failed:", err.Error())

			}
			response, err := protocol.ParseResponse(conn)
			if err != nil {
				log.Println("Error while reading response: ", err)
			}
			switch action := response.(type) {
			case *protocol.GetResponse:
				log.Println("Value:", string(action.Value))
			case *protocol.ErrorResponse:
				log.Println("Error:", action.Error.Error())

			}

		case "EXIT":
			exit = true
		default:
			log.Println("Unkown command: ", command)

		}

	}

}

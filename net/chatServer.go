package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	portServer = flag.Int("pServer", 3003, "specifies the port")
	hostServer = flag.String("hServer", "localhost", "host to listen on")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)
	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the chat server, your name is: %s\n", clientName)
	messages <- fmt.Sprintf("New client is here, its name is: %s\n", clientName)
	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}
	leavingClients <- message
	messages <- fmt.Sprintf("%s said goodby\n", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

/*
Esta funcion va a manejar el como se conectar entre clientes
*/
func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages: //Aca reparte el sms de alguien a atodos
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients: //Agrega el nuevo cliente a la lista
			clients[newClient] = true
		case leavingCLient := <-leavingClients:
			delete(clients, leavingCLient)
			close(leavingCLient)
		}
	}
}

func main() {
	//net.Dial es para conectarse
	//net.Listen para escuchar
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *hostServer, *portServer))
	if err != nil {
		log.Fatal(err)
	}
	//Esta funcion es independiente, por eso no hay wg
	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}

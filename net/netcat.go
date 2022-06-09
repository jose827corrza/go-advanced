package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3003, "specifies the port")
	host = flag.String("h", "localhost", "host to listen on")
)

/*
Este codigo vendria siendo el del cliente
*/
func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected..")
	done := make(chan struct{})

	//Start goroutine to receive data from the server
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	// Copy what we got to the console line
	CopyContent(conn, os.Stdin)
	<-done //bloquea todo
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}
}

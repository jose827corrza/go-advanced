package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "host to be scanned")

func main() {
	flag.Parse() //esto toca llamarlo antes de usar el flag en el codigo
	//SIN CONCURRENCIA
	//for port := 0; port < 100; port++ {
	//	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
	//	if err != nil {
	//		continue
	//	}
	//	conn.Close()
	//	fmt.Printf("The port %d, is open\n", port)
	//}

	//CON CONCURRENCIA
	var wg sync.WaitGroup
	for i := 0; i < 65000; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("The port %d, is open\n", port)
		}(i)
	}
	wg.Wait()
}

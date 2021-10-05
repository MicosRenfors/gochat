package main

import (
	"log"
	"net"
)

func main() {
	// Initializes new server and starts goroutine for run() function
	s := newServer()
	go s.run()

	// Starts tcp server listening for incoming connections
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on :8888")

	// Starts goroutine for every new connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}

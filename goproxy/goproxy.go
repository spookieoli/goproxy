package goproxy

import (
	"flag"
	"fmt"
	"io"
	"net"
)

// Proxy Main type of the package
type Proxy struct {
	localAddr, remoteAddr string
}

// NewProxy Constructor for Proxy
func NewProxy() *Proxy {
	localAddr := flag.String("localAddr", "0.0.0.0:8080", "Local address")
	remoteaddr := flag.String("remoteAddr", "127.0.0.1:8080", "Remote address")
	flag.Parse()
	return &Proxy{*localAddr, *remoteaddr}
}

// handleConnection handles an incoming connection from a client. It establishes a connection
// with the remote server, and then copies the data between the client and the server.
// The connection with the client is closed when the method execution finishes.
func (p *Proxy) handleConnection(conn net.Conn) {
	defer conn.Close()

	// connect to the server
	server, err := net.Dial("tcp", p.remoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	// copy data between client and server
	go io.Copy(server, conn)
	io.Copy(conn, server)
}

// Serve Function to start the proxy
func (p *Proxy) Serve() {
	fmt.Println("Starting proxy on", p.localAddr, "forwarding to", p.remoteAddr)
	listener, err := net.Listen("tcp", p.localAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go p.handleConnection(conn)
	}
}

package src

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr     string
	ln             net.Listener
	quit           chan struct{}
	messageChannel chan []byte
	connectedPeers map[string]net.Conn
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr:     listenAddr,
		quit:           make(chan struct{}),
		messageChannel: make(chan []byte, 100),
		connectedPeers: make(map[string]net.Conn),
	}
}

func (s *Server) AllMessages() chan []byte {
	return s.messageChannel
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quit

	close(s.messageChannel)

	return nil
}

func (s *Server) Stop() {
	close(s.quit)
}

func (s *Server) acceptLoop() {

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}
		log.Println("New connection to the server from : ", conn.RemoteAddr())
		s.connectedPeers[conn.RemoteAddr().String()] = conn
		go s.readLoop(conn)
	}
	// accept
}

func (s *Server) readLoop(conn net.Conn) {
	// read
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error: ", err)
			continue
		}
		s.messageChannel <- buf[:n]
		fmt.Println("total connections are ", len(s.connectedPeers))

	}
}

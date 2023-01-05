package main

import (
	"fmt"
	"log"
	"net"
)


type Server struct {
	listnAddr string
	ln net.Listener
	quit chan struct{}
	msgch chan []byte
    peer  map[string]net.Conn

}


func NewServer(listnAddr string) *Server {
	return &Server{
		listnAddr: listnAddr,
		quit: make(chan struct{}),
		msgch: make(chan []byte ,100),
		peer: make(map[string]net.Conn),
	}
}


func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listnAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quit

	close(s.msgch)

	
	return nil
}

func (s *Server) Stop() {
	close(s.quit)
}

func(s *Server) acceptLoop() {

	for {
		conn , err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}
		log.Println("New connection to the server from : ", conn.RemoteAddr())
		s.peer[conn.RemoteAddr().String()] = conn
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
		s.msgch <- buf[:n]
		fmt.Println("total connections are " , len(s.peer))
		
	}
}


func main() {
	srv := NewServer(":3000")

	go func() {
		for msg := range srv.msgch {
			fmt.Println(string(msg))
		}
		
	}()
	log.Fatal(srv.Start())
}
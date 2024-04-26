package server

import (
	"io"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener,
	}, nil
}

func (s *Server) Serve() error {
	log.Printf("Serving on %v", s.listener.Addr())

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		log.Printf("Accepted connection")

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 256)

		if _, err := conn.Read(buf); err == io.EOF {
			log.Print("Closing connection")
			return
		} else if err != nil {
			log.Printf("Failed to read from socket: %v", err)
			continue
		}

		log.Print("Received message")

		go s.handleMessage(conn, buf)
	}
}

func (s *Server) handleMessage(conn net.Conn, msg []byte) {
	log.Printf("%v", string(msg))

	conn.Write([]byte("+PONG\r\n"))
}

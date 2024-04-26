package server

import (
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/resp/command"
	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
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
		// TODO bigger buffer
		buf := make([]byte, 1024)

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
	cmd, err := command.UnmarshalBinary(msg)
	if err != nil {
		log.Printf("Failed unmarshalling command: %v", err)
		return
	}

	var response []byte
	switch actualCmd := cmd.(type) {
	case *command.Echo:
		response, err = s.handleEcho(actualCmd)
	case *command.Ping:
		response, err = s.handlePing()
	}

	if err != nil {
		log.Printf("Failed to handle command: %v", err)
	} else {
		conn.Write(response)
	}
}

// TODO make more generic (return Data, then marshal)
func (s *Server) handleEcho(echo *command.Echo) ([]byte, error) {
	resp := data.NewBulkStringWithData(echo.Data())

	return resp.MarshalBinary()
}

func (s *Server) handlePing() ([]byte, error) {
	resp := data.NewSimpleStringWithData("PONG")

	return resp.MarshalBinary()
}

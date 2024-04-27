package server

import (
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/resp/command"
	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

type Server struct {
	listener        net.Listener
	store           *store.Store
	commandChannel  chan command.Command
	responseChannel chan data.Data
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener:        listener,
		store:           store.NewStore(),
		commandChannel:  make(chan command.Command),
		responseChannel: make(chan data.Data),
	}, nil
}

func (s *Server) Serve() error {
	log.Printf("Serving on %v", s.listener.Addr())

	go s.eventLoop()

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

	s.commandChannel <- cmd

	response := <-s.responseChannel

	if responseBinary, err := response.MarshalBinary(); err != nil {
		log.Printf("Failed to handle command: %v", err)
	} else {
		conn.Write(responseBinary)
	}
}

func (s *Server) eventLoop() {
	for {
		cmd := <-s.commandChannel

		log.Printf("Handling event %+v", cmd)

		var err error
		var response data.Data
		switch actualCmd := cmd.(type) {
		case *command.Echo:
			response, err = s.handleEcho(actualCmd)
		case *command.Ping:
			response, err = s.handlePing()
		case *command.Set:
			response, err = s.handleSet(actualCmd)
		case *command.Get:
			response, err = s.handleGet(actualCmd)
		}

		if err != nil {
			log.Printf("Failed handling command: %v", err)
		}

		s.responseChannel <- response
	}
}

func (s *Server) handleEcho(echo *command.Echo) (data.Data, error) {
	return data.NewBulkStringWithData(echo.Data()), nil
}

func (s *Server) handlePing() (data.Data, error) {
	return data.NewSimpleStringWithData("PONG"), nil
}

func (s *Server) handleSet(set *command.Set) (data.Data, error) {
	s.store.Set(set.Key(), set.Value())

	return data.NewSimpleStringWithData("OK"), nil
}

func (s *Server) handleGet(get *command.Get) (data.Data, error) {
	if value, present := s.store.Get(get.Key()); !present {
		return data.NewNullBulkString(), nil
	} else {
		return data.NewBulkStringWithData(value), nil
	}
}

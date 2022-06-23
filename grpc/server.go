package grpc

import (
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Port string
}

func New(port string) *Server {
	return &Server{
		Server: grpc.NewServer(),
		Port:   port,
	}
}

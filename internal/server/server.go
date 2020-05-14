package server

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"net"

	pb "github.com/cloudingcity/grpc-chat/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	password string
}

func New(password string) *Server {
	return &Server{
		password: password,
	}
}

func (s *Server) Listen(port int) {
	log.Infof("Password: %s", s.password)
	addr := fmt.Sprintf(":%d", port)
	log.Infof("Started listening on %s\n", addr)

	lis, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	pb.RegisterChatServer(server, &Chat{password: s.password})
	if err := server.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type Chat struct {
	password string
}

func (c *Chat) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Password != c.password {
		return nil, errors.New("invalid password")
	}
	return &pb.LoginResponse{
		Token: token(),
	}, nil
}

func token() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

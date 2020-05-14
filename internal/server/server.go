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

func Listen(port int, password string) {
	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Infof("Password: %s", password)
	log.Infof("Started listening on %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterChatServer(server, &Server{password: password})
	if err := server.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type Server struct {
	password string
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Password != s.password {
		return nil, errors.New("invalid password")
	}
	log.Infof("[%s] is logged in", req.Username)
	return &pb.LoginResponse{
		Token: token(),
	}, nil
}

func token() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

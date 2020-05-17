package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	pb "github.com/cloudingcity/grpc-chat/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Listen(port int, password string) {
	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Infof("Started listening on %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterChatServer(server, newServer(password))
	if err := server.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type Server struct {
	password string
	manager  *UserManager
}

func newServer(password string) *Server {
	return &Server{
		password: password,
		manager:  NewUserManager(),
	}
}

func (s *Server) Connect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	if req.Password != s.password {
		return nil, errors.New("invalid password")
	}
	log.Infof("[%s] is logged in", req.Username)

	tkn := genToken()
	s.manager.Register(tkn, req.Username)

	return &pb.ConnectResponse{
		Token: string(tkn),
	}, nil
}

func (s *Server) Stream(stream pb.Chat_StreamServer) error {
	tkn, err := s.getToken(stream)
	if err != nil {
		return err
	}
	user, err := s.manager.Get(tkn)
	if err != nil {
		return err
	}
	user.Stream = stream

	go func() {
		for {
			req, err := stream.Recv()
			if err != nil {
				s.manager.Deregister(tkn)
				return
			}
			resp := &pb.StreamResponse{
				Username: req.Username,
				Message:  req.Message,
			}
			s.manager.Broadcast(resp)
		}
	}()

	select {
	case <-stream.Context().Done():
	case <-user.Done:
	}
	log.Infof("[%s] is logged out", user.Name)
	return nil
}

func (s *Server) getToken(stream pb.Chat_StreamServer) (token, error) {
	md, _ := metadata.FromIncomingContext(stream.Context())
	tkn, ok := md["x-token"]
	if !ok {
		return "", errors.New("token not found")
	}
	return token(tkn[0]), nil
}

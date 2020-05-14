package client

import (
	"context"

	pb "github.com/cloudingcity/grpc-chat/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Client struct {
	username string
	password string
}

func New(username, password string) *Client {
	return &Client{
		username: username,
		password: password,
	}
}

func (c *Client) Connect(addr string) {
	log.Infof("UserName: %s", c.username)
	log.Infof("Password: %s", c.password)
	log.Infof("Connect to server: %s", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	chat := &Chat{
		client: pb.NewChatClient(conn),
	}
	token, err := chat.Login(c.username, c.password)
	if err != nil {
		log.Fatalln(err)
	}
	log.Infof("Token: %s", token)
}

type Chat struct {
	client pb.ChatClient
}

func (c *Chat) Login(username, password string) (string, error) {
	req := &pb.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := c.client.Login(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

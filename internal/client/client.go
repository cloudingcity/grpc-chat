package client

import (
	log "github.com/sirupsen/logrus"
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
}

package server

import (
	"errors"
	"sync"

	pb "github.com/cloudingcity/grpc-chat/proto"
)

type User struct {
	Token  token
	Name   string
	Stream pb.Chat_StreamServer
	Done   chan struct{}
}

type UserManager struct {
	users map[token]*User
	mux   sync.RWMutex
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[token]*User),
	}
}

func (m *UserManager) Users() map[token]*User {
	return m.users
}

func (m *UserManager) Register(tkn token, username string) {
	m.mux.Lock()
	m.users[tkn] = &User{
		Token: tkn,
		Name:  username,
		Done:  make(chan struct{}),
	}
	m.mux.Unlock()
}

func (m *UserManager) Get(tkn token) (*User, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if user, ok := m.users[tkn]; ok {
		return user, nil
	}
	return nil, errors.New("token not valid")
}

func (m *UserManager) Deregister(tkn token) {
	m.mux.Lock()
	close(m.users[tkn].Done)
	delete(m.users, tkn)
	m.mux.Unlock()
}

func (m *UserManager) Broadcast(resp *pb.StreamResponse) {
	m.mux.Lock()
	defer m.mux.Unlock()

	for _, user := range m.users {
		if user.Stream == nil {
			continue
		}
		if err := user.Stream.Send(resp); err != nil {
			m.Deregister(user.Token)
			continue
		}
	}
}

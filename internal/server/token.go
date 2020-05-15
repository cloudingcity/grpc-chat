package server

import (
	"crypto/rand"
	"fmt"
)

type token string

func genToken() token {
	b := make([]byte, 4)
	rand.Read(b)
	return token(fmt.Sprintf("%x", b))
}

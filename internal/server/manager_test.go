package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestManager(t *testing.T) {
	const tkn1, tkn2 = "token 1", "token 2"
	m := NewUserManager()

	m.Register(tkn1, "Foo")
	m.Register(tkn2, "Bar")
	require.Equal(t, 2, len(m.Users()))

	_, err := m.Get("not exists token")
	require.Error(t, err)

	user, err := m.Get(tkn1)
	require.NoError(t, err)
	require.Equal(t, "Foo", user.Name)

	m.Deregister(tkn1)
	require.Equal(t, 1, len(m.users))
}

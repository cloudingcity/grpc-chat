package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenToken(t *testing.T) {
	t1 := genToken()
	t2 := genToken()
	assert.NotEqual(t, t1, t2)
}

package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalStorage(t *testing.T) {
	tra := NewLocalTransport("a")
	trb := NewLocalTransport("b")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.Addr()], trb)
	assert.Equal(t, trb.peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("a")
	trb := NewLocalTransport("b")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello")
	tra.SendMessage(trb.Addr(), msg)

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.From, tra.Addr())
	assert.Equal(t, rpc.payload, msg)
}

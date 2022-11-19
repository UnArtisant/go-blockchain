package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalStorage(t *testing.T) {
	tra := NewLocalStorage("a")
	trb := NewLocalStorage("b")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalStorage("a")
	trb := NewLocalStorage("b")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello")
	tra.sendMessage(trb.addr, msg)

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.From, tra.addr)
	assert.Equal(t, rpc.payload, msg)
}

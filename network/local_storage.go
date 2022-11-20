package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	peers     map[NetAddr]*LocalTransport
	consumeCh chan RPC
	lock      sync.RWMutex
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		peers:     make(map[NetAddr]*LocalTransport),
		consumeCh: make(chan RPC, 1024),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]

	if !ok {
		return fmt.Errorf("%s: no peer with address %s", t.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    t.addr,
		payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

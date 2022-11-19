package network

import (
	"fmt"
	"sync"
)

type LocalStorage struct {
	addr      NetAddr
	peers     map[NetAddr]*LocalStorage
	consumeCh chan RPC
	lock      sync.RWMutex
}

func NewLocalStorage(addr NetAddr) *LocalStorage {
	return &LocalStorage{
		addr:      addr,
		peers:     make(map[NetAddr]*LocalStorage),
		consumeCh: make(chan RPC, 1024),
	}
}

func (t *LocalStorage) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalStorage) Connect(tr *LocalStorage) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

func (t *LocalStorage) sendMessage(to NetAddr, payload []byte) error {
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

func (t *LocalStorage) Addr() NetAddr {
	return t.addr
}

package main

import (
	"blockchain/network"
	"time"
)

func main() {

	trLocal := network.NewLocalTransport("local")
	trRemote := network.NewLocalTransport("remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello"))
			time.Sleep(1 * time.Second)
		}
	}()

	opt := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}
	s := network.NewServer(opt)
	s.Start()
}

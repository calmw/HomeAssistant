package internal

import (
	"net"
	"sync"
)

type Client struct {
	ClientId             string
	Ip                   string
	Conn                 *net.Conn
	LastActiveTime       string
	HeartbeatFailedTime  int64
	HeartbeatFailedTimes int
}

var Clients []Client
var clientMx = &sync.Mutex{}

func AddClient(cli Client) {
	var exist bool
	for _, client := range Clients {
		if cli.ClientId == client.ClientId {
			exist = true
		}
	}
	if !exist {
		clientMx.Lock()
		defer clientMx.Unlock()
		Clients = append(Clients, cli)
	}
}

func RemoveClient(cli Client) {
	for i, client := range Clients {
		if cli.ClientId == client.ClientId {
			clientMx.Lock()
			defer clientMx.Unlock()
			Clients = append(Clients[:i], Clients[i+1:]...)
			return
		}
	}
}

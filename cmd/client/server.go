package main

import (
	"encoding/json"
	"home_assistant/internal"
	"home_assistant/log"
	"io"
	"net"
	"time"
)

func main() {

	var conn net.Conn

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Logger.Sugar().Fatal("socket listen error", err)
	}
	defer listen.Close()

	go func() {
		ticker := time.NewTicker(internal.HeartbeatDuration * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			for _, client := range internal.Clients {
				heartbeat(client)
			}

		}
	}()
	
	for {
		conn, err = listen.Accept()
		if err != nil {
			log.Logger.Sugar().Errorf("socket accept error", err)
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Logger.Sugar().Debugf("the connection is closed")
			} else {
				log.Logger.Sugar().Errorf("read error: %s", err)
			}
			return
		}
		var receive internal.Receive
		err = json.Unmarshal(buf[:n], &receive)
		if err != nil {
			log.Logger.Sugar().Errorf("json unmarshal error: %s", err)
			return
		}
		
		internal.AddClient(internal.Client{
			ClientId:       receive.ClientId,
			Ip:             conn.RemoteAddr().String(),
			Conn:           &conn,
			LastActiveTime: time.Now().Format("2006-01-02 15:04:05"),
		})

		switch receive. {
		
		}
	}
}

func heartbeat(cli internal.Client) {
	_, err := (*cli.Conn).Write([]byte("ping"))
	if err != nil {
		log.Logger.Sugar().Errorf("heartbeat write error: %s", err)
		return
	}

	if (time.Now().Unix() - cli.HeartbeatFailedTime) >= internal.HeartbeatDuration*2 {
		cli.HeartbeatFailedTimes += 1
	}

	if cli.HeartbeatFailedTimes >= 3 {
		internal.RemoveClient(cli)
		return
	}
}

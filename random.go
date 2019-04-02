package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

func SendRandomData(client *websocket.Conn) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		w, err := client.NextWriter(websocket.TextMessage)
		if err != nil {
			ticker.Stop()
			break
		}

		msg := newRandomData()
		w.Write(msg)
		w.Close()

		<-ticker.C
	}
}

func newRandomData() []byte {
	data, _ := json.Marshal(map[string]int{
		"x": rand.Intn(10),
	})
	return data
}

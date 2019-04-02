package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		go SendRandomData(ws)
	})

	http.HandleFunc("/chat", handleConnections)
	go handleMessages()

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Received: %s\n", message)
		if err = conn.WriteMessage(messageType, []byte("Reply :"+string(message))); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", echo)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

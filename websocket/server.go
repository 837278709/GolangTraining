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
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Input big or small")
}

func reader(conn *websocket.Conn) {
	for {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 7
		rollNO := rand.Intn(max-min) + min
		var result string
		if rollNO >= 4 {
			result = "big"
		} else {
			result = "small"
		}
		fmt.Println(rollNO, result)

		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var rtMsg string
		if result == string(p) {
			rtMsg = "True"
		} else {
			rtMsg = "False"
		}

		switch string(p) {
		case "big":
		case "small":
		default:
			rtMsg = "Input big or small"
			log.Println("incorret input: ", string(p))
		}

		if err := conn.WriteMessage(messageType, []byte(rtMsg)); err != nil {
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected")

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

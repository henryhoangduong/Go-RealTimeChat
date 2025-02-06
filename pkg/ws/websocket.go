package ws

import "github.com/gorilla/websocket"

type Client struct {
}

type Message struct {
}

var clients = make(map[*Client]bool)
var broadcast = make(chan *model.Chat)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}
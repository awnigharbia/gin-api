package notifications

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "chrome-extension://cbcbkhdmedgianpaifchdaddpnmgnknn"
	},
}
var clients = make(map[*websocket.Conn]bool)

func NotificationsListener(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		clients[conn] = true

		for client := range clients {
			if client != conn {
				err := client.WriteMessage(websocket.TextMessage, []byte("New user joined"))
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				delete(clients, conn)
				break
			}
			for client := range clients {
				err := client.WriteMessage(messageType, append([]byte("hello"), message...))
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
}
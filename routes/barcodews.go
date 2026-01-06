package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)

func BarcodeReaderWS(c *gin.Context) {
	// Setup and upgrade websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// Clean Up connection
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	// Pingpong to check if websocket is connected
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Setup a ticker to send a ping to the webpage every 30 seconds
	pingTicker := time.NewTicker(10 * time.Second)
	defer pingTicker.Stop()

	clients[conn] = true

	for {
		for range pingTicker.C {
			err := conn.WriteControl(
				websocket.PingMessage,
				[]byte{},
				time.Now().Add(10*time.Second),
			)
			if err != nil {
				return
			}
		}
	}
}

func SendWSBarcodes(barcode string) {
	if barcode != "" {
		for clientConn, _ := range clients {
			clientConn.WriteMessage(websocket.TextMessage, []byte(barcode))
		}
	}
}

package websocket

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})
	if err != nil {
		log.Printf("WebSocket accept error: %v", err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "connection closed")

	log.Println("New WebSocket connection established")

	ctx, cancel := context.WithTimeout(r.Context(), time.Hour)
	defer cancel()
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				log.Println("Client closed connection")
			} else {
				log.Printf("Read error: %v", err)
			}
			break
		}

		log.Printf("Received: %s", message)

		err = conn.Write(ctx, websocket.MessageText, []byte("Echo: "+string(message)))
		if err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

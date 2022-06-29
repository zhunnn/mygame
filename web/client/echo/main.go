package main

import (
	"mygame/internal/logrot"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:6600/echo", nil)
	if err != nil {
		logrot.Log.Panicln("[Dial]:", err)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		// Write
		err = c.WriteMessage(websocket.TextMessage, []byte("I'm client."))
		if err != nil {
			logrot.Log.Errorln("[Write]:", err)
			return
		}
		// Read
		_, msg, err := c.ReadMessage()
		if err != nil {
			logrot.Log.Errorln("[Read]:", err)
			return
		}
		logrot.Log.Infoln("Receive:", string(msg))
		time.Sleep(time.Second)
	}
}

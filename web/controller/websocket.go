package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HeartBeat(c *gin.Context) {
	// 升級協議
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("[HeartBeat 升級協議失敗]:", err)
		return
	}
	defer ws.Close()

	for {
		// 讀取 ws 中的數據
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("[HeartBeat 讀取數據失敗]:", err)
			break
		}
		if string(msg) == "heart" {
			msg = []byte("beat")
		}
		// 寫入 ws 數據
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			log.Println("[HeartBeat 寫入數據失敗]:", err)
			return
		}
	}
}

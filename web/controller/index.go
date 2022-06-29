package controller

import (
	"mygame/internal/logrot"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	Index    index
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type index struct{}

func (i *index) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (i *index) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Server alive!")
}

func (i *index) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}

func (i *index) Greet(c *gin.Context) {
	logrot.Log.Debugln("Receive:", c.Query("msg"))
	if c.Query("msg") == "Hello" {
		c.JSON(http.StatusOK, "World!")
	} else {
		c.JSON(http.StatusOK, "Where is your Hello?")
	}
}

func (i *index) Echo(c *gin.Context) {
	// 升級協議
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrot.Log.Errorln("[Echo 升級協議失敗]:", err)
		return
	}
	defer func() {
		ws.Close()
	}()

	for {
		// 讀取 ws 中的數據
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			logrot.Log.Errorln("[Echo 讀取數據失敗]:", err)
			break
		}
		logrot.Log.Infoln("Receive:", string(msg))
		msg = append([]byte("Echo: "), msg...)
		// 寫入 ws 數據
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			logrot.Log.Errorln("[Echo 寫入數據失敗]:", err)
			return
		}
	}
}

package controller

import (
	"mygame/internal/logrot"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	Game game
)

type game struct{}

func (g *game) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "game.html", nil)
}

// 遊戲連線
func (g *game) Connect(c *gin.Context) {
	// 升級協議
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrot.Log.Errorln("[Connect 升級協議失敗]:", err)
		return
	}
	defer func() {
		defer logrot.Log.Warnln("Goroutine 連線結束")
		conn.Close()
	}()

	// // Get Request Header
	// account := c.Request.Header.Get("Account")
	// password := c.Request.Header.Get("Password")
	// logrot.Log.Infoln("Account:", account)
	// logrot.Log.Infoln("Password:", password)

	// Login
	for {
		err = conn.WriteMessage(websocket.TextMessage, []byte("請輸入帳號"))
		if err != nil {
			logrot.Log.Errorln("[Connect 寫入數據失敗]:", err)
			return
		}
		_, account, err := conn.ReadMessage()
		if err != nil {
			logrot.Log.Errorln("[Connect 讀取數據失敗]:", err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, []byte("請輸入密碼"))
		if err != nil {
			logrot.Log.Errorln("[Connect 寫入數據失敗]:", err)
			return
		}
		_, password, err := conn.ReadMessage()
		if err != nil {
			logrot.Log.Errorln("[Connect 讀取數據失敗]:", err)
			return
		}
		logrot.Log.Infof("登入資訊 Account: %v, Password: %v", string(account), string(password))
		if string(account) == "nick" && string(password) == "1234" {
			logrot.Log.Infoln("登入成功")
			err = conn.WriteMessage(websocket.TextMessage, []byte("登入成功!"))
			if err != nil {
				logrot.Log.Errorln("[Connect 寫入數據失敗]:", err)
				return
			}
			break
		} else {
			logrot.Log.Infoln("帳號密碼錯誤")
			err = conn.WriteMessage(websocket.TextMessage, []byte("帳號密碼錯誤!"))
			if err != nil {
				logrot.Log.Errorln("[Connect 寫入數據失敗]:", err)
				return
			}
		}
	}

	// Write
	go func() {
		defer logrot.Log.Warnln("Goroutine 寫入賽果結束")
		for {
			// 寫入 ws 數據
			err = conn.WriteMessage(websocket.TextMessage, []byte("賽果"))
			if err != nil {
				logrot.Log.Errorln("[Connect 寫入數據失敗]:", err)
				return
			}
			time.Sleep(time.Second)
		}
	}()
	// Read
	for {
		// 讀取 ws 中的數據
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logrot.Log.Errorln("[Connect 讀取數據失敗]:", err)
			break
		}
		logrot.Log.Infoln("Receive:", string(msg))
		// 輸出 ws 數據
		msg = []byte("先不要插嘴")
		logrot.Log.Infoln("Response:", string(msg))
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logrot.Log.Errorln("[Connect 寫入數據失敗]:", err)
			break
		}
	}
}

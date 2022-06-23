package main

import (
	"encoding/json"
	"mygame/config"
	"mygame/internal/logrot"
	"mygame/web/router"
)

func main() {
	conf, _ := json.Marshal(config.Config)
	logrot.Info("啟動服務 設定檔:", string(conf))
	router := router.New(config.Config.System.Environment)
	router.Init()
	router.Start(config.Config.Server.Http.Port)
}

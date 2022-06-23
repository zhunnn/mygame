package main

import (
	"mygame/config"
	"mygame/internal/logrot"
	"mygame/web/router"
)

func main() {
	// Start
	logrot.Warn("啟動服務: ", config.Config.System.ServiceName)
	logrot.Warn("服務設定: ", config.Config.Print())
	// HTTP Router
	router := router.New(config.Config.System.Environment)
	router.Init()
	router.Start(config.Config.Server.Http.Port)
}

package main

import (
	"mygame/config"
	"mygame/internal/logrot"
	"mygame/web/router"
)

func main() {
	// Start
	logrot.Log.Infoln("開啟服務:", config.Config.System.ServiceName)
	logrot.Log.Info("服務設定:\n", config.Config.Print())
	// HTTP Router
	router := router.New(config.Config.System.Environment)
	router.Init()
	router.Start(config.Config.Server.Http.Port)
}

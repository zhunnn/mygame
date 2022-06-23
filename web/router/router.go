package router

import (
	"mygame/config"
	"mygame/config/enum"
	"mygame/internal/logrot"
	"mygame/web/controller"
	"mygame/web/middleware"
	"net"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func New(env string) *Router {
	gin.DefaultWriter = logrot.StandardLogger().Out
	// Set gin mode
	switch env {
	case enum.Environment_Local:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	// New
	return &Router{
		Engine: gin.New(),
	}
}

func (r *Router) Init() {
	logrot.Info("Router Init...")
	// Middleware
	r.Use(middleware.GetMiddleware()...)
	// Html
	r.LoadHTMLGlob(config.Config.System.ProjectRootPath + "/web/view/template/*")
	r.Static("/static", config.Config.System.ProjectRootPath+"/web/view/static")
	// r.StaticFile("/favicon.ico", config.Config.System.ProjectRootPath+"/favicon.ico")
	// Register
	r.Register()
	// Proxy
	addr, err := net.LookupHost("LocalHost")
	if err != nil {
		logrot.Error("[查詢 Proxy 錯誤]:", err)
	}
	err = r.Engine.SetTrustedProxies(addr)
	if err != nil {
		logrot.Error("[設定 Proxy 錯誤]:", err)
	}
}

func (r *Router) Start(port string) {
	logrot.Info("Router Start on ", port)
	err := r.Run(port)
	if err != nil {
		logrot.Fatal("[Router 啟動錯誤]:", err)
	}
}

func (r *Router) Register() {
	// Websocket 心跳長連線
	r.GET("/heartbeat", controller.HeartBeat)

	// 檢查
	r.GET("/", controller.Index)
	r.GET("/index", controller.Index)
	r.GET("/health", controller.HealthCheck)
	r.GET("/ping", controller.Ping)
	r.GET("/message", controller.Message)

	// 設定路由群組
	user := r.Group("/user")
	user.GET("/index", controller.User.Index)

	// 設定路由群組
	game := r.Group("/game")
	game.GET("/index", controller.Game.Index)

	// 設定路由群組
	post := r.Group("/post")
	post.GET("/index", controller.Post.Index)
}

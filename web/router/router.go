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
	gin.DefaultWriter = logrot.Log.MultiWriter
	// Set gin mode
	switch env {
	case enum.Environment_Local:
		// gin.SetMode(gin.DebugMode)
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	// New
	return &Router{
		Engine: gin.New(),
	}
}

func (r *Router) Init() {
	logrot.Log.Infoln("Router Init...")
	// Middleware
	r.Use(middleware.GetMiddleware()...)
	// Html
	r.LoadHTMLGlob(config.Config.System.ProjectRootPath + "/web/view/template/*")
	r.Static("/static", config.Config.System.ProjectRootPath+"/web/view/static")
	// Register
	r.Register()
	// Proxy
	addr, err := net.LookupHost("LocalHost")
	if err != nil {
		logrot.Log.Errorln("[查詢 Proxy 錯誤]:", err)
	}
	err = r.Engine.SetTrustedProxies(addr)
	if err != nil {
		logrot.Log.Errorln("[設定 Proxy 錯誤]:", err)
	}
}

func (r *Router) Start(port string) {
	logrot.Log.Infoln("Router Serving on", port)
	err := r.Run(port)
	if err != nil {
		logrot.Log.Panicln("[Router 啟動錯誤]:", err)
	}
}

func (r *Router) Register() {
	// 基礎測試
	r.GET("/", controller.Index.Home)
	r.GET("/health", controller.Index.HealthCheck)
	r.GET("/ping", controller.Index.Ping)
	r.GET("/greet", controller.Index.Greet)
	// 長連線
	r.GET("/echo", controller.Index.Echo)

	// 設定路由群組
	user := r.Group("/user")
	user.GET("/", controller.User.Index)

	// 設定路由群組
	game := r.Group("/game")
	game.GET("/", controller.Game.Index)
	game.GET("/connect", controller.Game.Connect)

	// 設定路由群組
	post := r.Group("/post")
	post.GET("/", controller.Post.Index)
}

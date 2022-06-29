package middleware

import (
	"mygame/internal/logrot"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GetMiddleware 統一註冊中間件接口
func GetMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		Cors(),
		Recovery(),
		Logger(),
	}
}

// Cors 誇域設定
func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("POST, OPTIONS, GET, PUT, DELETE")
	corsConfig.AddAllowHeaders("Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
	return cors.New(corsConfig)
}

// Recovery 捕捉 panic
func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}

// Logger 設定
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 開始時間
		startTime := time.Now()
		// 處理請求
		c.Next()
		// 結束時間
		endTime := time.Now()
		// 執行時間
		latencyTime := endTime.Sub(startTime)
		// 請求方式
		reqMethod := c.Request.Method
		// 請求路由
		reqUri := c.Request.RequestURI
		// 狀態碼
		statusCode := c.Writer.Status()
		// 請求 IP
		clientIP := c.ClientIP()
		// 依照狀態碼區分日誌等級
		switch statusCode / 100 {
		case 1:
			logrot.Log.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		case 2:
			logrot.Log.Debugf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		case 3:
			logrot.Log.Warnf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		case 4:
			logrot.Log.Errorf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		case 5:
			logrot.Log.Errorf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		default:
			logrot.Log.Warnf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
		}
	}
}

// AllowOriginSpecfic 允許特定對象誇域設定
func AllowOriginSpecfic(allowOrigins []string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("POST, OPTIONS, GET, PUT, DELETE")
	corsConfig.AddAllowHeaders("Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
	return cors.New(corsConfig)
}

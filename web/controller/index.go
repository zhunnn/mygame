package controller

import (
	"mygame/internal/logrot"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Server alive!")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}

func Message(c *gin.Context) {
	logrot.Debug("Receive: ", c.Query("msg"))
	if c.Query("msg") == "Hello" {
		c.JSON(http.StatusOK, "World!")
	} else {
		c.JSON(http.StatusOK, "Where is your Hello?")
	}
}

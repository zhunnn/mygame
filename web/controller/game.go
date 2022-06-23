package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Game game
)

type game struct{}

func (g *game) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "game.html", nil)
}

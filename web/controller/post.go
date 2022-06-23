package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Post post
)

type post struct{}

func (p *post) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", nil)
}

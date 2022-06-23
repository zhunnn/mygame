package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	User user
)

type user struct{}

func (u *user) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", nil)
}

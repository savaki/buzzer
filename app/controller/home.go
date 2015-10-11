package controller

import (
	"io"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	io.WriteString(c.Writer, "hello world")
}

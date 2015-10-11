package controller

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	content := TwiML{
		Say: "hello world",
	}
	c.Header("Content-Type", "text/xml")
	xml.NewEncoder(c.Writer).Encode(content)
}

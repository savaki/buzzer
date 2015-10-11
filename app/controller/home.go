package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	content := TwiML{
		Say: "hello world",
	}
	c.Header("Content-Type", "text/xml")
	json.NewEncoder(c.Writer).Encode(content)
}

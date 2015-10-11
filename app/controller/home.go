package controller

import (
	"encoding/xml"
	"io"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	content := TwiML{
		Say: "hello world",
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

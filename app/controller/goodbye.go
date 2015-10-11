package controller

import (
	"encoding/xml"
	"io"

	"github.com/gin-gonic/gin"
)

func Goodbye(c *gin.Context) {
	content := TwiML{
		Say: "good bye",
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

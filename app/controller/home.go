package controller

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	for key, values := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
	}
	defer c.Request.Body.Close()
	io.Copy(os.Stdout, c.Request.Body)

	content := TwiML{
		Say: "hello world",
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

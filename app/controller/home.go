package controller

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	input := Form{}
	if err := c.Bind(&input); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	json.NewEncoder(os.Stdout).Encode(input)

	content := TwiML{
		Say: "hello world",
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

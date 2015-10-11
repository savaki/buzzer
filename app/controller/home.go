package controller

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
)

var whitelist = whiteListFromEnv()

func Home(c *gin.Context) {
	input := Form{}
	if err := c.Bind(&input); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(whitelist)
	content := TwiML{
		Say: "hello world",
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

func whiteListFromEnv() []string {
	phones := []string{}
	for _, env := range os.Environ() {
		parts := strings.Split(env, "=")
		if len(parts) != 2 {
			continue
		}

		if !strings.HasPrefix(parts[0], "PH_") {
			continue
		}

		phones = append(phones, parts[1])
	}
	return phones
}

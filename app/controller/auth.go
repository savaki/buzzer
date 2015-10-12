package controller

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var whitelist = whiteListFromEnv()

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

func Auth(c *gin.Context) {
	input := struct {
		Digits  string
		Attempt int
	}{}

	if err := c.Bind(&input); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Received:", input.Digits)

	for _, validCode := range whitelist {
		if input.Digits == validCode {
			Buzz(c)
			return
		}
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/?Attempt=%d", input.Attempt+1))
}

func Buzz(c *gin.Context) {
	content := TwiML{
		Say: "enter",
		Play: &Play{
			Digits: "#9",
		},
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

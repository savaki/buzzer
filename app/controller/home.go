package controller

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	input := Form{}
	if err := c.Bind(&input); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if input.Attempt >= 3 {
		c.Redirect(http.StatusFound, "/goodbye")
		return
	}

	content := TwiML{
		Gather: &Gather{
			Action: fmt.Sprintf("/auth?Attempt=%d", input.Attempt),
			Method: "POST",
			Say:    "Please enter your authorization code",
		},
		Redirect: &Redirect{
			Method: "GET",
			Action: fmt.Sprintf("/?Attempt=%d", input.Attempt+1),
		},
	}
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>`)
	xml.NewEncoder(c.Writer).Encode(content)
}

func Dial(c *gin.Context) {
	c.Header("Content-Type", "text/xml")

	io.WriteString(c.Writer, `<?xml version="1.0" encoding="UTF-8" ?>
<Response>
	<Dial timeout="10">510-529-9511</Dial>
</Response>`)
}

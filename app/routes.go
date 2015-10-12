package app

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/savaki/buzzer/app/controller"
)

func Routes() http.Handler {
	router := gin.New()
	router.GET("/", controller.Home)
	router.POST("/", controller.Home)

	router.GET("/auth", controller.Auth)
	router.POST("/auth", controller.Auth)

	router.GET("/buzz", controller.Buzz)
	router.GET("/goodbye", controller.Goodbye)

	fs, err := filepath.Abs(".")
	if err != nil {
		log.Fatalln(err)
	}
	audioContent := func(c *gin.Context) {
		c.Header("Content-Type", "audio/basic")

		filename, err := filepath.Abs(filepath.Join(fs, c.Request.URL.Path))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		log.Println(filename)

		if _, err := os.Stat(filename); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		log.Println("a")

		// attempted hack, ignore this
		if !strings.HasPrefix(filename, fs) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		log.Println("b")

		f, err := os.Open(filename)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		defer f.Close()

		io.Copy(c.Writer, f)
	}
	router.GET("/static/*path", audioContent)

	return router
}

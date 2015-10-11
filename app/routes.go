package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savaki/buzzer/app/controller"
)

func Routes() http.Handler {
	router := gin.New()
	router.GET("/", controller.Home)

	return router
}

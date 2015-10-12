package app

import (
	"net/http"

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

	router.Static("/static", "static")

	return router
}

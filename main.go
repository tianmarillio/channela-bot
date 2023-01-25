package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tianmarillio/channela-backend/config"
	"github.com/tianmarillio/channela-backend/src/controllers"
	"github.com/tianmarillio/channela-backend/src/middlewares"
	"github.com/tianmarillio/channela-backend/src/scheduler"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()
	scheduler.Start()

	r := gin.Default()
	config.SetupCors(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/login", controllers.LoginController.Login)

	channelRouter := r.Group("/channels", middlewares.Authenticate)
	channelRouter.POST("", controllers.ChannelController.Create)
	channelRouter.GET("", controllers.ChannelController.FindAll)
	channelRouter.GET("/search-youtube", controllers.ChannelController.SearchYoutubeChannel)
	channelRouter.GET("/:channelId", controllers.ChannelController.FindById)
	channelRouter.PATCH("/:channelId", controllers.ChannelController.Update)
	channelRouter.DELETE("/:channelId", controllers.ChannelController.Delete)

	r.Run(os.Getenv("PORT"))
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/perfectogo/upload/api/handlers"
	"github.com/perfectogo/upload/config"
	"github.com/perfectogo/upload/service"
)

type Options struct {
	Config  config.Config
	Service service.InterfaceServer
}

func Runner(options Options) {
	handlers := handlers.NewHandler(handlers.CfgHandler{
		Config:  options.Config,
		Service: options.Service,
	})

	router := gin.Default()
	router.POST("/upload", handlers.UploadImg)
	router.GET("/get", handlers.GetImages)
	router.Run(options.Config.HTTPPort)
}

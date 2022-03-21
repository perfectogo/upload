package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/perfectogo/upload/api/handlers"
	"github.com/perfectogo/upload/config"
	"github.com/perfectogo/upload/pkg/logger"
	"github.com/perfectogo/upload/service"
)

type Options struct {
	Config  config.Config
	Service service.InterfaceServer
	Log     logger.Logger
}

func Runner(options Options) error {
	handlers := handlers.NewHandler(handlers.CfgHandler{
		Config:  options.Config,
		Service: options.Service,
		Log:     options.Log,
	})

	router := gin.Default()
	router.POST("/upload", handlers.UploadImg)
	router.GET("/get", handlers.GetImages)
	if err := router.Run(options.Config.HTTPPort); err != nil {
		return err
	}
	return nil
}

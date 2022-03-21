package main

import (
	"github.com/perfectogo/upload/api/routes"
	"github.com/perfectogo/upload/config"
	"github.com/perfectogo/upload/pkg/database"
	"github.com/perfectogo/upload/pkg/logger"
	"github.com/perfectogo/upload/service"
	"github.com/perfectogo/upload/storage"
)

func main() {
	cfg, _ := config.Load()
	log := logger.New(cfg.LogLevel, "crud_app")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger ", logger.Error(err))
		}
	}(log)

	sqlxDB, err := database.ConnectionToDB(cfg)
	defer sqlxDB.Close()
	if err != nil {
		log.Fatal("db connection error", logger.Error(err))
		return
	}
	storage := storage.NewStoragePg(sqlxDB)
	service := service.NewService(storage)

	if err := routes.Runner(routes.Options{
		Config:  cfg,
		Log:     log,
		Service: service,
	}); err != nil {
		log.Fatal(err.Error())
	}

}

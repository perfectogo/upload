package main

import (
	"github.com/perfectogo/upload/api/routes"
	"github.com/perfectogo/upload/config"
	"github.com/perfectogo/upload/pkg/database"
	"github.com/perfectogo/upload/service"
	"github.com/perfectogo/upload/storage"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, _ := config.Load()
	sqlxDB, err := database.ConnectionToDB(cfg)
	defer sqlxDB.Close()
	if err != nil {
		logrus.Fatalf("db connection error: %s", err.Error())
		return
	}
	storage := storage.NewStoragePg(sqlxDB)
	service := service.NewService(storage)

	routes.Runner(routes.Options{
		Config:  cfg,
		Service: service,
	})

}

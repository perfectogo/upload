package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	HTTPPort         string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
}

// Load loads environment vars and inflates Config
func Load() (Config, error) {
	cfg := Config{}

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg.HTTPPort = viper.GetString("port")
	cfg.PostgresHost = viper.GetString("db.host")
	cfg.PostgresPort = viper.GetInt("db.port")
	cfg.PostgresDatabase = viper.GetString("db.dbname")
	cfg.PostgresUser = viper.GetString("db.username")
	cfg.PostgresPassword = os.Getenv("DB_PASSWORD")
	cfg.LogLevel = viper.GetString("db.logLevel")
	return cfg, nil
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

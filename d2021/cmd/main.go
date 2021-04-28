package main

import (
	"github.com/YRIDZE/DIPLOMA2021/main"
	"github.com/YRIDZE/DIPLOMA2021/main/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title HRM API
// @version 1.0
// @description API Server for HRM system

// @host localhost:8001
// @BasePath /path

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := handler.NewHandler()
	srv := new(d2021.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

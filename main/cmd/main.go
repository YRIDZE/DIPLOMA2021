package main

import (
	"encoding/json"
	"fmt"
	api "github.com/YRIDZE/DIPLOMA2021/main/api"
	data "github.com/YRIDZE/DIPLOMA2021/main/data"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig(configPaths string) {
	file, err := os.Open(configPaths)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data.Config)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

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

	router := api.InitRoutes()
	router.Run(":8001")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

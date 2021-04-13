package main

import (
	"encoding/json"
	"fmt"
	api "github.com/YRIDZE/DIPLOMA2021/main/api"
	"os"
)

// @title HRM API
// @version 1.0
// @description API Server for HRM system

// @host localhost:8001
// @BasePath /path

func LoadConfig(configPaths string) {
	file, err := os.Open(configPaths)
	if err != nil {
		fmt.Println("error1:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&api.Config)
	if err != nil {
		fmt.Println("error2: ", err)
	}
}

func main() {
	LoadConfig("main/configuration.json")
	router := api.Routes()
	router.Run(":8001")
}

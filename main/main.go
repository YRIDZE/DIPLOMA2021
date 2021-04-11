package main

import (
	api "github.com/YRIDZE/DIPLOMA2021/main/api"
)

// @title HRM API
// @version 1.0
// @description API Server for HRM system

// @host localhost:8001
// @BasePath /path

func main() {

	router := api.Routes()
	router.Run(":8001")
}

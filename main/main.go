package main

import (
	"github.com/YRIDZE/DIPLOMA2021/main/api"
)

func main() {

	router := api.Routes()
	router.Run(":8080")
}

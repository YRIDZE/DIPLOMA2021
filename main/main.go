package main

import (
	"encoding/json"
	"fmt"
	api "github.com/YRIDZE/DIPLOMA2021/main/api"
	"os"
)

var data = []byte(`{
    "employee": { 
        "idEmployee": 2,
        "firstName": "John",
        "lastName": "Doll",
        "email": "JohnDoll@gmail.com",
        "password": "123454321",
        "country": "USA"
    },
    "course": [  
        {
            "idCourse": 2,
            "title": "C#",
            "status": "suggested"
        },
        {
            "idCourse": 4,
            "title": "C",
            "status": "in progress"
        }
    ]
}`)

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

// @title HRM API
// @version 1.0
// @description API Server for HRM system

// @host localhost:8001
// @BasePath /path

func main() {

	fmt.Println(api.Convert_EmployeeCourses(api.JsonToString(data)))
	LoadConfig("main/configuration.json")

	router := api.Routes()
	router.Run(":8001")
}

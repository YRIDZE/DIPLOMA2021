package main

import (
	api "github.com/YRIDZE/DIPLOMA2021/main/api"
)

func main() {

	api.All_employees = append(api.All_employees, api.Employee{1, "Irina", "Laktionova", "IraLa999@ukr.net", "12345678", "Ukraine"})
	api.All_courses = append(api.All_courses, api.Courses{1, "JAVA-AVA"})
	api.Techn = append(api.Techn, api.Technologies{IdTechnology: 1, Title: "Python"})

	var empl api.Employee
	for _, item := range api.All_employees {
		if item.IdEmployee == 1 {
			empl = item
		}

	}
	var course api.Courses
	for _, item := range api.All_courses {
		if item.IdCourse == 1 {
			course = item
		}

	}

	var techn api.Technologies
	for _, item := range api.Techn {
		if item.IdTechnology == 1 {
			techn = item
		}

	}

	api.Employee_courses = append(api.Employee_courses, api.EmployeeCourses{empl, course, api.S})
	api.Employee_technologies = append(api.Employee_technologies, api.EmployeeTechnology{empl, techn})

	router := api.Routes()
	router.Run(":8001")
}

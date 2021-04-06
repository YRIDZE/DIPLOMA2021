package main

type Employee struct {
	IdEmployee int    `json:"idEmployee"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Country    string `json:"country"`
}
type Courses struct {
	IdCourse int    `json:"idCourse"`
	Title    string `json:"title"`
}

type Technologies struct {
	IdTechnology int    `json:"idTechnoligy"`
	Title        string `json:"title"`
}

type EmployeeCourses struct {
	Employee
	Courses
	Status bool `json:"status"`
}

type EmployeeTechnology struct {
	Employee
	Technologies
}

var employees []Employee
var courses []Courses
var employeeCourses []EmployeeCourses
var technologies []Technologies
var employeeTechnology []EmployeeTechnology

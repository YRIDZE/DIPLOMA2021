package api

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
	IdTechnology int    `json:"idTechnologies"`
	Title        string `json:"title"`
}

type EmployeeCourses struct {
	Employee
	Courses
	Status Stat `json:"status"`
}

type EmployeeTechnology struct {
	Employee
	Technologies
}

type Stat string

const (
	F  Stat = "finished"
	IP Stat = "in progress"
	S  Stat = "suggested"
)

var employees []Employee
var courses []Courses
var employeeCourses []EmployeeCourses
var technologies []Technologies
var employeeTechnology []EmployeeTechnology

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
	IdTechnology int    `json:"idTechnology"`
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

var All_employees []Employee
var All_courses []Courses
var Employee_courses []EmployeeCourses
var Techn []Technologies
var Employee_technologies []EmployeeTechnology

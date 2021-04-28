package d2021

type Employee struct {
	IdEmployee string `json:"idEmployee"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	//Email      string `json:"email"`
	//Password   string `json:"password"`
	//Country    string `json:"country"`
}
type Courses struct {
	IdCourse int    `json:"idCourse"`
	Title    string `json:"title"`
	Status   Stat   `json:"status"`
}

type Technologies struct {
	IdTechnology int    `json:"idTechnology"`
	Title        string `json:"title"`
}

type EmployeeCourses struct {
	Employee `json:"employee"`
	C        []Courses `json:"course"`
}

type EmployeeTechnology struct {
	Employee `json:"employee"`
	T        []Technologies `json:"technology"`
}

type Stat string

const (
	F  Stat = "finished"
	IP Stat = "in progress"
	S  Stat = "suggested"
)

type appConf struct {
	HRM_empl_endp string
	LMS_ip_endp   string
	LMS_f_endp    string
	LMS_s_endp    string
	AL_sug_endp   string
}

var Config appConf

//var AllEmployees []Employee
//var AllCourses []Courses
//var Employee_courses []EmployeeCourses
//var AllTechnologies []Technologies
//var EmployeeTechnologies []EmployeeTechnology

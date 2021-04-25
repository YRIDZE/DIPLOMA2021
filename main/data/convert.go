package data

import (
	"github.com/tidwall/gjson"
)

func JsonToString(data []byte) string {
	return string(data)
}

func ConvertEmployeeCourses(data string) EmployeeCourses {
	var employeeCourse EmployeeCourses

	employeeCourse.Employee =
		Employee{
			IdEmployee: (gjson.Get(data, "employee.idEmployee")).String(),
			FirstName:  (gjson.Get(data, "employee.firstName")).String(),
			LastName:   (gjson.Get(data, "employee.lastName")).String(),
			//Email:      (gjson.Get(data, "employee.email")).String(),
			//Password:   (gjson.Get(data, "employee.password")).String(),
			//Country:    (gjson.Get(data, "employee.country")).String(),
		}

	result := gjson.Get(data, "course")
	var i = 0
	for _, v := range result.Array() {
		employeeCourse.C = append(employeeCourse.C, Courses{})
		employeeCourse.C[i].IdCourse = int(v.Get("idCourse").Int())
		employeeCourse.C[i].Title = v.Get("title").String()
		employeeCourse.C[i].Status = Stat(v.Get("status").String())
		i++
	}
	return employeeCourse
}

func ConvertEmployeeTechnologies(data string) EmployeeTechnology {
	var employeeTechnologies EmployeeTechnology

	employeeTechnologies.Employee =
		Employee{
			IdEmployee: (gjson.Get(data, "employee.idEmployee")).String(),
			FirstName:  (gjson.Get(data, "employee.firstName")).String(),
			LastName:   (gjson.Get(data, "employee.lastName")).String(),
			//Email:      (gjson.Get(data, "employee.email")).String(),
			//Password:   (gjson.Get(data, "employee.password")).String(),
			//Country:    (gjson.Get(data, "employee.country")).String(),
		}

	result := gjson.Get(data, "technology")
	var i = 0
	for _, v := range result.Array() {
		employeeTechnologies.T = append(employeeTechnologies.T, Technologies{})
		employeeTechnologies.T[i].IdTechnology = int(v.Get("idTechnology").Int())
		employeeTechnologies.T[i].Title = v.Get("title").String()
		i++
	}
	return employeeTechnologies
}

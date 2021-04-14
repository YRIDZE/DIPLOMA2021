package api

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

func Decode(s1, s2 string, data []byte) (m map[string]interface{}, m2 []map[string]interface{}) {

	result := gjson.Get(string(data), s1)
	result2 := gjson.Get(string(data), s2)

	err := json.Unmarshal([]byte(result.String()), &m)
	err = json.Unmarshal([]byte(result2.String()), &m2)

	if err != nil {
		fmt.Println("error:", err)
	}
	return
}

func Convert(m map[string]interface{}, m2 []map[string]interface{}) EmployeeCourses {
	var employeeCourse EmployeeCourses

	for key, value := range m {
		switch key {
		case "idEmployee":
			employeeCourse.Employee.IdEmployee = int(value.(float64))
		case "firstName":
			employeeCourse.Employee.FirstName = value.(string)
		case "lastName":
			employeeCourse.Employee.LastName = value.(string)
		case "email":
			employeeCourse.Employee.Email = value.(string)
		case "password":
			employeeCourse.Employee.Password = value.(string)
		case "country":
			employeeCourse.Employee.Country = value.(string)
		}
	}
	for key, value := range m2 {
		employeeCourse.C = append(employeeCourse.C, Courses{})
		for key2, value2 := range value {
			switch key2 {
			case "idCourse":
				employeeCourse.C[key].IdCourse = int(value2.(float64))
			case "title":
				employeeCourse.C[key].Title = value2.(string)
			case "status":
				employeeCourse.C[key].Status = Stat(value2.(string))
			}
		}
	}
	return employeeCourse
}

func JsonToString(data []byte) string {
	return string(data)
}

func Convert_EmployeeCourses(data string) EmployeeCourses {
	var employeeCourse EmployeeCourses

	employeeCourse.Employee =
		Employee{
			IdEmployee: int((gjson.Get(data, "employee.idEmployee")).Int()),
			FirstName:  (gjson.Get(data, "employee.firstName")).String(),
			LastName:   (gjson.Get(data, "employee.lastName")).String(),
			Email:      (gjson.Get(data, "employee.email")).String(),
			Password:   (gjson.Get(data, "employee.password")).String(),
			Country:    (gjson.Get(data, "employee.country")).String(),
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

func Convert_EmployeeTechnologies(data string) EmployeeTechnology {
	var employeeTechnologies EmployeeTechnology

	employeeTechnologies.Employee =
		Employee{
			IdEmployee: int((gjson.Get(data, "employee.idEmployee")).Int()),
			FirstName:  (gjson.Get(data, "employee.firstName")).String(),
			LastName:   (gjson.Get(data, "employee.lastName")).String(),
			Email:      (gjson.Get(data, "employee.email")).String(),
			Password:   (gjson.Get(data, "employee.password")).String(),
			Country:    (gjson.Get(data, "employee.country")).String(),
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

package data

import (
	"github.com/tidwall/gjson"
)

func JsonToString(data []byte) string {
	return string(data)
}

func ConvertEmployeeCourses(data string) []Courses {
	var employeeCourse []Courses

	gjson.ForEachLine(data, func(line gjson.Result) bool {
		for _, v := range line.Array() {
			employeeCourse = append(employeeCourse, Courses{
				IdCourse: int(v.Get("idCourse").Int()),
				Title:    v.Get("title").String(),
				Status:   Stat(v.Get("status").String())})
		}
		return true
	})
	return employeeCourse
}

func ConvertEmployeeTechnologies(data string) []Technologies {
	var employeeTechnologies []Technologies

	gjson.ForEachLine(data, func(line gjson.Result) bool {
		for _, v := range line.Array() {
			employeeTechnologies = append(employeeTechnologies, Technologies{
				IdTechnology: int(v.Get("idTechnology").Int()),
				Title:        v.Get("title").String()})
		}
		return true
	})
	return employeeTechnologies
}

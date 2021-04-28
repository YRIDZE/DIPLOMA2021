package d2021

import (
	"github.com/tidwall/gjson"
)

type CoursesConv struct {
	Title string `json:"title"`
}

func JsonToString(data []byte) string {
	return string(data)
}

func ConvertEmployeeCourses(data string) []CoursesConv {
	var employeeCourse []CoursesConv

	gjson.ForEachLine(data, func(line gjson.Result) bool {
		for _, v := range line.Array() {
			employeeCourse = append(employeeCourse, CoursesConv{
				Title: v.Get("title").String(),
			})
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
				//IdTechnology: int(v.Get("idTechnology").Int()),
				Title: v.Get("title").String()})
		}
		return true
	})
	return employeeTechnologies
}

package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

//даем HRM инфу по назначеным курсам *
func GetEmployeeCourse(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range Employee_courses {
		if item.Employee.IdEmployee == idEmployee {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, Employee_courses)
}

//даем HRM инфу по курсам, которые в процессе прохождения
func GetStartedEmplCourse(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range Employee_courses {
		if item.Employee.IdEmployee == idEmployee && item.Status == IP {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, Employee_courses)
}

//даем HRM инфу про законченые курсы
func GetFinishedCourses(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range Employee_courses {
		if item.Employee.IdEmployee == idEmployee && item.Status == F {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, Employee_courses)
}

// передаем данные курса и работника, который его начал
func PostEmployeeCourse(context *gin.Context) {
	var employeeCourse EmployeeCourses
	x, _ := ioutil.ReadAll(context.Request.Body)

	result := gjson.Get(string(x), "employee")
	result2 := gjson.Get(string(x), "courses")
	result3 := gjson.Get(string(x), "status")

	json.Unmarshal([]byte(result.String()), &employeeCourse.Employee)
	json.Unmarshal([]byte(result2.String()), &employeeCourse.Courses)

	if result3.String() == "finished" {
		employeeCourse.Status = F
	} else if result3.String() == "in progress" {
		employeeCourse.Status = IP
	} else if result3.String() == "suggested" {
		employeeCourse.Status = S
	}

	Employee_courses = append(Employee_courses, employeeCourse)
	context.JSON(http.StatusOK, employeeCourse)
	fmt.Println(Employee_courses)
}

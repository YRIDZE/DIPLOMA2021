package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)
//даем HRM инфу по назначеным курсам *
//TODO: проверить!!

func GetEmployeeCourse(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range employeeCourses {
		if item.Employee.IdEmployee == idEmployee {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, employeeCourses)
}

//даем HRM инфу по курсам, которые в процессе прохождения
//TODO: проверить!!
func GetStartedEmplCourse(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range employeeCourses {
		if item.Employee.IdEmployee == idEmployee && item.Status == IP {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, employeeCourses)
}

//даем HRM инфу про законченые курсы
//TODO: проверить!!
func GetFinishedCourses(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range employeeCourses {
		if item.Employee.IdEmployee == idEmployee && item.Status == F {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, employeeCourses)
}

func PostEmployeeCourse(context *gin.Context) {
	var employeeСourse EmployeeCourses
	body:= context.Request.Body
	x, _ := ioutil.ReadAll(body)
	fmt.Printf("body is: %s \n", string(x))

	_ = json.Unmarshal(x, &employeeСourse)
	prettyJSON, _ := json.MarshalIndent(employeeСourse, "", "    ")
	fmt.Printf("%s\n", string(prettyJSON))

	context.BindJSON(&employeeСourse)
	employeeCourses = append(employeeCourses, employeeСourse)
	context.JSON(http.StatusOK, employeeСourse)
	fmt.Println(employeeCourses)
}


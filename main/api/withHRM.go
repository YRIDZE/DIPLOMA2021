package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

//даем HRM инфу по назначеным курсам *

// GetEmployeeCourse
// @Summary GetCourse
// @Tags Course
// @Description Get suggested courses for employee
// @ID get-suggested-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {object} api.EmployeeCourses
// @Failure 500 {object} api.EmployeeCourses
// @Router /suggested/{idEmployee} [get]
func GetEmployeeCourse(c *gin.Context) {
	var emplCourses EmployeeCourses
	idEmployee, err := strconv.Atoi(c.Params.ByName("idEmployee"))

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	resp, err := http.Get("http://localhost:8002/path/suggested/" + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &emplCourses)

	Employee_courses = append(Employee_courses, emplCourses)

	for _, item := range Employee_courses {
		if item.Employee.IdEmployee == idEmployee {
			fmt.Println(item)
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusOK, Employee_courses)
}

//даем HRM инфу по курсам, которые в процессе прохождения

// GetStartedEmplCourse
// @Summary GetStartedCourses
// @Tags Course
// @Description Get started by employee courses
// @ID get-started-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {object} api.EmployeeCourses
// @Failure 500 {object} api.EmployeeCourses
// @Router /progress/{idEmployee} [get]
func GetStartedEmplCourse(c *gin.Context) {
	idEmployee, _ := strconv.Atoi(c.Params.ByName("IdEmployee"))
	for i, item := range Employee_courses {
		if item.Employee.IdEmployee == idEmployee && item.C[i].Status == IP {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusOK, Employee_courses)
}

//даем HRM инфу про законченые курсы

// GetFinishedCourses
// @Summary GetFinishedCourses
// @Tags Course
// @Description Get finished by employee courses
// @ID get-finished-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {object} api.EmployeeCourses
// @Failure 500 {object} api.EmployeeCourses
// @Router /finished/{idEmployee} [get]
func GetFinishedCourses(c *gin.Context) {
	idEmployee, _ := strconv.Atoi(c.Params.ByName("IdEmployee"))
	for i, item := range Employee_courses {
		if item.Employee.IdEmployee == idEmployee && item.C[i].Status == F {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusOK, Employee_courses)
}

// передаем данные курса и работника, который его начал

// PostEmployeeCourse
// @Summary PostEmployeeCourses
// @Tags Course
// @Description Start taking the course assigned to the employee
// @ID post-course-to-start
// @Accept json
// @Produce json
// @Param input body api.EmployeeCourses true "Employee courses"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {object} api.EmployeeCourses
// @Failure 500 {object} api.EmployeeCourses
// @Router /employee/course [post]
func PostEmployeeCourse(c *gin.Context) {
	var employeeCourse EmployeeCourses
	body, _ := ioutil.ReadAll(c.Request.Body)

	result := gjson.Get(string(body), "employee")
	result2 := gjson.Get(string(body), "course")

	_ = json.Unmarshal([]byte(result.String()), &employeeCourse.Employee)
	_ = json.Unmarshal([]byte(result2.String()), &employeeCourse.C)

	for i, item := range employeeCourse.C {
		if item.Status == "finished" {
			employeeCourse.C[i].Status = F
		} else if item.Status == "in progress" {
			employeeCourse.C[i].Status = IP
		} else if item.Status == "suggested" {
			employeeCourse.C[i].Status = S
		}
	}

	Employee_courses = append(Employee_courses, employeeCourse)
	c.JSON(http.StatusOK, employeeCourse)
	fmt.Println(Employee_courses)
}

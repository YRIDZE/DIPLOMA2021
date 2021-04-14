package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetSuggesterCourse
// @Summary GetCourse
// @Tags Course
// @Description Get suggested courses for employee
// @ID get-suggested-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /suggested/{idEmployee} [get]
func GetSuggesterCourse(c *gin.Context) {

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get(Config.AL_sug_endp + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, Convert_EmployeeCourses(JsonToString(body)))
}

// GetStartedEmplCourse
// @Summary GetStartedCourses
// @Tags Course
// @Description Get started by employee courses
// @ID get-started-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /progress/{idEmployee} [get]
func GetStartedEmplCourse(c *gin.Context) {

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get(Config.LMS_ip_endp + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, Convert_EmployeeCourses(JsonToString(body)))
}

// GetFinishedCourses
// @Summary GetFinishedCourses
// @Tags Course
// @Description Get finished by employee courses
// @ID get-finished-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} status	"Bad request"
// @Failure 500 {string} status	" Internal Server Error"
// @Router /finished/{idEmployee} [get]
func GetFinishedCourses(c *gin.Context) {

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get(Config.LMS_f_endp + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, Convert_EmployeeCourses(JsonToString(body)))
}

// PostEmployeeCourse
// @Summary PostEmployeeCourses
// @Tags Course
// @Description Start taking the course assigned to the employee
// @ID post-course-to-start
// @Accept json
// @Produce json
// @Param input body api.EmployeeCourses true "Employee courses"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /employee/course [post]
func PostEmployeeCourse(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	resp, err := http.Post(Config.LMS_s_endp, "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, Convert_EmployeeCourses(JsonToString(body)))
}

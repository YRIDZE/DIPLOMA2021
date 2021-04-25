package api

import (
	"bytes"
	"fmt"
	"github.com/YRIDZE/DIPLOMA2021/main/data"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// GetStartedEmplCourse
// @Summary GetStartedCourses
// @Tags Course
// @Description Get started by employee courses
// @ID get-started-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} data.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /progress/{idEmployee} [get]
func GetStartedEmplCourse(c *gin.Context) {

	resp, err := http.Get(viper.GetString("LMS_ip_endp") + c.Params.ByName("idEmployee"))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, data.ConvertEmployeeCourses(data.JsonToString(body)))
}

// GetFinishedCourses
// @Summary GetFinishedCourses
// @Tags Course
// @Description Get finished by employee courses
// @ID get-finished-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} data.EmployeeCourses
// @Failure 400,404 {string} status	"Bad request"
// @Failure 500 {string} status	" Internal Server Error"
// @Router /finished/{idEmployee} [get]
func GetFinishedCourses(c *gin.Context) {

	resp, err := http.Get(viper.GetString("LMS_f_endp") + c.Params.ByName("idEmployee"))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, data.ConvertEmployeeCourses(data.JsonToString(body)))
}

// PostEmployeeCourse
// @Summary PostEmployeeCourses
// @Tags Course
// @Description Start taking the course assigned to the employee
// @ID post-course-to-start
// @Accept json
// @Produce json
// @Param input body data.EmployeeCourses true "Employee courses"
// @Success 200 {object} data.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /employee/course [post]
func PostEmployeeCourse(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	resp, err := http.Post(viper.GetString("LMS_s_endp"), "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, data.ConvertEmployeeCourses(data.JsonToString(body)))
}

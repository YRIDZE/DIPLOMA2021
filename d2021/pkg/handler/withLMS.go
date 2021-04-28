package handler

import (
	"bytes"
	"fmt"
	d2021 "github.com/YRIDZE/DIPLOMA2021/main"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetStartedEmplCourse
// @Summary GetStartedCourses
// @Security ApiKeyAuth
// @Tags Course
// @Description Get started by employee courses
// @ID get-started-course
// @Accept json
// @Produce json
// @Success 200 {object} d2021.CoursesConv
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /progress [get]
func (h *Handler) getStartedEmplCourse(c *gin.Context) {

	client := http.Client{}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[0]

	req, err := http.NewRequest("GET", viper.GetString("routs.LMS_ip_endp"), nil)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, d2021.ConvertEmployeeCourses(d2021.JsonToString(body)))
}

// GetFinishedCourses
// @Summary GetFinishedCourses
// @Security ApiKeyAuth
// @Tags Course
// @Description Get finished by employee courses
// @ID get-finished-course
// @Accept json
// @Produce json
// @Success 200 {object} d2021.CoursesConv
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /finished [get]
func (h *Handler) getFinishedCourses(c *gin.Context) {

	client := http.Client{}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[0]

	req, err := http.NewRequest("GET", viper.GetString("routs.LMS_f_endp"), nil)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, d2021.ConvertEmployeeCourses(d2021.JsonToString(body)))
}

// PostEmployeeCourse
// @Summary PostEmployeeCourses
// @Security ApiKeyAuth
// @Tags Course
// @Description Start taking the course assigned to the employee
// @ID post-course-to-start
// @Accept json
// @Produce json
// @Param input body d2021.CoursesConv true "Employee course"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /employee/course [post]
func (h *Handler) postEmployeeCourse(c *gin.Context) {

	client := http.Client{}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[0]

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}

	req, err := http.NewRequest("PUT", viper.GetString("routs.LMS_s_endp"), bytes.NewBuffer(body))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

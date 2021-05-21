package handler

import (
	"bytes"
	"fmt"
	d2021 "github.com/YRIDZE/DIPLOMA2021/main"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetSuggesterCourse
// @Summary GetCourse
// @Security ApiKeyAuth
// @Tags Course
// @Description Get suggested courses for employee
// @ID get-suggested-course
// @Accept json
// @Produce json
// @Success 200 {object} d2021.CoursesConv
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /suggested [get]
func (h *Handler) getSuggesterCourse(c *gin.Context) {

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1] //0

	//path:= viper.GetString("routs.AL_sug_endp")
	req, err := http.NewRequest("GET", "http://localhost:8002/path/suggested", nil)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	client := http.Client{}
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
	//path=viper.GetString("routs.HRM_sug_endp")
	request, err := http.NewRequest("POST", "http://localhost:8000/path/suggcourses", bytes.NewBuffer(body))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	_, err = client.Do(request)

	c.JSON(http.StatusOK, d2021.ConvertEmployeeCourses(d2021.JsonToString(body)))
}

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

	client := http.Client{}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[0]

	req, err := http.NewRequest("GET", viper.GetString("routs.AL_sug_endp"), nil)
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

	request, err := http.NewRequest("POST", viper.GetString("routs.LMS_sug_endp"), bytes.NewBuffer(body))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	request.Header.Set("Content-Type", "application/json")
	_, err = client.Do(request)

	c.JSON(http.StatusOK, d2021.ConvertEmployeeCourses(d2021.JsonToString(body)))
}

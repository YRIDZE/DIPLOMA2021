package api

import (
	"fmt"
	"github.com/YRIDZE/DIPLOMA2021/main/data"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// GetSuggesterCourse
// @Summary GetCourse
// @Tags Course
// @Description Get suggested courses for employee
// @ID get-suggested-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} data.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /suggested/{idEmployee} [get]
func GetSuggesterCourse(c *gin.Context) {

	resp, err := http.Get(viper.GetString("AL_sug_endp") + c.Params.ByName("idEmployee"))
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

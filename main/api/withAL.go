package api

import (
	"fmt"
	"github.com/YRIDZE/DIPLOMA2021/main/data"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetEmployeeSkills
// @Summary GetSkills
// @Tags Technologies
// @Description Get technology skill
// @ID get-technologies
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} data.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /employee[get]
func GetEmployeeSkills(c *gin.Context) {
	client := http.Client{}

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	req, err := http.NewRequest(http.MethodGet, viper.GetString("HRM_empl_endp"), nil)
	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, data.ConvertEmployeeTechnologies(data.JsonToString(body)))
}

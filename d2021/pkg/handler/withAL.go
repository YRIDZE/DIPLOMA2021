package handler

import (
	"fmt"
	"github.com/YRIDZE/DIPLOMA2021/main"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetEmployeeSkills
// @Summary GetEmployeeSkills
// @Security ApiKeyAuth
// @Tags Technology
// @Description Get employee technology skills
// @ID get-technologies
// @Accept json
// @Produce json
// @Success 200 {object} d2021.Technologies
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /empltechn [get]
func (h *Handler) getEmployeeSkills(c *gin.Context) {

	client := http.Client{}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	//path:=viper.GetString("routs.HRM_empl_endp")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/path/empltechn", nil)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, d2021.ConvertEmployeeTechnologies(d2021.JsonToString(body)))
}

package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetEmployeeSkills
// @Summary GetSkills
// @Tags Technologies
// @Description Get technology skill
// @ID get-technologies
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /employee/{idEmployee} [get]
func GetEmployeeSkills(c *gin.Context) {
	var emplTechnology EmployeeTechnology

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get("http://localhost:8000/path/empltechn/" + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		httputil.NewError(c, http.StatusServiceUnavailable, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	err = json.Unmarshal(body, &emplTechnology)
	if err != nil {
		fmt.Println("error:", err)
	}

	c.JSON(http.StatusOK, emplTechnology)
}

package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetEmployeeSkills
//@Summary GetSkills
// @Tags Technologies
// @Description Get technology skill
// @ID get-technologies
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {object} api.EmployeeCourses
// @Failure 500 {object} api.EmployeeCourses
// @Router /employee/{idEmployee} [get]
func GetEmployeeSkills(c *gin.Context) {
	var emplTechnology EmployeeTechnology

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))

	resp, err := http.Get("http://localhost:8000/path/empltechn/" + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &emplTechnology)

	Employee_technologies = append(Employee_technologies, emplTechnology)

	for _, item := range Employee_technologies {
		if item.Employee.IdEmployee == idEmployee {
			//TODO: idTechnologies = 0 WTF?
			fmt.Println(item)
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusOK, Employee_technologies)
}

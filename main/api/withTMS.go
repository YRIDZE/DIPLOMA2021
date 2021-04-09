package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// with TMS
func GetEmployeeSkills(context *gin.Context) {
	idEmployee, _ := strconv.Atoi(context.Params.ByName("IdEmployee"))
	for _, item := range Employee_technologies {
		if item.Employee.IdEmployee == idEmployee {
			context.JSON(http.StatusOK, item)
			return
		}
	}
	context.JSON(http.StatusOK, Employee_technologies)
}

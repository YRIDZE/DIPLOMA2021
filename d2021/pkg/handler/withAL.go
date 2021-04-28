package handler

import (
	"fmt"
	"github.com/YRIDZE/DIPLOMA2021/main"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

func (h *Handler) getEmployeeSkills(c *gin.Context) {

	client := http.Client{}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	fmt.Println(token)

	req, err := http.NewRequest(http.MethodGet, viper.GetString("routs.HRM_empl_endp"), nil)
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
		fmt.Printf("error reading the body: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, d2021.ConvertEmployeeTechnologies(d2021.JsonToString(body)))
}

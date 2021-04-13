package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetSuggesterCourse
// @Summary GetCourse
// @Tags Course
// @Description Get suggested courses for employee
// @ID get-suggested-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /suggested/{idEmployee} [get]
func GetSuggesterCourse(c *gin.Context) {
	var emplCourses EmployeeCourses

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get(Config.AL_sug_endp + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	err = json.Unmarshal(body, &emplCourses)
	if err != nil {
		fmt.Println("error:", err)
	}

	c.JSON(http.StatusOK, emplCourses)
}

// GetStartedEmplCourse
// @Summary GetStartedCourses
// @Tags Course
// @Description Get started by employee courses
// @ID get-started-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /progress/{idEmployee} [get]
func GetStartedEmplCourse(c *gin.Context) {
	var emplCourses EmployeeCourses

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get(Config.LMS_ip_endp + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}

	result := gjson.Get(string(body), "employee")
	result2 := gjson.Get(string(body), "course")

	_ = json.Unmarshal([]byte(result.String()), &emplCourses.Employee)
	_ = json.Unmarshal([]byte(result2.String()), &emplCourses.C)

	c.JSON(http.StatusOK, emplCourses)
}

// GetFinishedCourses
// @Summary GetFinishedCourses
// @Tags Course
// @Description Get finished by employee courses
// @ID get-finished-course
// @Accept json
// @Produce json
// @Param idEmployee path int true "Employee ID"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} status	"Bad request"
// @Failure 500 {string} status	" Internal Server Error"
// @Router /finished/{idEmployee} [get]
func GetFinishedCourses(c *gin.Context) {
	var emplCourses EmployeeCourses

	idEmployee, _ := strconv.Atoi(c.Params.ByName("idEmployee"))
	resp, err := http.Get(Config.LMS_f_endp + strconv.Itoa(idEmployee))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	err = json.Unmarshal(body, &emplCourses)
	if err != nil {
		fmt.Println("error:", err)
	}

	c.JSON(http.StatusOK, emplCourses)
}

// PostEmployeeCourse
// @Summary PostEmployeeCourses
// @Tags Course
// @Description Start taking the course assigned to the employee
// @ID post-course-to-start
// @Accept json
// @Produce json
// @Param input body api.EmployeeCourses true "Employee courses"
// @Success 200 {object} api.EmployeeCourses
// @Failure 400,404 {string} string	"Bad request"
// @Failure 500 {string} string	" Internal Server Error"
// @Router /employee/course [post]
func PostEmployeeCourse(c *gin.Context) {
	var employeeCourse EmployeeCourses
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}
	resp, err := http.Post(Config.LMS_s_endp, "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		newErrorResponse(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}

	result := gjson.Get(string(body), "employee")
	result2 := gjson.Get(string(body), "course")

	err = json.Unmarshal([]byte(result.String()), &employeeCourse.Employee)
	err = json.Unmarshal([]byte(result2.String()), &employeeCourse.C)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i, item := range employeeCourse.C {
		if item.Status == "finished" {
			employeeCourse.C[i].Status = F
		} else if item.Status == "in progress" {
			employeeCourse.C[i].Status = IP
		} else if item.Status == "suggested" {
			employeeCourse.C[i].Status = S
		}
	}

	Employee_courses = append(Employee_courses, employeeCourse)
	c.JSON(http.StatusOK, employeeCourse)
	fmt.Println(Employee_courses)
}

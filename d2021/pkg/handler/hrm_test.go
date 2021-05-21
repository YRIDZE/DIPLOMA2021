package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_getEmployeeSkills(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	testTable := []struct {
		name                string
		headerValue         string
		headerName          string
		token               string
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			headerName:          "Authorization",
			headerValue:         "Bearer",
			token:               os.Getenv("token"),
			expectedStatusCode:  200,
			expectedRequestBody: `[{"title":"PHP"},{"title":"jQuery"}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			handler := NewHandler()

			r := gin.New()
			r.GET("/empltechn", handler.getEmployeeSkills)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/empltechn", nil)
			req.Header.Set(testCase.headerName, fmt.Sprintf("%s %s", testCase.headerValue, testCase.token))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_getFinishedCourses(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	testTable := []struct {
		name                string
		headerValue         string
		headerName          string
		token               string
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			headerName:          "Authorization",
			headerValue:         "Bearer",
			token:               os.Getenv("token"),
			expectedStatusCode:  200,
			expectedRequestBody: `[{"title":"The Complete Oracle SQL Certification Course"},{"title":"Learn API Technical Writing: JSON and XML for Writers"}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			handler := NewHandler()

			r := gin.New()
			r.GET("/finished", handler.getFinishedCourses)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/finished", nil)
			req.Header.Set(testCase.headerName, fmt.Sprintf("%s %s", testCase.headerValue, testCase.token))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_getStartedEmplCourse(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	testTable := []struct {
		name                string
		headerValue         string
		headerName          string
		token               string
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			headerName:          "Authorization",
			headerValue:         "Bearer",
			token:               os.Getenv("token"),
			expectedStatusCode:  200,
			expectedRequestBody: `[{"title":"Git Complete: The definitive, step-by-step guide to Git"},{"title":"Windows PowerShell"},{"title":"C# Intermediate: Classes, Interfaces and OOP"}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			handler := NewHandler()

			r := gin.New()
			r.GET("/progress", handler.getStartedEmplCourse)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/progress", nil)
			req.Header.Set(testCase.headerName, fmt.Sprintf("%s %s", testCase.headerValue, testCase.token))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_getSuggesterCourse(t *testing.T) {

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	testTable := []struct {
		name                string
		headerValue         string
		headerName          string
		token               string
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			headerName:          "Authorization",
			headerValue:         "Bearer",
			token:               os.Getenv("token"),
			expectedStatusCode:  200,
			expectedRequestBody: `[{"title":"C# Intermediate: Classes, Interfaces and OOP"},{"title":"C++: From Beginner to Expert"}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			handler := NewHandler()

			r := gin.New()
			r.GET("/suggested", handler.getSuggesterCourse)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/suggested", nil)
			req.Header.Set(testCase.headerName, fmt.Sprintf("%s %s", testCase.headerValue, testCase.token))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_postEmployeeCourse(t *testing.T) {

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	testTable := []struct {
		name                string
		inputBody           string
		headerValue         string
		headerName          string
		token               string
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			inputBody:           `{"title":"C# Intermediate: Classes, Interfaces and OOP"}`,
			headerName:          "Authorization",
			headerValue:         "Bearer",
			token:               os.Getenv("token"),
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"ok"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			handler := NewHandler()

			r := gin.New()
			r.POST("/employee/course", handler.postEmployeeCourse)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/employee/course",
				bytes.NewBufferString(testCase.inputBody))
			req.Header.Set(testCase.headerName, fmt.Sprintf("%s %s", testCase.headerValue, testCase.token))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})

	}

}

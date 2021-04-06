package api

import "github.com/gin-gonic/gin"

func Routes() *gin.Engine {
	router := gin.Default()
	r := router.Group("/path")
	{
		// предоставляем для HRM
		//1. GET даем HRM инфу по назначеным курсам *
		//2. GET даем HRM инфу по курсам, которые в процессе прохождения
		//3. GET даем HRM инфу про законченые курсы
		//4. POST начни для меня на курс **
		//>>> GET у HRM данные по работнику
		r.GET("/suggested/:idEmployee", GetEmployeeCourse)
		r.GET("/progress/:idEmployee", GetStartedEmplCourse)
		r.GET("/finished/:idEmployee", GetFinishedCourses)
		r.POST("/employee/course", PostEmployeeCourse)

		// предоставляем для TMS
		// 1. GET даем TMS взтые ранее у HRM данные по работрику
		//>>> GET у TMS данные по назначенным курсам, которые отдавем в HRM через (*)
		r.GET("/employee", GetEmployeeSkills)

		//реализуем инт. LMS
		//>>> GET данные по курсам, которые находятся в процессе
		//>>> GET данные по законченным курсам
		//>>> POST передачем данные пользователя и курса, который он хочет начать проходить (**)

	}
	return router
}

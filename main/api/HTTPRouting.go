package api

import (
	_ "github.com/YRIDZE/DIPLOMA2021/main/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {

	router := gin.Default()
	r := router.Group("/path")
	{
		// предоставляем для HRM
		//1. GET даем HRM инфу по назначеным курсам *
		//2. GET даем HRM инфу по курсам, которые в процессе прохождения
		//3. GET даем HRM инфу про законченые курсы
		//4. POST начни для меня на курс **
		//>>> GET у HRM данные по работнику - вложен в *
		r.GET("/suggested/:idEmployee", GetEmployeeCourse)
		r.GET("/progress/:idEmployee", GetStartedEmplCourse)
		r.GET("/finished/:idEmployee", GetFinishedCourses)
		r.POST("/employee/course", PostEmployeeCourse)

		// предоставляем для TMS
		// 1. GET даем TMS взятые ранее у HRM данные по работнику
		//>>> GET у TMS данные по назначенным курсам, которые отдаем в HRM через (*)
		r.GET("/employee/:idEmployee", GetEmployeeSkills)

		//реализуем инт. LMS
		//>>> GET данные по курсам, которые находятся в процессе
		//>>> GET данные по законченным курсам
		//>>> POST передаем данные пользователя и курса, который он хочет начать проходить (**)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}

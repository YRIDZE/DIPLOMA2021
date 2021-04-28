package handler

import (
	_ "github.com/YRIDZE/DIPLOMA2021/main/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/path")
	{
		// предоставляем для HRM
		//1. GET даем HRM инфу по назначеным курсам *
		//2. GET даем HRM инфу по курсам, которые в процессе прохождения
		//3. GET даем HRM инфу про законченые курсы
		//4. POST начни для меня на курс **
		//>>> GET у HRM данные по работнику - вложен в *
		api.GET("/suggested", h.getSuggesterCourse)
		api.GET("/progress", h.getStartedEmplCourse)
		api.GET("/finished", h.getFinishedCourses)
		api.POST("/employee/course", h.postEmployeeCourse)

		// предоставляем для TMS
		// 1. GET даем TMS взятые ранее у HRM данные по работнику
		//>>> GET у TMS данные по назначенным курсам, которые отдаем в HRM через (*)
		api.GET("/empltechn", h.getEmployeeSkills)

		//реализуем инт. LMS
		//>>> GET данные по курсам, которые находятся в процессе
		//>>> GET данные по законченным курсам
		//>>> POST передаем данные пользователя и курса, который он хочет начать проходить (**)

	}
	return router
}

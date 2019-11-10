package router

import (
	"App/http/controllers"
	"App/http/middlewares"
	"github.com/gin-gonic/gin"
)

func DeclareRoutes(engine *gin.Engine) {

	engine.Use(middlewares.Log)
	engine.GET("/", controllers.MainController{}.Homepage)
	engine.GET("/ping", controllers.MainController{}.Ping)

	courses := engine.Group("/courses")
	courses.
		Use(middlewares.LogLatency)
	{
		courseController := controllers.CourseController{}
		courses.GET("", courseController.List)
		courses.POST("", courseController.Post)
		courses.PATCH("/:courseId", courseController.Patch)
		courses.GET("/:courseId", courseController.Get)
		courses.DELETE("/:courseId", courseController.Delete)
	}
}

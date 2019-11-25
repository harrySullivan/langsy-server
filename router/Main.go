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

	users := engine.Group("/users")
	users.
		Use(middlewares.LogLatency)
	{
		usersController := controllers.UserController{}
		users.GET("", usersController.List)
		users.POST("", usersController.Post)
		users.PATCH("/:userId", usersController.Patch)
		users.GET("/:userId", usersController.Get)
		users.DELETE("/:userId", usersController.Delete)
	}

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

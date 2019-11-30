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

	auth := engine.Group("/auth")
	auth.
		Use(middlewares.LogLatency)
	{
		authController := controllers.AuthController{}
		auth.POST("/login", authController.Login)
		auth.POST("/signUp", authController.SignUp)
	}

	users := engine.Group("/users")
	users.
		Use(middlewares.LogLatency).
		Use(middlewares.AuthRequired)
	{
		usersController := controllers.UserController{}
		users.GET("", usersController.List)
		users.PATCH("/:userId", usersController.Patch)
		users.GET("/:userId", usersController.Get)
		users.DELETE("/:userId", usersController.Delete)
	}

	courses := engine.Group("/courses")
	courses.
		Use(middlewares.LogLatency).
		Use(middlewares.AuthRequired)
	{
		courseController := controllers.CourseController{}
		courses.GET("", courseController.List)
		courses.POST("", courseController.Post)
		courses.PATCH("/:courseId", courseController.Patch)
		courses.GET("/:courseId", courseController.Get)
		courses.DELETE("/:courseId", courseController.Delete)
	}

	phrases := engine.Group("/phrases")
	phrases.
		Use(middlewares.LogLatency).
		Use(middlewares.AuthRequired)
	{
		phraseController := controllers.PhraseController{}
		phrases.GET("", phraseController.List)
		phrases.GET("/random/:lang", phraseController.GetRandom)
	}

	sources := engine.Group("/sources")
	sources.
		Use(middlewares.LogLatency).
		Use(middlewares.AuthRequired)
	{
		sourceController := controllers.SourceController{}
		sources.GET("", sourceController.List)
		sources.POST("", sourceController.Post)
	}
}

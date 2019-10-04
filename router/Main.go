package router

import (
	"App/http/controllers"
	"App/http/middlewares"
	"github.com/gin-gonic/gin"
)

func DeclareRoutes(engine *gin.Engine) {
	engine.Use(middlewares.Etag())

	engine.Use(middlewares.Log)
	engine.GET("/ping", controllers.MainController{}.Ping)

	api := engine.Group("/api")
	{
		api.GET("/", controllers.MainController{}.Homepage)

		users := api.Group("/users")
		users.
			Use(middlewares.LoadUser).
			Use(middlewares.LogLatency)
		{
			userController := controllers.UserController{}
			users.GET("", userController.List)
			users.POST("", userController.Post)
			users.PATCH("/:userId", userController.Patch)
			users.GET("/:userId", userController.Get)
			users.DELETE("/:userId", userController.Delete)
		}
	}
	engine.NoRoute(controllers.MainController{}.Client)
	engine.Static("/static", "./dist")
}

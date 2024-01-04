package routers

import (
	"github.com/gin-gonic/gin"
	"webSurfer/handlers"
)

func SetupUsersRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("", handlers.GetUsers)
		userRouter.GET("/:id", handlers.GetUser)
		userRouter.POST("", handlers.CreateUser)
		userRouter.DELETE("/:id", handlers.DeleteUser)
	}
}

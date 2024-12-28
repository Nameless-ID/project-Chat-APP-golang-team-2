package routes

import (
	"user-service/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *controller.UserController) {
    
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/:user_id", userHandler.GetUserInfo)   
        userRoutes.GET("/online", userHandler.GetOnlineUsers)  
    }
}

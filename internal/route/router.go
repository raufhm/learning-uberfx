package route

import (
	"github.com/gin-gonic/gin"
	"github.com/raufhm/learning-uberfx/internal/handler"
)

func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	v1 := router.Group("iam/v1")
	{
		v1.GET("/user:id", userHandler.GetUserByID)
		// Add more routes as needed
	}
}

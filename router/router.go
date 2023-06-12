package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raufhm/learning-uberfx/handler"
)

func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	router.GET("/users", userHandler.GetUserByID)
	// Add more routes as needed
}

package main

import (
	"WebService/common/middleware"
	"WebService/handler"

	"github.com/gin-gonic/gin"
)

// customizeRegister register customize routers.
func register(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	userGroup := r.Group("/user", middleware.Login())
	{
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/passwd/reset", handler.ResetPasswd)
	}
}

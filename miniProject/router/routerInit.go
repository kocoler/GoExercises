package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kocoler/GoExercises/miniProject/handler"
	"github.com/kocoler/GoExercises/miniProject/middleware"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	Router.POST("/login",handler.UserLogin)
	Router.Use(middleware.JwtAAuth())
	Router.GET("/test",handler.Test)

}
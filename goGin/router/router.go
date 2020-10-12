package router

import (
	"github.com/gin-gonic/gin"
	"goGin/funcs"
	"goGin/handler/login"
	"goGin/handler/redirect"
	"goGin/handler/register"
	"goGin/handler/userPassword"
	"goGin/handler/usersInfo"
	"goGin/middleware"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	//Router.Use(gin.Logger())
	Router.Use(middleware.TimeLog())
	Router.POST("/register", middleware.Welcome(), register.Register)
	Router.POST("/login", login.Login)
	Router.POST("/password/forget", userPassword.Forget)
	Router.GET("/", redirect.RedirectToMain)
	Router.POST("/magicalTest", funcs.Test1)
	Router.GET("/home", middleware.Welcome(), redirect.Home)
	Router.GET("/userInfoR", usersInfo.UserInfoR)
	Router.POST("/userInfoW", usersInfo.UsersInfoWrite)
	return
}

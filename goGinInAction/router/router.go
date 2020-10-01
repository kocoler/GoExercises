package router

import (
	"github.com/gin-gonic/gin"
	"goGinInAction/handler"
	"goGinInAction/middleware"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	Router.Use(middleware.Time())
	//增添
	Router.POST("/student/creat", handler.CreatStudent)
	Router.POST("/classroom/creat", handler.CreatClassroom)
	Router.POST("/course/creat", handler.CreatCourse)

	//基本信息
	Router.GET("/student/info", handler.StudentInfo)     //student=
	Router.GET("/classroom/info", handler.ClassroomInfo) //classroom=
	Router.GET("/course/info", handler.CourseInfo)       //course=

	//信息
	Router.GET("/student/classroom", handler.StudentClassroom) //student=
	Router.GET("/student/course", handler.StudentCourse)       //student=

	//分配
	Router.POST("/student/dis", handler.DisStudentCourse) //student=
	Router.POST("/course/dis", handler.DisCourseClass)    //course=

	//更改
	Router.POST("/student/info", handler.UpdateStudentInfo)     //student=
	Router.POST("/classroom/info", handler.UpdateClassroomInfo) //classroom=
	Router.POST("/course/info", handler.UpdateCourseInfo)       //course=

	Router.POST("/")
	return
}

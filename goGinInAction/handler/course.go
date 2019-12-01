package handler

import (
	"github.com/gin-gonic/gin"
	"goGinInAction/model"
)

func CreatCourse(c *gin.Context) {
	type courses struct {
		Course	[]string	`json:"course"`
	}
	var tmpCourse courses
	if err := c.BindJSON(&tmpCourse); err != nil {
		c.JSON(400,gin.H{"message":"Bad Request!"})
		return
	}
	for _,v := range tmpCourse.Course {
		c.JSON(200,gin.H{
			"message":model.CreatCourse(v),
			})
	}
	c.JSON(200,gin.H{
		"message":"Creat Success",
	})
}

func UpdateCourseInfo(c *gin.Context) {
	courseName:= c.Query("name")
	if !model.CheckCourse(courseName) {
		c.JSON(200,gin.H{
			"message":"课程不存在",
		})
		return
	}
	var tmpCourse model.Course
	tmpCourse.Course = courseName
	if err := c.BindJSON(&tmpCourse); err != nil {
		c.JSON(400,gin.H{
			"message":"Bad Request!",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":model.UpdateCourseInfo(tmpCourse),
	})
	return
}

func DisCourseClass(c *gin.Context) {
	className := c.Query("name")
	if !model.CheckClassRoom(className) {
		c.JSON(200,gin.H{
			"message":"教室不存在",
		})
		return
	}
	type tmpCourse struct {
		Course	[]string	`json:"course"`
	}
	var tmpCourses tmpCourse
	if err := c.BindJSON(&tmpCourses); err != nil {
		c.JSON(400,gin.H{"message":"Bad Request!"})
		return
	}
	for _,v := range tmpCourses.Course {
		c.JSON(200,gin.H{
			"message":model.DisCourseClass(v,className),
		})

	}
	c.JSON(200,gin.H{
		"message":"Success!",
	})
	return
}

func CourseInfo(c *gin.Context) {
	var tmpCourse model.Course
	tmpCourse.Course = c.Query("name")
	if !model.CheckCourse(tmpCourse.Course) {
		c.JSON(200,gin.H{
			"message":"课程不存在",
		})
		return
	}
	tmpCourse, content := model.CourseInfo(tmpCourse.Course)
	if len(content) != 0 {
		c.JSON(400,gin.H{
			"message":content,
		})
		return
	}
	c.JSON(200,tmpCourse)
}
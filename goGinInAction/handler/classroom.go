package handler

import (
	"github.com/gin-gonic/gin"
	"goGinInAction/model"
)

type classrooms struct {
	Classroom []string `json:"classroom"`
}

func CreatClassroom(c *gin.Context) {

	var tmpClassroom classrooms
	if err := c.BindJSON(&tmpClassroom); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	for _, v := range tmpClassroom.Classroom {
		c.JSON(200, gin.H{
			"message": model.CreatClassroom(v),
		})
	}
	c.JSON(200, gin.H{
		"message": "Creat Success",
	})
}

func UpdateClassroomInfo(c *gin.Context) {
	var tmpClassroom model.Classroom
	tmpClassroom.Classroom = c.Query("name")
	if !model.CheckClassRoom(tmpClassroom.Classroom) {
		c.JSON(200, gin.H{
			"message": "教室不存在",
		})
		return
	}
	if err := c.BindJSON(&tmpClassroom); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": model.UpdateClassroomInfo(tmpClassroom),
	})
	return
}

func ClassroomInfo(c *gin.Context) {
	var tmpClassroom model.Classroom
	tmpClassroom.Classroom = c.Query("name")
	if !model.CheckClassRoom(tmpClassroom.Classroom) {
		c.JSON(200, gin.H{
			"message": "教室不存在",
		})
		return
	}
	tmpClassroom, content := model.ClassroomInfo(tmpClassroom.Classroom)
	if len(content) != 0 {
		c.JSON(400, gin.H{
			"message": content,
		})
		return
	}
	c.JSON(200, tmpClassroom)
}

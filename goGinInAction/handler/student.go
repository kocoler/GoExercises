package handler

import (
	"github.com/gin-gonic/gin"
	"goGinInAction/model"
	"log"
)

type students struct {
	Student []string `json:"student"`
}

type tmpClassroom struct {
	Classroom []string `json:"classroom"`
}

func CreatStudent(c *gin.Context) {
	//var tmpStudent model.Student

	var tmpStudent students
	if err := c.BindJSON(&tmpStudent); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request!"})
		return
	}
	for _, v := range tmpStudent.Student {
		c.JSON(200, gin.H{
			"message": model.CreatStudent(v),
		})
	}
	c.JSON(200, gin.H{
		"message": "Creat Success",
	})
}

func StudentCourse(c *gin.Context) {
	studentName := c.Query("name")
	type tmpCourse struct {
		Course []string `json:"course"`
	}
	var course tmpCourse
	course.Course = model.StudentCourse(studentName)
	c.JSON(200, course)
}

func StudentClassroom(c *gin.Context) {
	studentName := c.Query("name")

	var classroom tmpClassroom
	classroom.Classroom = model.StudentClassroom(studentName)
	c.JSON(200, classroom)
}

func UpdateStudentInfo(c *gin.Context) {
	var tmpStudent model.Student
	tmpStudent.StudentName = c.Query("name")
	if !model.CheckStudent(tmpStudent.StudentName) {
		c.JSON(200, gin.H{
			"message": "学生不存在",
		})
		return
	}
	if err := c.BindJSON(&tmpStudent); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request!"})
		log.Println(err)
		return
	}
	content := model.UpdateStudentInfo(tmpStudent)
	c.JSON(200, gin.H{
		"message": content,
	})
	return
}

func DisStudentCourse(c *gin.Context) {
	courseName := c.Query("name")
	if !model.CheckCourse(courseName) {
		c.JSON(200, gin.H{
			"message": "课程不存在",
		})
		return
	}

	var tmpStudents students
	if err := c.BindJSON(&tmpStudents); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request!"})
		return
	}
	for _, v := range tmpStudents.Student {
		c.JSON(200, gin.H{
			"message": model.DisStudentCourse(v, courseName),
		})
	}
	c.JSON(200, gin.H{
		"message": "Success!",
	})
	return
}

func StudentInfo(c *gin.Context) {
	var tmpStudent model.Student
	tmpStudent.StudentName = c.Query("name")
	if !model.CheckStudent(tmpStudent.StudentName) {
		c.JSON(200, gin.H{
			"message": "学生不存在",
		})
		return
	}
	tmpStudent, content := model.StudentInfo(tmpStudent.StudentName)
	if len(content) != 0 {
		c.JSON(400, gin.H{
			"message": content,
		})
		return
	}
	c.JSON(200, tmpStudent)
}

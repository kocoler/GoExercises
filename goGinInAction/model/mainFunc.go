package model

import (
	"log"
)

type Student struct {
	Number      int    `gorm:"number" json:"-"`
	Course      string `gorm:"course" json:"-"`
	StudentInfo string `gorm:"student_info" json:"student_info"`
	StudentName string `gorm:"student_name" json:"student_name"`
}

type Classroom struct {
	ClassroomId   int    `gorm:"classroom_id" json:"classroom_id"`
	Classroom     string `gorm:"class_room" json:"classroom"`
	ClassroomInfo string `gorm:"classroom_info" json:"classroom_info"`
}

type Course struct {
	CourseId   int    `gorm:"course_id" json:"course_id"`
	Course     string `gorm:"course" json:"course"`
	CourseInfo string `gorm:"course_info" json:"course_info"`
	Classroom  string `gorm:"classroom" json:"classroom"`
}

func CheckStudent(studentName string) bool {
	var tmpUser Student
	if Db.Self.Model(&Student{}).Where(Student{StudentName: studentName}).First(&tmpUser); len(tmpUser.StudentName) != 0 {
		return true
	}
	return false
}

func CreatStudent(studentName string) string {
	if CheckStudent(studentName) {
		return "学生" + studentName + "已存在"
	}
	var tmpUser = Student{StudentName: studentName}
	if err := Db.Self.Model(&Student{}).Create(&tmpUser).Error; err != nil {
		log.Print("CreatStudent:")
		log.Println(err)
		return "创建" + studentName + "错误"
	}
	return "创建" + studentName + "成功"
}

func CheckClassRoom(classRoom string) bool {
	var tmpClassRoom Classroom
	if Db.Self.Model(&Classroom{}).Where(Classroom{Classroom: classRoom}).First(&tmpClassRoom); len(tmpClassRoom.Classroom) != 0 {
		return true
	}
	return false
}

func CreatClassroom(classroom string) string {
	if CheckClassRoom(classroom) {
		return "教室" + classroom + "已存在"
	}
	var tmpClassRoom Classroom
	tmpClassRoom.Classroom = classroom
	if err := Db.Self.Model(&Classroom{}).Create(&tmpClassRoom).Error; err != nil {
		log.Print("CreatClassroom:")
		log.Println(err)
		return "创建" + classroom + "错误"
	}
	return "创建" + classroom + "成功"
}

func CheckCourse(course string) bool {
	var tmpCourse Course
	if Db.Self.Model(&Course{}).Where(Course{Course: course}).First(&tmpCourse); len(tmpCourse.Course) != 0 {
		return true
	}
	return false
}

func CreatCourse(course string) string {
	if CheckCourse(course) {
		return "课程" + course + "已存在"
	}
	var tmpCourse Course
	tmpCourse.Course = course
	if err := Db.Self.Model(&Course{}).Create(&tmpCourse).Error; err != nil {
		log.Print("CreatCourse")
		log.Println(err)
		return "创建" + course + "错误"
	}
	return "创建" + course + "成功"
}

func StudentCourse(student string) []string {
	var courses []Student
	var a []string
	if err := Db.Self.Model(&Student{}).Where(Student{StudentName: student}).Find(&courses).Error; err != nil {
		log.Print("StudentCourse")
		log.Println(err)
		return a
	}
	for _, v := range courses {
		a = append(a, v.Course)
	}
	return a
}

func CourseClassroom(course string) []string {
	var classroom []Course
	var a []string
	if err := Db.Self.Model(&Course{}).Where(Course{Course: course}).Find(&classroom).Error; err != nil {
		log.Print("CourseClassroom:")
		log.Println(err)
		return a
	}
	for _, v := range classroom {
		a = append(a, v.Classroom)
	}
	return a
}

func StudentClassroom(student string) []string {
	var courses []string
	var classroom []string
	courses = StudentCourse(student)
	for _, v := range courses {
		for _, a := range CourseClassroom(v) {
			classroom = append(classroom, a)
		}
	}
	return classroom
}

func UpdateStudentInfo(tmpUser Student) string {
	if err := Db.Self.Model(&Student{}).Where(Student{StudentName: tmpUser.StudentName}).Update(&tmpUser).Error; err != nil {
		log.Print("UpdateStudentInfo")
		log.Println(err)
		return "数据库错误"
	}
	return "更新学生信息成功"
}

func UpdateClassroomInfo(tmpClassroom Classroom) string {
	if err := Db.Self.Model(&Classroom{}).Where(Classroom{Classroom: tmpClassroom.Classroom}).Update(&tmpClassroom).Error; err != nil {
		log.Print("UpdateClassroomInfo:")
		log.Println(err)
		return "数据库错误"
	}
	return "更新教室信息成功"
}

func UpdateCourseInfo(tmpCourse Course) string {
	if err := Db.Self.Model(&Course{}).Where(Course{Course: tmpCourse.Course}).Update(&tmpCourse).Error; err != nil {
		log.Print("UpdateCourseInfo:")
		log.Println(err)
		return "数据库错误"
	}
	return "更新课程信息成功"
}

func DisStudentCourse(studentName string, course string) string {
	if !CheckStudent(studentName) {
		return studentName + "学生不存在"
	}
	var tmpStudent Student
	tmpStudent.StudentName = studentName
	tmpStudent.Course = course
	if err := Db.Self.Model(&Student{}).Where(Student{StudentName: studentName}).Update(&tmpStudent).Error; err != nil {
		log.Print("DisStudentCourse:")
		log.Println(err)
		return "数据库错误"
	}
	return "分配学生" + studentName + "到" + course + "成功"
}

func DisCourseClass(courseName string, class string) string {
	if !CheckCourse(courseName) {
		return courseName + "课程不存在"
	}
	var tmpCourse Course
	tmpCourse.Course = courseName
	tmpCourse.Classroom = class
	if err := Db.Self.Model(&Course{}).Where(Course{Course: courseName}).Update(&tmpCourse).Error; err != nil {
		log.Print("DisCourseClass:")
		log.Println(err)
		return "数据库错误"
	}
	return "分配课程" + courseName + "到" + class + "成功"
}

func StudentInfo(studentName string) (Student, string) {
	var tmpStudent Student
	if err := Db.Self.Model(&Student{}).Where(Student{StudentName: studentName}).First(&tmpStudent).Error; err != nil {
		log.Print("获取学生信息错误")
		log.Println(err)
		return Student{}, "数据库错误"
	}
	return tmpStudent, ""
}

func ClassroomInfo(classroomName string) (Classroom, string) {
	var tmpClassroom Classroom
	if err := Db.Self.Model(&Classroom{}).Where(Classroom{Classroom: classroomName}).First(&tmpClassroom).Error; err != nil {
		log.Print("获取教室信息错误")
		log.Println(err)
		return Classroom{}, "数据库错误"
	}
	return tmpClassroom, ""
}

func CourseInfo(courseName string) (Course, string) {
	var tmpCourse Course
	if err := Db.Self.Model(&Course{}).Where(Course{Course: courseName}).First(&tmpCourse).Error; err != nil {
		log.Print("获取课程信息错误")
		log.Println(err)
		return Course{}, "数据库错误"
	}
	return tmpCourse, ""
}

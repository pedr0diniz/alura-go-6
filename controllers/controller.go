package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedr0diniz/alura-go-5/database"
	"github.com/pedr0diniz/alura-go-5/models"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.JSON(200, students)
}

func FindStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}

	c.JSON(200, student)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says": "Hey " + name + ", what's up?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Student successfully deleted",
	})
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func FindStudentByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}

	c.JSON(200, student)
}

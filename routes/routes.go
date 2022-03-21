package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedr0diniz/alura-go-5/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCpf)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}

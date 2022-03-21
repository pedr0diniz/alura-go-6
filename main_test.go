package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pedr0diniz/alura-go-5/controllers"
	"github.com/pedr0diniz/alura-go-5/database"
	"github.com/pedr0diniz/alura-go-5/models"
	"github.com/stretchr/testify/assert"
)

var studentID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{
		Name: "Test Student",
		RG:   "123456789",
		CPF:  "12345678901",
	}
	database.DB.Create(&student)
	studentID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, studentID)
}

func TestShould_Return_Greeting_With_Parameter(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)

	req, _ := http.NewRequest("GET", "/ziniD", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")

	responseMock := `{"API says":"Hey ziniD, what's up?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody))

}

func TestShould_List_All_Students(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShould_Get_Student_by_Cpf(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()

	r.GET("/students/cpf/:cpf", controllers.FindStudentByCpf)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShould_Find_Student_By_Id(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()

	r.GET("/students/:id", controllers.FindStudentById)

	resourceId := strconv.Itoa(studentID)
	req, _ := http.NewRequest("GET", "/students/"+resourceId, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Test Student", studentMock.Name, "Names must be equal")
	assert.Equal(t, "123456789", studentMock.RG, "RGs must be equal")
	assert.Equal(t, "12345678901", studentMock.CPF, "CPFs must be equal")
}

func TestShould_Delete_Student(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	r := SetupTestRoutes()

	r.DELETE("/students/:id", controllers.DeleteStudent)

	resourceId := strconv.Itoa(studentID)
	req, _ := http.NewRequest("DELETE", "/students/"+resourceId, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShould_Edit_Student(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()

	r.PATCH("/students/:id", controllers.EditStudent)

	student := models.Student{
		Name: "Test Student",
		RG:   "123456700",
		CPF:  "47123456789",
	}
	jsonStudent, _ := json.Marshal(student)

	resourceId := strconv.Itoa(studentID)
	req, _ := http.NewRequest("PATCH", "/students/"+resourceId, bytes.NewBuffer(jsonStudent))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var updatedMockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &updatedMockStudent)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, student.Name, updatedMockStudent.Name, "Names must be equal")
	assert.Equal(t, student.RG, updatedMockStudent.RG, "RGs must be equal")
	assert.Equal(t, student.CPF, updatedMockStudent.CPF, "CPFs must be equal")
}

package api

import (
	"github.com/SrVariable/mongo-exporter/internal/student"

	"github.com/gin-gonic/gin"
)

func addStudentRoutes(rg *gin.RouterGroup) {
	students := rg.Group("/students")

	students.GET("/", student.GetStudentsHandler)
	students.GET("/:id", student.GetStudentByIDHandler)
}

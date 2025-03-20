package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudentsHandler(c *gin.Context) {
	s := GetStudents()
	c.JSON(http.StatusOK, s)
}

func GetStudentByIDHandler(c *gin.Context) {
	id := c.Param("id")
	s, err := findStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

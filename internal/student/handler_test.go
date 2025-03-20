package student

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStudent(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetStudentsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := []student{
		{ID: "1", Name: "Ana", Age: 18},
		{ID: "2", Name: "Ben", Age: 20},
		{ID: "3", Name: "Casey", Age: 22},
		{ID: "4", Name: "Denise", Age: 24},
		{ID: "5", Name: "Elmo", Age: 23},
	}
	var got []student
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

func TestGetStudentByID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "2"})

	GetStudentByIDHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := student{ID: "2", Name: "Ben", Age: 20}
	var got student
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

func TestGetStudentByIDNegative(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "-1"})

	GetStudentByIDHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "student not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

func TestGetStudentByIDLessThanOne(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "0"})

	GetStudentByIDHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "student not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

// This test is an edge case with `id` off by 1.
// It will probably have to change when I connect to a database.
func TestGetStudentByIDGreaterThanFive(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "6"})

	GetStudentByIDHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "student not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

func TestGetStudentByIDNotNumber(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "a"})

	GetStudentByIDHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "student not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

package student

import "errors"

// This will eventually be removed because I'll retrieve the data from a database.
// So it only works to have some data to test the API
var students = []student{
	{ID: "1", Name: "Ana", Age: 18},
	{ID: "2", Name: "Ben", Age: 20},
	{ID: "3", Name: "Casey", Age: 22},
	{ID: "4", Name: "Denise", Age: 24},
	{ID: "5", Name: "Elmo", Age: 23},
}

func GetStudents() []student {
	return students
}

func GetStudentByID(id string) (*student, error) {
	for _, s := range students {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, errors.New("student not found")
}

package student

func findAllStudents() []student {
	return GetStudents()
}

func findStudentByID(id string) (*student, error) {
	s, err := GetStudentByID(id)
	return s, err
}

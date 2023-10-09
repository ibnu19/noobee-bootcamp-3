package sesi1introdution

import (
	"fmt"
	"testing"
)

type Student struct {
	Name  string
	Class int8
}

func (s *Student) SetMyName(newName string) {
	s.Name = newName
}

func (s *Student) CallMyName() {
	fmt.Println("Hello, My name is", s.Name)
}

func Test(t *testing.T) {
	student := Student{
		Name:  "David Maulana",
		Class: 12,
	}
	fmt.Println("Hello, My name is", student.Name)

	student.SetMyName("Rizki Ridho")
	student.CallMyName()
}

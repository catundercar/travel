package student

import (
	"go.study.org/mongoDemo/driver"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

func (s *Student) Add(e driver.Engine) error {
	err := e.InsertOne("students", s)
	if err != nil {
		return err
	}
	return nil
}

type StudentList struct {
	Students []*Student
	Total    int
}

func (s *StudentList) Add(e driver.Engine) error {
	ds := make([]interface{}, 0, len(s.Students))
	ds = append(ds, s.Students)
	err := e.InsertMany("students", ds)
	if err != nil {
		return err
	}

	return nil
}

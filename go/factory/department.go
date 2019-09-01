package factory

import "fmt"

type DepartmentIf interface {
	Create(*department) error
	Get(int) *department
}

func NewDepartment(id int, name string) *department {
	return &department{
		id:   id,
		name: name,
	}
}

type department struct {
	id   int
	name string
}

type MySQLDepartment struct {
	// private fields
}

func (m *MySQLDepartment) Create(d *department) error {
	fmt.Printf("create department:%+v\n", d)
	return nil
}

func (m *MySQLDepartment) Get(id int) *department {
	fmt.Println("get department ...")
	return nil
}

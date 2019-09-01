package main

import "github.com/blurty/design-pattens/go/factory"

func main() {
	f := factory.NewDataAccess()
	db := f.CreateUser()
	user := factory.NewUser(1, "bob")
	db.Create(user)
	db.Get(1)

	departmentDb := f.CreateDepartment()
	department := factory.NewDepartment(2, "didi")
	departmentDb.Create(department)
	departmentDb.Get(2)
}

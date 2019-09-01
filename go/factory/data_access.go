package factory

import (
	"reflect"
)

var (
	mysqlUser       = reflect.TypeOf(&MySQLUser{})
	mysqlDepartment = reflect.TypeOf(&MySQLDepartment{})
)

var typeRegistry = make(map[string]reflect.Type)

func init() {
	// items in myTypes must implement UserIf interface{}
	myTypes := []reflect.Type{
		mysqlUser,
		mysqlDepartment,
	}
	for _, t := range myTypes {
		typeRegistry[t.String()] = t
	}
}

func NewDataAccess() *dataAccess {
	return &dataAccess{
		dbUser:       mysqlUser.String(),
		dbDepartment: mysqlDepartment.String(),
	}
}

type dataAccess struct {
	dbUser       string
	dbDepartment string
}

func (da *dataAccess) CreateUser() UserIf {
	elem := reflect.New(typeRegistry[da.dbUser]).Elem().Interface()
	return elem.(UserIf)
}

func (da *dataAccess) CreateDepartment() DepartmentIf {
	elem := reflect.New(typeRegistry[da.dbDepartment]).Elem().Interface()
	return elem.(DepartmentIf)
}

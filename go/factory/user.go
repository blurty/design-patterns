package factory

import "fmt"

type UserIf interface {
	Create(user *User) error
	Get(id int) *User
}

func NewUser(id int, name string) *User {
	return &User{
		Id:   id,
		Name: name,
	}
}

type User struct {
	Id   int
	Name string
}

type MySQLUser struct {
	// private fields
}

func (m *MySQLUser) Create(user *User) error {
	fmt.Printf("create user:%+v\n", user)
	return nil
}

func (m *MySQLUser) Get(id int) *User {
	fmt.Println("get user ...")
	return nil
}

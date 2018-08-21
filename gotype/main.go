package main

import (
	"base01/gotype/dao"
	"base01/gotype/domain"
	"base01/gotype/service"
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u user) sayName() {
	fmt.Println("my name is:", u.name)
}

//原来user的age不变,新返回的user的age为修改后的
func (u user) changeAge(newAge int) user {
	u.age = newAge
	return u
}

/*原来的和新返回的age都会改变
	user := &user{
			name: "wmz",
			age:  10,
	}

	user1 := user.changeOriAge(20)
	fmt.Println("user1 age is:", user1.age)
	fmt.Println("user age is:", user.age)
    user1 age is: 20
    user age is: 20
*/
func (u *user) changeOriAge(newAge int) user {
	u.age = newAge
	return *u
}

func createPerson(creator dao.Dao) domain.Person {
	return creator.Create()
}

func main() {
	creator := service.PersonFactory{}

	person := createPerson(creator)

	fmt.Println("person is:", person)
}

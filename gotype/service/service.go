package service

import (
	"base01/gotype/domain"
)

type PersonFactory struct {
}

func (creator PersonFactory) Create() domain.Person {
	person := domain.Person{
		Sex:  "男",
		Name: "wmz",
		Age:  10,
	}
	return person
}

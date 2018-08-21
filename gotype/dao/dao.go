package dao

import "base01/gotype/domain"

type Dao interface {
	Create() domain.Person
}

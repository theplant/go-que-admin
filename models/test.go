package models

import "github.com/jinzhu/gorm"

type TestUser struct {
	gorm.Model

	Name string
	Age  int
}

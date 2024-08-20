package schemas

import (
	"gorm.io/gorm"
)

type BookModel struct {
	gorm.Model
	Title    string
	Author   string
	Location string
}

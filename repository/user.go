package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       *int         `gorm:"default:18"`
	Active    sql.NullBool `gorm:"default:true"`
	CompanyID int
	Company   Company
}

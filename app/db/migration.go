package db

import (
	"github.com/jinzhu/gorm"
)

// Init ...
func Init(db *gorm.DB, t interface{}) {

	defer db.Close()

	db.AutoMigrate(t)

}

// CreateMigrate ...,HasTable
func CreateMigrate(db *gorm.DB, t interface{}) {

	defer db.Close()

	db.AutoMigrate(t)
}

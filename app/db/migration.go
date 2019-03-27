package db

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Init ...
func Init(db *gorm.DB, t ...interface{}) {

	defer db.Close()

	db.AutoMigrate(t)

}

// CreateMigrate ...
func CreateMigrate(db *gorm.DB, t ...interface{}) {

	defer db.Close()

	for _, n := range t {
		d := db.AutoMigrate(n)
		if d.Error != nil {
			log.Fatal(d.Error)
		}
	}
}

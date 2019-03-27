package db

import (
	"github.com/jinzhu/gorm"
)

// GetAll return all register database
func GetAll(db *gorm.DB) (*OS, error) {
	OS := OS{}
	if result := db.First(&OS); result.Error != nil {
		return nil, result.Error
	}
	return &OS, nil
}

// Get Find register
func Get(db *gorm.DB, nome string) (*OS, error) {
	OS := OS{}
	db.Where("nomeCliente = ?", nome).Find(&OS)
	if db.Error != nil {
		return nil, db.Error
	}
	return &OS, nil
}

// DeleteID ...
func DeleteID(db *gorm.DB, id int) error {
	if db.Delete(OS{}, "id = ?", id); db.Error != nil {
		return db.Error
	}
	return nil
}

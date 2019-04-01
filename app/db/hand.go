package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

// GetAll return all register database
func GetAll(db *gorm.DB) ([]OS, error) {
	OS := []OS{}
	if result := db.First(&OS); result.Error != nil {
		return nil, result.Error
	}
	return OS, nil
}

// GetNome Find register
func GetNome(db *gorm.DB, nome string) ([]OS, error) {
	OS := []OS{}
	db.Where("nomeCliente = ?", nome).Find(&OS)
	if db.Error != nil {
		return nil, db.Error
	}
	return OS, nil
}

// DeleteID ...
func DeleteID(db *gorm.DB, id int) error {
	if db.Delete(OS{}, "id = ?", id); db.Error != nil {
		return db.Error
	}
	return nil
}

// Get ...
func Get(db *gorm.DB, f map[string]interface{}) (*[]OS, error) {
	OS := []OS{}
	if len(f) == 0 {
		return nil, errors.New("Parameter is empty")
	}
	query := db.New()
	for col, val := range f {
		query.Where(fmt.Sprintf("%s = ?", col), val)
	}
	query.Find(&OS)

	return &OS, nil
}

// Insert ...
func Insert(db *gorm.DB, OS OS) (OS, error) {
	result := OS

	db.Create(&result)
	if db.Error != nil {
		return result, db.Error
	}

	return result, nil
}

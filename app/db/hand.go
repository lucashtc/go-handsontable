package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// GetAll return all register database
func GetAll(db *gorm.DB) ([]OS, error) {
	var os []OS
	db.Find(&os)
	if db.Error != nil {
		return nil, errors.Wrap(db.Error,"fail to get all")
	}
	return os, nil
}

// GetNome Find register
func GetNome(db *gorm.DB, nome string) ([]OS, error) {
	OS := []OS{}
	db.Where("nomeCliente = ?", nome).Find(&OS)
	if db.Error != nil {
		return nil, errors.Wrap(db.Error,"fail to get nome")
	}
	return OS, nil
}

// DeleteID ...
func DeleteID(db *gorm.DB, id int) error {
	if db.Delete(OS{}, "id = ?", id); db.Error != nil {
		return errors.Wrap(db.Error,"fail delete")
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
		return result, errors.Wrap(db.Error,"fail insert")
	}

	return result, nil
}

// Update ...
func Update(db *gorm.DB,os OS) error {
	db.Save(&os)
	if db.Error != nil {
		return errors.Wrap(db.Error, "fail update")
	}
	return nil
}
package db

import (
	"github.com/jinzhu/gorm"
)

// OS ...
type OS struct {
	gorm.Model
	NumeroOs    int
	NomeCliente string
	Cnpj14      int
}

// User ...
type User struct {
	gorm.Model
	User string
	Password string
}


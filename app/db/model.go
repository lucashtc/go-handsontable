package db

import (
	"github.com/jinzhu/gorm"
)

// OS ...
type OS struct {
	gorm.Model
	numeroOs    int
	nomeCliente string
	cnpj14      int
}

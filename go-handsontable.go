package main

import (
	"log"

	// Drive
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucashtc/go-handsontable/app/db"

	"github.com/jinzhu/gorm"
)

func main() {

	//Exemplo para criar migration

	// Conn database
	conn, err := gorm.Open("mysql", "root@/test")
	if err != nil {
		log.Fatal(err)
	}

	// Define struct
	type table struct {
		gorm.Model
		Name string
		Age  int
	}

	db.CreateMigrate(conn, &table{})

}

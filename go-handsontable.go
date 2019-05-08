package main

import (
	"log"

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

	//os := db.OS{}
	user := db.User{}

	//db.CreateMigrate(conn, &os)
	db.CreateMigrate(conn, &user)

}

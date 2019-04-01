package db

import (
	"log"
	"testing"

	"github.com/jinzhu/gorm"
)

func conn() *gorm.DB {
	conn, err := gorm.Open("mysql", "root@/test")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
func TestGetAll(t *testing.T) {
	conn := conn()

	_, err := GetAll(conn)
	if err != nil {
		t.Error(err)
	}
}

func TestInsert(t *testing.T) {
	
}

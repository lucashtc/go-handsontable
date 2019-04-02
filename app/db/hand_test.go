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
	data := []OS{
		OS{
			NomeCliente: "Lucas Da costa",
			NumeroOs:    11,
			Cnpj14:      123123123,
		},
		OS{
			NomeCliente: "Outro Lucas",
			NumeroOs:    12,
			Cnpj14:      33213,
		},
	}

	for _, os := range data {
		conn := conn()
		_, err := Insert(conn, os)
		if err != nil {
			t.Fatal(err)
		}
	}

}

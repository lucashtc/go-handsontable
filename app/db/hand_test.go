package db

import (
	"log"
	"testing"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func conn() *gorm.DB {
	conn, err := gorm.Open("mysql", "root@/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if conn.Error != nil {
		log.Fatal(conn.Error)
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

func TestUpdate(t *testing.T) {
	conn := conn()

	os, err := GetAll(conn)
	if err != nil {
		assert.Error(t, err,"fail to get all data")
	}

	for i, v := range os {
		
		v.NomeCliente = fmt.Sprintf("addteste %d",i)
		v.Cnpj14 = 1234567890

		err = Update(conn, v)
		if err != nil {
			assert.Error(t,err,"fail to perfom update")
		}
	}	
}

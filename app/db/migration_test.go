package db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func TestCreateMigrate(t *testing.T) {
	conn, err := gorm.Open("mysql", "root@/test")
	if err != nil {
		t.Error(err)
	}

	type test struct {
		gorm.Model
		Name      string
		OtherName string
	}
	type otherTest struct {
		gorm.Model
		Name      string
		OtherName string
	}

	CreateMigrate(conn, &OS{})

}

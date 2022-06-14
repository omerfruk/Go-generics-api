package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	DBNAME = "go_generics"
	DBUSER = "go_generics"
	DBPASS = "go_generics"
	DBHOST = "localhost"
	DBPORT = "5432"
)

func DBConnect() {
	var err error
	dsn := fmt.Sprintf("host=%s user='%s' password=%s dbname=%s port=%s sslmode=disable", DBHOST, DBUSER, DBPASS, DBNAME, DBPORT)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return db
}

func AutoMigrate() {
	db.Migrator()
}
